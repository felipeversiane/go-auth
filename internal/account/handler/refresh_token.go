package handler

import (
	"net/http"

	"github.com/felipeversiane/go-auth/config/httperr"
	"github.com/felipeversiane/go-auth/config/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uh *userHandlerInterface) RefreshToken(c *gin.Context) {
	logger.Info("Init verifyToken handler",
		zap.String("journey", "refreshToken"))

	token := c.GetHeader("Authorization")
	if token == "" {
		logger.Error("Authorization token not provided", nil,
			zap.String("journey", "refreshToken"))
		c.JSON(http.StatusBadRequest, httperr.NewUnauthorizedRequestError("Authorization token not provided."))
		return
	}

	acessToken, err := uh.service.RefreshTokenService(token)
	if err != nil {
		logger.Error("Error trying to refresh token", err,
			zap.String("journey", "refreshToken"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("verifyToken handler executed successfully",
		zap.String("journey", "refreshToken"))

	c.JSON(http.StatusOK, gin.H{
		"acess_token": acessToken,
	})
}
