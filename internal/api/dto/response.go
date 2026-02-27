package dto

type UserResponse struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Role     string `json:"role"`
}

type TransactionResponse struct {
    ID          int     `json:"id"`
    FromUserID  int     `json:"from_user_id"`
    ToUserID    int     `json:"to_user_id"`
    Amount      float64 `json:"amount"`
    Description string  `json:"description"`
    CreatedAt   string  `json:"created_at"`
}

type BalanceResponse struct {
    UserID int     `json:"user_id"`
    Balance float64 `json:"balance"`
}

type HealthCheckResponse struct {
    Status string `json:"status"`
}