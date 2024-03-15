package service

import (
	"github.com/felipeversiane/go-auth/config/httperr"
	domain "github.com/felipeversiane/go-auth/internal/account"
)

func (ud *userDomainService) LoginUserService(userDomain *domain.UserDomainInterface,
) (domain.UserDomainInterface, string, *httperr.HttpError) {
	return nil, "", nil
}
