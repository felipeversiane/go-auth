package request

type UserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6,containsany=!@&*%$#"`
	FirstName string `json:"first_name" binding:"min=4,max=50"`
	LastName  string `json:"last_name" binding:"min=4,max=50"`
	Age       int8   `json:"age" binding:"required,numeric,min=12,max=140"`
}

type UserUpdateRequest struct {
	FirstName string `json:"first_name" binding:"required,min=4,max=100"`
	LastName  string `json:"last_name" binding:"required,min=4,max=100"`
	Age       int8   `json:"age" binding:"required,min=1,max=140"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email" example:"test@test.com"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*" example:"password#@#@!2121"`
}
