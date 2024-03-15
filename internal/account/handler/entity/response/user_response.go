package response

type UserResponse struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int8   `json:"age"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}
