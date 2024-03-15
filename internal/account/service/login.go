package service

import (
	"github.com/felipeversiane/golang-auth/config/httperr"
	domain "github.com/felipeversiane/golang-auth/internal/account"
)

func (ud *userDomainService) LoginUserService(userDomain *domain.UserDomainInterface,
) (domain.UserDomainInterface, string, *httperr.HttpError) {
	return nil, "", nil
}
