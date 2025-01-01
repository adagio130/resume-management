package repo

import (
	"gorm.io/gorm"
	"resume/internal/entities"
)

type ResumeRepository interface {
	Find(id string) (*entities.Resume, error)
	List(param entities.ResumeQueryParam) ([]*entities.Resume, error)
	Create(resume entities.Resume) error
	Update(resume entities.Resume) error
	Delete(id string) error
}

type resumeRepository struct {
	conn *gorm.DB
}

func NewResumeRepository(conn *gorm.DB) ResumeRepository {
	return &resumeRepository{conn: conn}
}

func (r *resumeRepository) Find(id string) (*entities.Resume, error) {
	return nil, nil
}

func (r *resumeRepository) List(param entities.ResumeQueryParam) ([]*entities.Resume, error) {
	return nil, nil
}

func (r *resumeRepository) Create(resume entities.Resume) error {
	return nil
}

func (r *resumeRepository) Update(resume entities.Resume) error {
	return nil
}

func (r *resumeRepository) Delete(id string) error {
	return nil
}
