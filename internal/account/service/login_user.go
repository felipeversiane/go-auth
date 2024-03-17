package service

import (
	"github.com/felipeversiane/go-auth/config/httperr"
	"github.com/felipeversiane/go-auth/config/logger"
	domain "github.com/felipeversiane/go-auth/internal/account"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserService(userDomain domain.UserDomainInterface,
) (domain.UserDomainInterface, string, *httperr.HttpError) {

	logger.Info("Init loginUser model.",
		zap.String("journey", "loginUser"))

	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPasswordService(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info(
		"loginUser service executed successfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "loginUser"))
	return user, token, nil
}
