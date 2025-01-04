package internal

import (
	"errors"
	"go.uber.org/zap"
	"resume/internal/models"
	"resume/internal/repo"
	"resume/internal/reqs"
)

type Service interface {
	CreateUser(user *reqs.CreateUserRequest) (string, error)
	GetUser(id string) (*models.User, error)
	CreateResume(resume *reqs.CreateResumeRequest) (string, error)
	GetResume(id string) (*models.Resume, error)
	UpdateResume(request *reqs.UpdateResumeRequest) (string, error)
	DeleteResume(id string) error
}

type service struct {
	logger     *zap.Logger
	resumeRepo repo.ResumeRepository
	userRepo   repo.UserRepository
}

func NewService(logger *zap.Logger, userRepo repo.UserRepository, ResumeRepo repo.ResumeRepository) Service {
	return &service{
		logger:     logger,
		userRepo:   userRepo,
		resumeRepo: ResumeRepo,
	}
}

func (s *service) CreateUser(req *reqs.CreateUserRequest) (string, error) {
	s.logger.Info("CreateUser")
	userModel := models.NewUser(req.Name, req.Account, req.Gender, req.Location)
	userId, err := s.userRepo.CreateUser(userModel)
	if err != nil {
		if errors.Is(err, models.ErrUserExist) {
			return "", models.ErrUserExist
		}
		s.logger.Error("Failed to create user", zap.Error(err))
		return "", err
	}
	return userId, nil
}

func (s *service) GetUser(id string) (*models.User, error) {
	return nil, nil
}

func (s *service) CreateResume(req *reqs.CreateResumeRequest) (string, error) {
	return "", nil
}

func (s *service) GetResume(id string) (*models.Resume, error) {
	return nil, nil
}

func (s *service) UpdateResume(req *reqs.UpdateResumeRequest) (string, error) {
	return "", nil
}

func (s *service) DeleteResume(id string) error {
	return nil
}
