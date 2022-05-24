package main

import (
	"github.com/gin-gonic/gin"
	"invoice-core/cmd/server/handler"
	"invoice-core/cmd/server/internal/logger"
)

type App struct {
	router  *gin.Engine
	logger  *logger.Logger
	handler *handler.Handler
}

func NewApp(router *gin.Engine, logger *logger.Logger) *App {
	return &App{
		router: gin.Default(),
		logger: logger,
	}
}

func (a *App) defineRoutes() {
	a.logger.Info("defining routes....")
	a.router.POST("/invoice", a.handler.CreateInvoice)
	a.router.GET("/healthcheck", a.handler.HealthCheck)
	a.logger.Info("started....")
	_ = a.router.Run()
}
