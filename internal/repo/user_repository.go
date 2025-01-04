package repo

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"resume/internal/entities"
	"resume/internal/models"
	"strings"
)

type UserRepository interface {
	CreateUser(user *models.User) (string, error)
	GetUser(id string) *models.User
}

type userRepo struct {
	logger *zap.Logger
	conn   *gorm.DB
}

func NewUserRepository(logger *zap.Logger, conn *gorm.DB) UserRepository {
	return &userRepo{
		conn:   conn,
		logger: logger,
	}
}

func (u *userRepo) CreateUser(user *models.User) (string, error) {
	entity := entities.NewUserEntity(user.ID, user.Account, user.Name, user.Gender, user.Location)
	result := u.conn.Create(entity)
	if result.Error != nil {
		u.logger.Error("Failed to create user", zap.Error(result.Error))
		if strings.Contains(result.Error.Error(), "1062") {
			return "", models.ErrUserExist
		}
		return "", result.Error
	}
	return entity.ID, nil
}

func (u *userRepo) GetUser(id string) *models.User {
	return nil
}
