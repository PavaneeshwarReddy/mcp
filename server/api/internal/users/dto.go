package users

/*
Module: Users DTO
Usage: Defines request and response bodies for user requests
*/

type RegisterUserRequest struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Email    *string `json:"email,omitempty"`
	Age      uint    `json:"age"`
}

type RegisterUserResponse struct{}
