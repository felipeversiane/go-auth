package router

import (
	"github.com/felipeversiane/go-auth/config/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func InitializeRoutes(database *mongo.Database) error {

	logger.Info("Initializing Routes.",
		zap.String("journey", "Initialize Routes"))

	router := gin.Default()

	AppRoutes(router, database)

	err := router.Run()

	if err != nil {
		return err
	}

	return nil
}
