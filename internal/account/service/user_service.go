package service

import (
	"github.com/felipeversiane/go-auth/config/httperr"
	domain "github.com/felipeversiane/go-auth/internal/account"
	"github.com/felipeversiane/go-auth/internal/account/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserService(domain.UserDomainInterface) (
		domain.UserDomainInterface, *httperr.HttpError)

	FindUserByIDService(
		id string,
	) (domain.UserDomainInterface, *httperr.HttpError)
	FindUserByEmailService(
		email string,
	) (domain.UserDomainInterface, *httperr.HttpError)

	UpdateUserService(string, domain.UserDomainInterface) *httperr.HttpError
	DeleteUserService(string) *httperr.HttpError
	LoginUserService(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, string, string, *httperr.HttpError)
	VerifyTokenService(string) (domain.UserDomainInterface, *httperr.HttpError)
	RefreshTokenService(string) (string, *httperr.HttpError)

}
