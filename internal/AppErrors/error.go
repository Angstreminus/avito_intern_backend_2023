package apperrors

import "net/http"

type ResponseError struct {
	Message string `json:"message"`
	Status  int    `json:"-"`
}

type InvalidDataErr struct {
	Message string
}

func (error *InvalidDataErr) Error() string {
	return error.Message
}

type DBoperationErr struct {
	Message string
}

func (error *DBoperationErr) Error() string {
	return error.Message
}

type TransactionErr struct {
	Message string
}

func (error *TransactionErr) Error() string {
	return error.Message
}

type AppError interface {
	Error() string
}

func MatchError(appErr AppError) *ResponseError {
	switch ae := appErr.(type) {
	case *DBoperationErr:
		return &ResponseError{
			Message: ae.Message,
			Status:  http.StatusInternalServerError,
		}
	case *TransactionErr:
		return &ResponseError{
			Message: ae.Message,
			Status:  http.StatusInternalServerError,
		}
	case *InvalidDataErr:
		return &ResponseError{
			Message: ae.Message,
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
