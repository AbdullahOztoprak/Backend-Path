package apperror

import "fmt"

// AppError represents a custom error type for the application.
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// New creates a new AppError with the given code and message.
func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// Error implements the error interface for AppError.
func (e *AppError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// Is checks if the error is of a specific type.
func Is(err error, target *AppError) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == target.Code
	}
	return false
}

// Wrap wraps an existing error with a new message and returns an AppError.
func Wrap(err error, message string) *AppError {
	return &AppError{
		Code:    500, // Default to internal server error
		Message: fmt.Sprintf("%s: %v", message, err),
	}
}