package routes

import (
	"github.com/felipeversiane/go-auth/api/middleware"
	"github.com/felipeversiane/go-auth/internal/account/handler"
	"github.com/felipeversiane/go-auth/internal/account/repository"
	"github.com/felipeversiane/go-auth/internal/account/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func AccountRoutes(v1 *gin.RouterGroup, database *mongo.Database) *gin.RouterGroup {

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	handler := handler.NewUserHandlerInterface(service)

	account := v1.Group("/account")
	{
		account.GET("/get_user_by_id/:userId", handler.FindUserByID)
		account.GET("/get_user_by_email/:userEmail", middleware.VerifyTokenMiddleware, handler.FindUserByEmail)
		account.POST("/create_user", handler.CreateUser)
		account.PUT("/update_user/:userId", handler.UpdateUser)
		account.DELETE("/delete_user/:userId", handler.DeleteUser)

		account.POST("/jwt/login", handler.LoginUser)
		account.POST("/jwt/refresh")
		account.POST("/jwt/verify", handler.VerifyToken)
		account.POST("/jwt/logout")

	}

	return account
}
