package handler

import (
	"net/http"

	"github.com/felipeversiane/go-auth/config/logger"
	"github.com/felipeversiane/go-auth/config/validation"
	domain "github.com/felipeversiane/go-auth/internal/account"
	"github.com/felipeversiane/go-auth/internal/account/handler/entity/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uh *userHandlerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser controller",
		zap.String("journey", "loginUser"),
	)
	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "loginUser"))
		errRest := validation.ValidateError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := domain.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)
	domainResult, token, err := uh.service.LoginUserService(domain)
	if err != nil {
		logger.Error(
			"Error trying to call loginUser service",
			err,
			zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"loginUser controller executed successfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "loginUser"))

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, token)
}
