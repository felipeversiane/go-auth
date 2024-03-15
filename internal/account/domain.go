package domain

import (
	"crypto/md5"
	"encoding/hex"
)

type userDomain struct {
	id         string
	email      string
	password   string
	firstName  string
	lastName   string
	age        int8
	isActive   bool
	created_at string
	updated_at string
}

type UserDomainInterface interface {
	SetID(id string)
	GetID() string
	GetEmail() string
	GetIsActive() bool
	SetIsActive(isActive bool)
	GetCreatedAt() string
	SetCreatedAt(created_at string)
	GetUpdatedAt() string
	SetUpdatedAt(updated_at string)
	GetFirstName() string
	GetLastName() string
	GetPassword() string
	GetAge() int8
	EncryptPassword()
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
func (ud *userDomain) GetCreatedAt() string {
	return ud.created_at
}
func (ud *userDomain) SetCreatedAt(created_at string) {
	ud.created_at = created_at
}
func (ud *userDomain) GetUpdatedAt() string {
	return ud.updated_at
}
func (ud *userDomain) SetUpdatedAt(updated_at string) {
	ud.updated_at = updated_at
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
