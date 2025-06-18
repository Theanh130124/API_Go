package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`         //loi goc lỗi golang bắt
	Message    string `json:"message"`   //Bao loi cho endUser
	Log        string `json:"log"`       // thong báo Loi lay tu rootErr (loi goc lỗi golang bắt)
	Key        string `json:"error_key"` //Custom key(da ngon ngu)
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}
func NewErrorBadRequest(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}
func NewUnauthorized(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

// Error (Error (Root Error) -> boc nhieu Error
func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError() // Lay rootError (neu no dc boc)
	}
	return e.RootErr // khong dc boc thi lay RootErr luon
}

// Bat ki ham nao Error() se la Error tron golang
func (e *AppError) Error() string {
	return e.RootError().Error()
}

func NewCustomError(root error, msg, log, key string) *AppError {
	if root != nil {
		return NewErrorBadRequest(root, msg, log, key)
	}
	return NewErrorBadRequest(errors.New(msg), msg, log, key)
}

// Loi cua database
func ErrDB(err error) *AppError { //err la loi goc
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something went wrong with db", err.Error(), "DB_ERROR")
}

// Loi parseBody

func ErrInvalidRequest(err error) *AppError {
	return NewErrorBadRequest(err, "Invalid request", err.Error(), "ErrInvalidRequest")
}

// Loi Runtime tren server
func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something went wrong in the server", err.Error(), "ErrInternal")
}

func RecordNotFound(err error) *AppError {
	return NewFullErrorResponse(http.StatusNotFound, err, "Record not found", err.Error(), "ErrRecordNotFound")
}
