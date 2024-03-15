package converter

import (
	domain "github.com/felipeversiane/golang-auth/internal/account"
	"github.com/felipeversiane/golang-auth/internal/account/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) domain.UserDomainInterface {
	domain := domain.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.FirstName,
		entity.LastName,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex())
	return domain
}
