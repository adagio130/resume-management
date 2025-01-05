package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"resume/internal/models"
	"resume/internal/reqs"

	customerors "resume/internal/errors"
)

var validate = validator.New()

type Handler struct {
	logger  *zap.Logger
	service Service
}

func NewHandler(logger *zap.Logger, service Service) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

func (h *Handler) GetResume(c *gin.Context) {
	resumeId := c.Param("id")
	if resumeId == "" {
		c.Error(customerors.GetError(customerors.ErrBadRequest))
		return
	}
	resume, err := h.service.GetResume(resumeId)
	if err != nil {
		h.logger.Error("Failed to get resume", zap.Error(err))
		c.Error(err)
	}
	if resume == nil {
		c.Error(customerors.GetError(customerors.ErrResumeNotFound))
		return
	}
	c.JSON(http.StatusOK, resume)
}

func (h *Handler) CreateResume(c *gin.Context) {
	var createResumeReq reqs.CreateResumeRequest
	c.Header("Content-Type", "application/json")
	if err := c.BindJSON(&createResumeReq); err != nil {
		h.logger.Error("Failed to bind JSON", zap.Error(err))
		c.Error(customerors.GetError(customerors.ErrBadRequest))
		return
	}

	if err := validate.Struct(&createResumeReq); err != nil {
		h.logger.Error("Validation failed", zap.Error(err))
		c.Error(customerors.GetError(customerors.ErrBadRequest))
		return
	}

	resumeId, err := h.service.CreateResume(&createResumeReq)
	if err != nil {
		h.logger.Error("Failed to create resume", zap.Error(err))
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": resumeId})
}

func (h *Handler) UpdateResume(c *gin.Context) {
	resumeId := c.Param("id")
	if resumeId == "" {
		c.Error(customerors.GetError(customerors.ErrBadRequest))
		return
	}
	var updateResumeReq reqs.UpdateResumeRequest
	if err := c.BindJSON(&updateResumeReq); err != nil {
		h.logger.Error("Failed to bind JSON", zap.Error(err))
		c.Error(customerors.GetError(customerors.ErrBadRequest))
		return
	}
	if err := validate.Struct(&updateResumeReq); err != nil {
		h.logger.Error("Validation failed", zap.Error(err))
		c.Error(customerors.GetError(customerors.ErrBadRequest))
		return
	}
	resumeId, err := h.service.UpdateResume(resumeId, &updateResumeReq)
	if err != nil {
		h.logger.Error("Failed to update resume", zap.Error(err))
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": resumeId})
}

func (h *Handler) DeleteResume(c *gin.Context) {
	resumeId := c.Param("id")
	if resumeId == "" {
		c.Error(customerors.GetError(customerors.ErrBadRequest))
		return
	}
	err := h.service.DeleteResume(resumeId)
	if err != nil {
		h.logger.Error("Failed to delete resume", zap.Error(err))
		c.Error(err)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) GetResumes(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.Error(customerors.GetError(customerors.ErrBadRequest))
		return
	}
	resumes, err := h.service.ListResume(userId)
	if err != nil {
		h.logger.Error("Failed to get resumes", zap.Error(err))
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, map[string][]*models.Resume{"resumes": resumes})
}

func (h *Handler) CreateUser(c *gin.Context) {
	var createUserReq reqs.CreateUserRequest
	c.Header("Content-Type", "application/json")
	if err := c.BindJSON(&createUserReq); err != nil {
		h.logger.Error("Failed to bind JSON", zap.Error(err))
		c.Error(customerors.GetError(customerors.ErrBadRequest))
		return
	}
	if err := validate.Struct(&createUserReq); err != nil {
		h.logger.Error("Validation failed", zap.Error(err))
		c.Error(customerors.GetError(customerors.ErrBadRequest))
		return
	}
	userId, err := h.service.CreateUser(&createUserReq)
	if err != nil {
		h.logger.Error("Failed to create user", zap.Error(err))
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": userId})
}
