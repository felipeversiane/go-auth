package router

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AppRoutes(router *gin.Engine, database *mongo.Database) *gin.RouterGroup {
	v1 := router.Group("/v1/api")
	{
	}
	return v1
}
