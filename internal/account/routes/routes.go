package routes

import (
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
		account.GET("/get_user_by_id/:user_id", handler.FindUserByID)
		account.GET("/get_user_by_email/:user_email", handler.FindUserByEmail)
		account.POST("/create_user", handler.CreateUser)
		account.PUT("/update_user/:user_id", handler.UpdateUser)
		account.DELETE("/delete_user/:user_id", handler.DeleteUser)
	}

	return account
}
