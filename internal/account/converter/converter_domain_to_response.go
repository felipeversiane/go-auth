package converter

import (
	domain "github.com/felipeversiane/go-auth/internal/account"
	"github.com/felipeversiane/go-auth/internal/account/handler/entity/response"
)

func ConvertDomainToResponse(
	userDomain domain.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:        userDomain.GetID(),
		Email:     userDomain.GetEmail(),
		FirstName: userDomain.GetFirstName(),
		LastName:  userDomain.GetLastName(),
		Age:       userDomain.GetAge(),
	}
}
