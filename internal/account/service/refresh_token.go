package service

import (
	"github.com/felipeversiane/go-auth/config/httperr"
	domain "github.com/felipeversiane/go-auth/internal/account"
)

func (ud *userDomainService) RefreshTokenService(refreshToken string) (string, *httperr.HttpError) {

	user, err := domain.VerifyRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}

	accessToken, err := user.GenerateAcessToken()
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
