package service

import (
	"github.com/felipeversiane/golang-auth/config/httperr"
	"github.com/felipeversiane/golang-auth/config/logger"
	domain "github.com/felipeversiane/golang-auth/internal/account"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserService(
	userId string,
	userDomain domain.UserDomainInterface,
) *httperr.HttpError {
	logger.Info("Init updateUser model.",
		zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateUser"))
		return err
	}

	logger.Info(
		"updateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))
	return nil
}
