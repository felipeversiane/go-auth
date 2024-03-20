package handler

import (
	"github.com/felipeversiane/go-auth/internal/account/service"
	"github.com/gin-gonic/gin"
)

func NewUserHandlerInterface(
	serviceInterface service.UserDomainService,
) UserHandlerInterface {
	return &userHandlerInterface{
		service: serviceInterface,
	}
}

type UserHandlerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	LoginUser(c *gin.Context)
	VerifyToken(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type userHandlerInterface struct {
	service service.UserDomainService
}
