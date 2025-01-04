package repo

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"resume/internal/models"
)

type ResumeRepository interface {
	Find(id string) (*models.Resume, error)
	List(userId string) ([]*models.Resume, error)
	Create(resume *models.Resume) error
	Update(resume *models.Resume) error
	Delete(id string) error
}

type resumeRepository struct {
	conn   *gorm.DB
	logger *zap.Logger
}

func NewResumeRepository(logger *zap.Logger, conn *gorm.DB) ResumeRepository {
	return &resumeRepository{
		conn:   conn,
		logger: logger,
	}
}

func (r *resumeRepository) Find(id string) (*models.Resume, error) {
	return nil, nil
}

func (r *resumeRepository) List(userId string) ([]*models.Resume, error) {
	return nil, nil
}

func (r *resumeRepository) Create(resume *models.Resume) error {
	return nil
}

func (r *resumeRepository) Update(resume *models.Resume) error {
	return nil
}

func (r *resumeRepository) Delete(id string) error {
	return nil
}
