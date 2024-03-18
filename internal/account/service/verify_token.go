package service

import (
	"github.com/felipeversiane/go-auth/config/httperr"
	"github.com/felipeversiane/go-auth/config/logger"
	domain "github.com/felipeversiane/go-auth/internal/account"
	"go.uber.org/zap"
)

func (ud *userDomainService) VerifyTokenService(tokenValue string) (domain.UserDomainInterface, *httperr.HttpError) {
	logger.Info("Init verifyToken service.",
		zap.String("journey", "verifyToken"))

	user, err := domain.VerifyToken(tokenValue)
	if err != nil {
		return nil, err
	}

	logger.Info("verifyToken service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "verifyToken"))

	return user, nil
}
