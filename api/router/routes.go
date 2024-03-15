package router

import (
	"github.com/felipeversiane/go-auth/internal/account/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AppRoutes(router *gin.Engine, database *mongo.Database) *gin.RouterGroup {
	v1 := router.Group("/v1/api")
	{
		routes.AccountRoutes(v1, database)
	}
	return v1
}
