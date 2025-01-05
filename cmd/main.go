package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"resume/config"
	"resume/internal"
	"resume/internal/entities"
	"resume/internal/middleware"
	"resume/internal/repo"
)

func main() {
	conf := initConfig()
	logger := initLogger(conf)
	db := initStorage(conf)
	handler := initHandler(logger, db)
	server := gin.Default()
	server.Use(middleware.ErrorHandler())
	server.POST("/user", func(c *gin.Context) {
		handler.CreateUser(c)
	})
	server.GET("/resume/:id", func(c *gin.Context) {
		handler.GetResume(c)
	})
	server.POST("/resume", func(c *gin.Context) {
		handler.CreateResume(c)
	})
	server.PUT("/resume/:id", func(c *gin.Context) {
		handler.UpdateResume(c)
	})
	server.DELETE("/resume/:id", func(c *gin.Context) {
		handler.DeleteResume(c)
	})
	server.GET("/resumes", func(c *gin.Context) {
		handler.GetResumes(c)
	})
	err := server.Run(fmt.Sprintf(":%s", conf.Server.Port))
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func initHandler(logger *zap.Logger, db *gorm.DB) *internal.Handler {
	resumeRepo := repo.NewResumeRepository(logger, db)
	userRepo := repo.NewUserRepository(logger, db)
	service := internal.NewService(logger, userRepo, resumeRepo)
	return internal.NewHandler(logger, service)
}

func initConfig() *config.Config {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	var conf config.Config
	if err := v.Unmarshal(&conf); err != nil {
		panic(fmt.Errorf("umarshal config error: %s \n", err))
	}
	v.WatchConfig()
	return &conf
}

// initStorage initializes the database connection
func initStorage(configs *config.Config) *gorm.DB {
	conn, err := gorm.Open(mysql.Open(configs.DB.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	db, err := conn.DB()
	if err != nil {
		log.Fatalf("failed to get database: %v", err)
		return nil
	}
	db.SetMaxIdleConns(configs.DB.MaxIdleConns)
	db.SetMaxOpenConns(configs.DB.MaxOpenConns)
	db.SetConnMaxLifetime(configs.DB.ConnMaxLifetime)
	err = conn.AutoMigrate(&entities.User{}, &entities.Resume{}, &entities.Education{}, &entities.Experience{})
	if err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
		return nil
	}
	return conn
}

func initLogger(configs *config.Config) *zap.Logger {
	switch configs.Log.Env {
	case "dev":
		logger, err := zap.NewDevelopment()
		if err != nil {
			log.Fatalf("failed to create logger: %v", err)
		}
		return logger
	case "test":
		logger, err := zap.NewDevelopment()
		if err != nil {
			log.Fatalf("failed to create logger: %v", err)
		}
		return logger
	}
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	return logger
}
