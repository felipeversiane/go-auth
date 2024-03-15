package converter

import (
	domain "github.com/felipeversiane/golang-auth/internal/account"
	"github.com/felipeversiane/golang-auth/internal/account/repository/entity"
)

func ConvertDomainToEntity(
	domain domain.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:     domain.GetEmail(),
		Password:  domain.GetPassword(),
		FirstName: domain.GetFirstName(),
		LastName:  domain.GetLastName(),
		Age:       domain.GetAge(),
	}
}
