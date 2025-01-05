package internal

import (
	"go.uber.org/zap"
	custom_error "resume/internal/errors"
	"resume/internal/models"
	"resume/internal/repo"
	"resume/internal/reqs"
)

type Service interface {
	CreateUser(user *reqs.CreateUserRequest) (string, error)
	GetUser(id string) (*models.User, error)
	CreateResume(resume *reqs.CreateResumeRequest) (string, error)
	GetResume(id string) (*models.Resume, error)
	UpdateResume(id string, request *reqs.UpdateResumeRequest) (string, error)
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
		s.logger.Error("Failed to create user", zap.Error(err))
		return "", err
	}
	return userId, nil
}

func (s *service) GetUser(id string) (*models.User, error) {
	return nil, nil
}

func (s *service) CreateResume(req *reqs.CreateResumeRequest) (string, error) {
	s.logger.Info("CreateResume")
	isUserExist, err := s.userRepo.GetUser(req.UserID)
	if err != nil {
		s.logger.Error("Failed to get user", zap.Error(err))
		return "", err
	}
	if isUserExist == nil {
		return "", custom_error.GetError(custom_error.ErrUserNotFound)
	}
	resumeModel := models.NewResumeFromReqs(req.UserID, req.Title, req.Email, req.Phone, req.Experience, req.Skills, req.Education)
	resumeId, err := s.resumeRepo.Create(resumeModel)
	if err != nil {
		s.logger.Error("Failed to create resume", zap.Error(err))
		return "", err
	}
	return resumeId, nil
}

func (s *service) GetResume(id string) (*models.Resume, error) {
	s.logger.Info("GetResume")
	resume, err := s.resumeRepo.Find(id)
	if err != nil {
		s.logger.Error("Failed to get resume", zap.Error(err))
		return nil, err
	}
	return resume, nil
}

func (s *service) UpdateResume(id string, req *reqs.UpdateResumeRequest) (string, error) {
	s.logger.Info("UpdateResume")

	return "", nil
}

func (s *service) DeleteResume(id string) error {
	return nil
}
