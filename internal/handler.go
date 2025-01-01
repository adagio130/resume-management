package internal

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"resume/internal/repo"
)

type Handler struct {
	resumeRepo repo.ResumeRepository
}

func NewHandler(logger *zap.Logger, repository repo.ResumeRepository) *Handler {
	return &Handler{
		resumeRepo: repository,
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
	c.AbortWithStatus(http.StatusOK)
}

func (h *Handler) UpdateResume(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)

}

func (h *Handler) DeleteResume(c *gin.Context) {

}

func (h *Handler) GetResumes(c *gin.Context) {

}
