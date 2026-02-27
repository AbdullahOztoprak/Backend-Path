package fixtures

import "github.com/AbdullahOztoprak/Backend-Path/internal/domain/entity"

var Users = []entity.User{
    {
        ID:       1,
        Username: "john_doe",
        Email:    "john@example.com",
        Password: "hashed_password_1", // Use a hashed password for testing
        Role:     "user",
    },
    {
        ID:       2,
        Username: "jane_doe",
        Email:    "jane@example.com",
        Password: "hashed_password_2", // Use a hashed password for testing
        Role:     "admin",
    },
    {
        ID:       3,
        Username: "admin_user",
        Email:    "admin@example.com",
        Password: "hashed_password_3", // Use a hashed password for testing
        Role:     "super_admin",
    },
}