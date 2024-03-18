package handler

import (
	"net/http"

	"github.com/felipeversiane/go-auth/config/httperr"
	"github.com/felipeversiane/go-auth/config/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (uh *userHandlerInterface) VerifyToken(c *gin.Context) {
	logger.Info("Init verifyToken controller",
		zap.String("journey", "verifyToken"))

	token := c.GetHeader("Authorization")
	if token == "" {
		logger.Error("Authorization token not provided", nil,
			zap.String("journey", "verifyToken"))
		c.JSON(http.StatusBadRequest, httperr.NewUnauthorizedRequestError("Authorization token not provided."))
		return
	}

	user, err := uh.service.VerifyTokenService(token)
	if err != nil {
		logger.Error("Error trying to verify token", err,
			zap.String("journey", "verifyToken"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("verifyToken controller executed successfully",
		zap.String("userId", user.GetID()), zap.String("journey", "verifyToken"))

	c.JSON(http.StatusOK, gin.H{
		"user_id": user.GetID(),
		"message": "Token verified successfully",
	})
}
