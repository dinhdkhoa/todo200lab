package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int `json:"status_code"`

	RootErr error `json:"-"`

	Message string `json:"message"`

	Log string `json:"log"`

	Key string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {

	return &AppError{

		StatusCode: http.StatusBadRequest,

		RootErr: root,

		Message: msg,

		Log: log,

		Key: key,
	}

}

func NewUnauthorizedError(root error, msg, log, key string) *AppError {

	return &AppError{

		StatusCode: http.StatusUnauthorized,

		RootErr: root,

		Message: msg,

		Log: log,

		Key: key,
	}

}

func NewCustomError(root error, msg, log, key string) *AppError {

	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {

	if err, ok := e.RootErr.(*AppError); ok {

		return err.RootError()

	}

	return e.RootErr

}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

// 15:24
func ErrDB(err error) *AppError {
	return NewErrorResponse(err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInternal(err error) *AppError {
	return NewErrorResponse(err, "internal error", err.Error(), "INTERNAL")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		"cannot list "+entity,
		err.Error(),
		"ErrCannotListEntity",
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		"cannot delete "+entity,
		err.Error(),
		"ErrCannotDeleteEntity",
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		"cannot get "+entity,
		err.Error(),
		"ErrCannotGetEntity",
	)
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		entity+" already exists",
		err.Error(),
		"ErrEntityExisted",
	)
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		entity+" not found",
		err.Error(),
		"ErrEntityNotFound",
	)
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		"cannot create "+entity,
		err.Error(),
		"ErrCannotCreateEntity",
	)
}

func ErrorNoPermission() *AppError {
	return NewCustomError(
		nil,
		"you have no permission",
		"you have no permission",
		"ErrorNoPermission",
	)
}

var ErrorRecord404 = errors.New("record not found")
