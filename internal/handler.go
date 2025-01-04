package internal

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"resume/internal/reqs"
)

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
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

}

func (h *Handler) CreateResume(c *gin.Context) {
	h.logger.Info("CreateResume")
	c.AbortWithStatus(http.StatusOK)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	userId, err := h.service.CreateUser(&createUserReq)
	if err != nil {
		h.logger.Error("Failed to create user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": userId})
}
