package common

import "fmt"

type ErrorCode string

const(
	ErrorCodeValidation ErrorCode = "VALIDATION_ERROR"
	ErrorCodeNotFound ErrorCode = "NOT_FOUND"
	ErrorCodeConflict ErrorCode = "CONFLICT"
	ErrorCodeForbidden ErrorCode = "FORBIDDEN"
	ErrorCodeInternal ErrorCode = "INTERNAL_ERROR"
)

type AppError struct {
	Code ErrorCode
	Message string
	Err error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s : %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewValidationError(Message string, Err error) *AppError {
	return &AppError{Code: ErrorCodeValidation, Message: Message, Err: Err}
}

func NewNotFoundError(Message string, Err error) *AppError {
	return &AppError{Code : ErrorCodeNotFound, Message : Message, Err : Err}
}

func NewConflictError(Message string, Err error) *AppError {
	return &AppError{Code : ErrorCodeConflict, Message : Message, Err : Err}
}

func NewForbiddenError(Message string, Err error) *AppError {
	return &AppError{Code : ErrorCodeForbidden, Message : Message, Err : Err}
}

func NewInternalError(Message string, Err error) *AppError {
	return &AppError{Code : ErrorCodeInternal, Message : Message, Err : Err}
}

func IsValidationError(err error) bool {
	if err == nil {
		return false
	}
	if apperr, ok := err.(*AppError); ok {
		return apperr.Code == ErrorCodeValidation
	}
	return false
}

func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	if apperr, ok := err.(*AppError); ok {
		return apperr.Code == ErrorCodeNotFound
	}
	return false
}

func IsConflictError(err error) bool {
	if err == nil {
		return false
	}
	if apperr, ok := err.(*AppError); ok {
		return apperr.Code == ErrorCodeConflict
	}
	return false
}

func IsForbiddenError(err error) bool {
	if err == nil {
		return false
	}		
	if apperr, ok := err.(*AppError); ok {
		return apperr.Code == ErrorCodeForbidden
	}	
	return false
}

func IsInternalError(err error) bool {
	if err == nil {
		return false
	}
	if apperr, ok := err.(*AppError); ok {
		return apperr.Code == ErrorCodeInternal
	}
	return false
}
