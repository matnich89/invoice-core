package handler

import (
	"github.com/gin-gonic/gin"
	"invoice-core/cmd/server/internal/logger"
	"invoice-core/cmd/server/internal/model"
	"log"
	"net/http"
)

type Handler struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *Handler {
	return &Handler{logger: logger}
}

func (h *Handler) HealthCheck(c *gin.Context) {
	h.logger.Info("recieved  health request... ")
	c.Writer.WriteHeader(http.StatusOK)
}

func (h *Handler) SayHelloToSonal(c *gin.Context) {
	log.Println("Hello Sonal")

}

func (h *Handler) CreateInvoice(c *gin.Context) {
	var data model.Data
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.logger.Info(data.Client)
	c.JSON(http.StatusOK, data)

}
