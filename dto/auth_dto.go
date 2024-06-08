package dto

type LoginResponse struct {
	Username string
	Token    string
}

type LoginRequest struct{
	Username string
	Password string
}
