package service

import (
	"github.com/felipeversiane/golang-auth/config/httperr"
	"github.com/felipeversiane/golang-auth/config/logger"
	domain "github.com/felipeversiane/golang-auth/internal/account"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDService(
	id string,
) (domain.UserDomainInterface, *httperr.HttpError) {
	logger.Info("Init findUserByID services.",
		zap.String("journey", "findUserById"))

	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailService(
	email string,
) (domain.UserDomainInterface, *httperr.HttpError) {
	logger.Info("Init findUserByID services.",
		zap.String("journey", "findUserById"))

	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) findUserByEmailAndPasswordService(
	email string,
	password string,
) (domain.UserDomainInterface, *httperr.HttpError) {
	logger.Info("Init findUserByID services.",
		zap.String("journey", "findUserById"))

	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
