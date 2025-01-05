package repo

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"resume/internal/entities"
	customerror "resume/internal/errors"
	"resume/internal/models"
)

type ResumeRepository interface {
	Find(id string) (*models.Resume, error)
	List(userId string) ([]*models.Resume, error)
	Create(resume *models.Resume) (string, error)
	Update(id string, resume *models.Resume) (string, error)
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
	r.logger.Info("Find")
	entity := &entities.Resume{}
	result := r.conn.Preload("Experiences").Preload("Educations").Where("id = ?", id).Take(entity)
	if result.Error != nil {
		r.logger.Error("Failed to find resume", zap.Error(result.Error))
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, customerror.GetError(customerror.ErrResumeNotFound)
	}
	resume := &models.Resume{
		ID:         entity.ID,
		UserID:     entity.UserID,
		Title:      entity.Title,
		Email:      entity.Email,
		Phone:      entity.Phone,
		Skills:     entity.Skills,
		Education:  make(map[string]*models.Education, 0),
		Experience: make(map[string]*models.Experience, 0),
	}

	for _, exp := range entity.Experiences {
		start := exp.StartDate.Format("2006-01")
		end := ""
		if exp.EndDate != nil {
			end = exp.EndDate.Format("2006-01")
		}
		resume.Experience[exp.ID] = &models.Experience{
			ID:          exp.ID,
			Company:     exp.Company,
			Position:    exp.Position,
			IsPresent:   exp.IsPresent,
			StartDate:   start,
			EndDate:     end,
			Description: exp.Description,
		}
	}

	for _, edu := range entity.Educations {
		start := edu.StartDate.Format("2006-01")
		end := ""
		if edu.EndDate != nil {
			end = edu.EndDate.Format("2006-01")
		}
		resume.Education[edu.ID] = &models.Education{
			ID:        edu.ID,
			School:    edu.School,
			Major:     edu.Major,
			Degree:    edu.Degree,
			StartDate: start,
			EndDate:   end,
		}
	}
	return resume, nil
}

func (r *resumeRepository) List(userId string) ([]*models.Resume, error) {
	r.logger.Info("List")

	var results []entities.Resume

	err := r.conn.Model(&entities.Resume{}).
		Select("id, title").
		Where("user_id = ?", userId).
		Find(&results).Error

	if err != nil {
		r.logger.Error("Failed to list resumes", zap.Error(err))
		return nil, err
	}

	var resumes []*models.Resume
	for _, res := range results {
		resumes = append(resumes, &models.Resume{
			ID:    res.ID,
			Title: res.Title,
		})
	}

	return resumes, nil
}

func (r *resumeRepository) Create(resume *models.Resume) (string, error) {
	r.logger.Info("Create")
	resumeEntity := entities.NewResumeEntity(resume.ID, resume.UserID, resume.Title, resume.Email, resume.Phone, resume.Skills)
	experiences := make([]*entities.Experience, len(resume.Experience))
	i := 0
	for _, exp := range resume.Experience {
		experiences[i] = entities.NewExperienceEntity(exp.ID, resume.ID, exp.Company, exp.Position, exp.IsPresent, exp.StartDate, exp.EndDate, exp.Description)
		i++
	}
	educations := make([]*entities.Education, len(resume.Education))
	i = 0
	for _, edu := range resume.Education {
		educations[i] = entities.NewEducationEntity(edu.ID, resume.ID, edu.School, edu.Major, edu.Degree, edu.StartDate, edu.EndDate)
		i++
	}
	err := r.conn.Transaction(func(tx *gorm.DB) error {
		result := tx.Create(resumeEntity)
		if result.Error != nil {
			r.logger.Error("Failed to create resume", zap.Error(result.Error))
			return result.Error
		}
		for _, exp := range experiences {
			result = tx.Create(exp)
			if result.Error != nil {
				r.logger.Error("Failed to create experience", zap.Error(result.Error))
				return result.Error
			}
		}
		for _, edu := range educations {
			result = tx.Create(edu)
			if result.Error != nil {
				r.logger.Error("Failed to create education", zap.Error(result.Error))
				return result.Error
			}
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return resume.ID, nil
}

func (r *resumeRepository) Update(id string, resume *models.Resume) (string, error) {
	r.logger.Info("Update")
	resumeEntity := entities.NewResumeEntity(resume.ID, resume.UserID, resume.Title, resume.Email, resume.Phone, resume.Skills)
	experiences := make([]*entities.Experience, len(resume.Experience))
	i := 0
	for _, exp := range resume.Experience {
		experiences[i] = entities.NewExperienceEntity(exp.ID, resume.ID, exp.Company, exp.Position, exp.IsPresent, exp.StartDate, exp.EndDate, exp.Description)
		i++
	}
	educations := make([]*entities.Education, len(resume.Education))
	i = 0
	for _, edu := range resume.Education {
		educations[i] = entities.NewEducationEntity(edu.ID, resume.ID, edu.School, edu.Major, edu.Degree, edu.StartDate, edu.EndDate)
		i++
	}
	err := r.conn.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("id = ?", id).Updates(resumeEntity)
		if result.Error != nil {
			r.logger.Error("Failed to update resume", zap.Error(result.Error))
			return result.Error
		}
		for _, exp := range experiences {
			result = tx.Where("id = ?", exp.ID).Updates(exp)
			if result.Error != nil {
				r.logger.Error("Failed to update experience", zap.Error(result.Error))
				return result.Error
			}
		}
		for _, edu := range educations {
			result = tx.Where("id = ?", edu.ID).Updates(edu)
			if result.Error != nil {
				r.logger.Error("Failed to update education", zap.Error(result.Error))
				return result.Error
			}
		}
		return nil

	})
	if err != nil {
		return "", err
	}
	return resume.ID, nil
}

func (r *resumeRepository) Delete(id string) error {
	r.logger.Info("Delete")
	result := r.conn.Where("id = ?", id).Delete(&entities.Resume{})
	if result.Error != nil {
		r.logger.Error("Failed to delete resume", zap.Error(result.Error))
		return result.Error
	}
	if result.RowsAffected == 0 {
		return customerror.GetError(customerror.ErrResumeNotFound)
	}
	return nil
}
