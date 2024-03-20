package service

import (
	"github.com/felipeversiane/go-auth/config/httperr"
	"github.com/felipeversiane/go-auth/config/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserService(userId string) *httperr.HttpError {
	logger.Info("Init deleteUser service.",
		zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "deleteUser"))
		return err
	}
	logger.Info(
		"deleteUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"))
	return nil
}
