package domain

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/felipeversiane/go-auth/config/httperr"
)

var (
	JWT_SECRET_KEY         = "JWT_SECRET_KEY"
	JWT_REFRESH_SECRET_KEY = "JWT_SECRET_REFRESH_KEY"
)

type userDomain struct {
	id        string
	email     string
	password  string
	firstName string
	lastName  string
	age       int8
	isActive  bool
	createdAt time.Time
	updatedAt time.Time
}

type UserDomainInterface interface {
	SetID(id string)
	GetID() string
	GetEmail() string
	GetIsActive() bool
	SetIsActive(isActive bool)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
	GetFirstName() string
	GetLastName() string
	GetPassword() string
	GetAge() int8
	EncryptPassword()
	GenerateToken() (string, string, *httperr.HttpError)
	GenerateRefreshToken() (string, *httperr.HttpError)
	GenerateAcessToken() (string, *httperr.HttpError)
}

func NewUserDomain(
	email string,
	password string,
	first_name string,
	last_name string,
	age int8,
) *userDomain {
	return &userDomain{
		email:     email,
		password:  password,
		firstName: first_name,
		lastName:  last_name,
		age:       age,
		isActive:  true,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}
func NewUserLoginDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
	}
}

func NewUserUpdateDomain(
	first_name string,
	last_name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		firstName: first_name,
		lastName:  last_name,
		age:       age,
		updatedAt: time.Now(),
	}
}

func NewTokenUserDomain(
	id string,
	email string,
	first_name string,
	last_name string,
	age int8,
) *userDomain {
	return &userDomain{
		id:        id,
		email:     email,
		firstName: first_name,
		lastName:  last_name,
		age:       age,
	}
}

func (ud *userDomain) GetID() string {
	return ud.id
}
func (ud *userDomain) SetID(id string) {
	ud.id = id
}
func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetIsActive() bool {
	return ud.isActive
}
func (ud *userDomain) SetIsActive(isActive bool) {
	ud.isActive = isActive
}
func (ud *userDomain) GetCreatedAt() time.Time {
	return ud.createdAt
}
func (ud *userDomain) SetCreatedAt(createdAt time.Time) {
	ud.createdAt = createdAt
}
func (ud *userDomain) GetUpdatedAt() time.Time {
	return ud.updatedAt
}
func (ud *userDomain) SetUpdatedAt(updatedAt time.Time) {
	ud.updatedAt = updatedAt
}
func (ud *userDomain) GetFirstName() string {
	return ud.firstName
}
func (ud *userDomain) GetLastName() string {
	return ud.lastName
}
func (ud *userDomain) GetAge() int8 {
	return ud.age
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}
func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}
