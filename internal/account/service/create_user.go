package service

import (
	"github.com/felipeversiane/golang-auth/config/httperr"
	"github.com/felipeversiane/golang-auth/config/logger"
	domain "github.com/felipeversiane/golang-auth/internal/account"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserService(
	userDomain domain.UserDomainInterface,
) (domain.UserDomainInterface, *httperr.HttpError) {
	logger.Info("Init createUser Service.",
		zap.String("journey", "createUser"))

	user, _ := ud.FindUserByEmailService(userDomain.GetEmail())
	if user != nil {
		return nil, httperr.NewBadRequestError("Email is already registered in another account")
	}

	userDomain.EncryptPassword()
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info(
		"CreateUser service executed successfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))
	return userDomainRepository, nil
}
