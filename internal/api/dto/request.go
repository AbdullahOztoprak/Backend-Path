package dto

type UserRegistrationRequest struct {
    Username     string `json:"username" validate:"required,min=3,max=30"`
    Email        string `json:"email" validate:"required,email"`
    Password     string `json:"password" validate:"required,min=8"`
    Role         string `json:"role" validate:"required,oneof=user admin"`
}

type UserLoginRequest struct {
    Username string `json:"username" validate:"required"`
    Password string `json:"password" validate:"required"`
}

type FundTransferRequest struct {
    FromUserID    int     `json:"from_user_id" validate:"required"`
    ToUserID      int     `json:"to_user_id" validate:"required"`
    Amount        float64 `json:"amount" validate:"required,min=0.01"`
    Description   string  `json:"description" validate:"max=255"`
}

type RefreshTokenRequest struct {
    Token string `json:"token" validate:"required"`
}