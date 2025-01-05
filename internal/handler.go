package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"resume/internal/reqs"

	customErrors "resume/internal/errors"
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
		c.Error(customErrors.GetError(customErrors.ErrBadRequest))
		return
	}
	resume, err := h.service.GetResume(resumeId)
	if err != nil {
		h.logger.Error("Failed to get resume", zap.Error(err))
		c.Error(err)
	}
	if resume == nil {
		c.Error(customErrors.GetError(customErrors.ErrResumeNotFound))
		return
	}
	c.JSON(http.StatusOK, resume)
}

func (h *Handler) CreateResume(c *gin.Context) {
	var createResumeReq reqs.CreateResumeRequest
	c.Header("Content-Type", "application/json")
	if err := c.BindJSON(&createResumeReq); err != nil {
		h.logger.Error("Failed to bind JSON", zap.Error(err))
		c.Error(customErrors.GetError(customErrors.ErrBadRequest))
		return
	}

	if err := validate.Struct(&createResumeReq); err != nil {
		h.logger.Error("Validation failed", zap.Error(err))
		c.Error(customErrors.GetError(customErrors.ErrBadRequest))
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
	c.AbortWithStatus(http.StatusOK)

}

func (h *Handler) DeleteResume(c *gin.Context) {

}

func (h *Handler) GetResumes(c *gin.Context) {

}

func (h *Handler) CreateUser(c *gin.Context) {
	var createUserReq reqs.CreateUserRequest
	c.Header("Content-Type", "application/json")
	if err := c.BindJSON(&createUserReq); err != nil {
		h.logger.Error("Failed to bind JSON", zap.Error(err))
		c.Error(customErrors.GetError(customErrors.ErrBadRequest))
		return
	}
	if err := validate.Struct(&createUserReq); err != nil {
		h.logger.Error("Validation failed", zap.Error(err))
		c.Error(customErrors.GetError(customErrors.ErrBadRequest))
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
