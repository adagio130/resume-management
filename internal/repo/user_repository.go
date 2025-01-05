package repo

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"resume/internal/entities"
	custom_error "resume/internal/errors"
	"resume/internal/models"
	"strings"
)

type UserRepository interface {
	CreateUser(user *models.User) (string, error)
	GetUser(id string) (*models.User, error)
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
			return "", custom_error.GetError(custom_error.ErrUserExist)
		}
		return "", result.Error
	}
	return entity.ID, nil
}

func (u *userRepo) GetUser(id string) (*models.User, error) {
	u.logger.Info("GetUser")
	var entity entities.User
	result := u.conn.Where("id = ?", id).Take(&entity)
	if result.Error != nil {
		u.logger.Error("Failed to get user", zap.Error(result.Error))
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, custom_error.GetError(custom_error.ErrUserNotFound)
	}
	model := &models.User{
		ID:       entity.ID,
		Account:  entity.Account,
		Name:     entity.Name,
		Location: entity.Location,
		Gender:   entity.Gender,
	}
	return model, nil
}
