package apperrors

import (
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func NewErrorResponse(message string, statusCode int) *ErrorResponse {
	return &ErrorResponse{
		Message:    message,
		StatusCode: statusCode,
	}
}

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrNotFound          = errors.New("resource not found")
	ErrInvalidInput      = errors.New("invalid input")
	ErrInternalServer    = errors.New("internal server error")
	ErrUnauthorized      = errors.New("user unauthorized")
)

func NewUserAlreadyExistsError() *ErrorResponse {
	return &ErrorResponse{
		Message:    "user already exists",
		StatusCode: http.StatusConflict,
	}
}

func NewNotFoundError(resource string) *ErrorResponse {
	return &ErrorResponse{
		Message:    resource + " not found",
		StatusCode: http.StatusNotFound,
	}
}

func NewBadRequestError() *ErrorResponse {
	return &ErrorResponse{
		Message:    "invalid input",
		StatusCode: http.StatusBadRequest,
	}
}

func NewInternalServerError() *ErrorResponse {
	return &ErrorResponse{
		Message:    "internal server error",
		StatusCode: http.StatusInternalServerError, // 500
	}
}

func NewUnauthorizedError() *ErrorResponse {
	return &ErrorResponse{
		Message:    "user unauthorized",
		StatusCode: http.StatusUnauthorized, // 401
	}
}
func NewAuthenticationError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message:    message,
		StatusCode: http.StatusUnauthorized, // 401
	}
}

func NewResourceAlreadyExistsError(resource string) *ErrorResponse {
	return &ErrorResponse{
		Message:    resource + " already exists",
		StatusCode: http.StatusConflict,
	}
}

func NewResourceNotAvailableError(resource string) *ErrorResponse {
	return &ErrorResponse{
		Message: resource + " not available",
		StatusCode: http.StatusNotFound,
	}
}

func NewCustomError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		StatusCode: http.StatusNotAcceptable,
	}
}
