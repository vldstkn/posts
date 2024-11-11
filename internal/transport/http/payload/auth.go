package payload

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	Id          int    `json:"id"`
	AccessToken string `json:"access_token"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Id          int    `json:"id"`
	AccessToken string `json:"access_token"`
}

type GetNewTokensResponse struct {
	AccessToken string `json:"access_token"`
}
