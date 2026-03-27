package handler

import (
	"net/http"

	"github.com/DennisMRitchie/go-nlp-service/internal/model"
	"github.com/DennisMRitchie/go-nlp-service/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Handler struct {
	service     *service.NLPService
	rateLimiter *rate.Limiter
}

func NewHandler(svc *service.NLPService) *Handler {
	return &Handler{
		service:     svc,
		rateLimiter: rate.NewLimiter(rate.Limit(100), 200), // 100 requests per second, burst 200
	}
}

func (h *Handler) Analyze(c *gin.Context) {
	if err := h.rateLimiter.Wait(c.Request.Context()); err != nil {
		c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
		return
	}

	var req model.AnalyzeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.Analyze(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
