package models

type CreateUserRequest struct {
	User User
}

type CreateUserResponse struct {
	Token string `json:"token"`
	Err   error  `json:"error,omitempty"`
}

type LoginRequest struct {
	User User
}

type LoginResponse struct {
	Token string `json:"token"`
	Err   error  `json:"error,omitempty"`
}
