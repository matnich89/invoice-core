package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"invoice-core/cmd/server/handler"
	"invoice-core/cmd/server/internal/logger"
	"log"
)

func main() {

	zapLogger, err := zap.NewProduction()

	if err != nil {
		log.Fatalln(err.Error())
	}

	logger := logger.New(zapLogger)

	handler := handler.New(logger)

	app := NewApp(gin.Default(), logger, handler)

	app.defineRoutes()

}
