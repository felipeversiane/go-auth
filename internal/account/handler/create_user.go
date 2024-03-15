package handler

import (
	"net/http"

	"github.com/felipeversiane/go-auth/config/logger"
	"github.com/felipeversiane/go-auth/config/validation"
	domain "github.com/felipeversiane/go-auth/internal/account"
	"github.com/felipeversiane/go-auth/internal/account/converter"
	"github.com/felipeversiane/go-auth/internal/account/handler/entity/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uh *userHandlerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser handler.",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := domain.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.FirstName,
		userRequest.LastName,
		userRequest.Age,
	)
	domainResult, err := uh.service.CreateUserService(domain)
	if err != nil {
		logger.Error(
			"Error trying to call CreateUser service",
			err,
			zap.String("journey", "createUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"CreateUser handler executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, converter.ConvertDomainToResponse(
		domainResult,
	))
}
