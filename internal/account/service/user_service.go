package service

import (
	"github.com/felipeversiane/golang-auth/config/httperr"
	domain "github.com/felipeversiane/golang-auth/internal/account"
	"github.com/felipeversiane/golang-auth/internal/account/repository"
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
	LoginUserService(userDomain *domain.UserDomainInterface) (domain.UserDomainInterface, string, *httperr.HttpError)
}
