package models

import (
    "errors"
    "regexp"
    "strings"
    "time"
)

type User struct {
    ID           int64     `json:"id"`
    Username     string    `json:"username"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"password_hash"`
    Role         string    `json:"role"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

// Validate checks if the user fields are valid
func (u *User) Validate() error {
    if len(u.Username) < 3 {
        return errors.New("username must be at least 3 characters")
    }
    if !strings.Contains(u.Email, "@") {
        return errors.New("invalid email address")
    }
    if u.Role == "" {
        return errors.New("role is required")
    }
    // Optional: Email regex check
    emailRegex := regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)
    if !emailRegex.MatchString(u.Email) {
        return errors.New("invalid email format")
    }
    return nil
}