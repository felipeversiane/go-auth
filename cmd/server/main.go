package main

import (
	"context"
	"os"

	"github.com/felipeversiane/go-auth/config/database/mongodb"
	"github.com/felipeversiane/go-auth/config/logger"
	"github.com/felipeversiane/go-auth/initialization"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func init() {
	logger.Info("Initialize methods is going to run...",
		zap.String("journey", "Initialize"))

	err := initialization.Initialize()

	if err != nil {
		logger.Fatal("Initialize error: ", err,
			zap.String("journey", "Initialize"))
	}

	logger.Info("Initialize methods runned...",
		zap.String("journey", "Initialize"))

	if os.Getenv("MODE") != "development" {
		gin.SetMode(gin.ReleaseMode)
	}
}
func main() {
	logger.Info("The API is going to run...",
		zap.String("journey", "Initialize"))

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		logger.Fatal("Error on create new DB connection.", err,
			zap.String("journey", "NewMongoDBConnection"))
	}

	logger.Info("The Database connection is open now...",
		zap.String("journey", "NewMongoDBConnection"))

}
