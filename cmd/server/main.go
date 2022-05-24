package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"invoice-core/cmd/server/internal/logger"
	"log"
)

func main() {

	zapLogger, err := zap.NewProduction()

	if err != nil {
		log.Fatalln(err.Error())
	}

	logger := logger.New(zapLogger)

	app := NewApp(gin.Default(), logger)

	app.defineRoutes()

}
