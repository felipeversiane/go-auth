package repository

import (
	"github.com/felipeversiane/go-auth/config/httperr"
	domain "github.com/felipeversiane/go-auth/internal/account"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_DB = "DB_USER"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain domain.UserDomainInterface,
	) (domain.UserDomainInterface, *httperr.HttpError)

	UpdateUser(
		userId string,
		userDomain domain.UserDomainInterface,
	) *httperr.HttpError

	DeleteUser(
		userId string,
	) *httperr.HttpError

	FindUserByEmail(
		email string,
	) (domain.UserDomainInterface, *httperr.HttpError)
	FindUserByEmailAndPassword(
		email string,
		password string,
	) (domain.UserDomainInterface, *httperr.HttpError)
	FindUserByID(
		id string,
	) (domain.UserDomainInterface, *httperr.HttpError)
}
