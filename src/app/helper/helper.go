package helper

import (
	"fmt"
	
)

// UnauthorizedError represents 401 Unauthorized error.
type UnauthorizedError struct {
	Message string
	Code    int
}

func (e *UnauthorizedError) Error() string {
	return fmt.Sprintf("UnauthorizedError: %s", e.Message)
}

// BadRequestError represents 400 Bad Request error.
type BadRequestError struct {
	Message string
	Code    int
}

func (e *BadRequestError) Error() string {
	return fmt.Sprintf("BadRequestError: %s", e.Message)
}

// ConflictError represents 409 Conflict error.
type ConflictError struct {
	Message string
	Code    int
}

func (e *ConflictError) Error() string {
	return fmt.Sprintf("ConflictError: %s", e.Message)
}

// InternalServerError represents 500 Internal Server Error.
type InternalServerError struct {
	Message string
	Code    int
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("InternalServerError: %s", e.Message)
}

// UnauthenticatedError represents 401 Unauthenticated error.
type UnauthenticatedError struct {
	Message string
	Code    int
}

func (e *UnauthenticatedError) Error() string {
	return fmt.Sprintf("UnauthenticatedError: %s", e.Message)
}

// NotFoundError represents 404 Not Found error.
type NotFoundError struct {
	Message string
	Code    int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("NotFoundError: %s", e.Message)
}

// ValidationError represents 400 Validation Error.
type ValidationError struct {
	Message string
	Code    int
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("ValidationError: %s", e.Message)
}

