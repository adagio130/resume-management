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
	"resume/internal/repo"
)

func main() {
	logger, _ := zap.NewProduction()
	conf := initConfig()
	db := initStorage(conf)
	defer func() {
		logger.Info("Server shutdown")
	}()
	handler := initHandler(logger, db)
	server := gin.Default()

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
	resumeRepo := repo.NewResumeRepository(db)
	return internal.NewHandler(logger, resumeRepo)
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
func initStorage(config *config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.DB.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(&entities.User{})
	if err != nil {
		log.Fatalf("failed to migrate schema: %v", err)
		return nil
	}
	return db
}
