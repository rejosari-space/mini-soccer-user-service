package error

import "errors"

var (
	InternalServerError = errors.New("internal server error")
	ErrSqlErr           = errors.New("database server error to execute query")
	ErrTooManyRequests  = errors.New("too many requests")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrInvalidToken     = errors.New("invalid token")
	ErrForbidden        = errors.New("forbidden")
)

var GeneralErrors = []error{
	InternalServerError,
	ErrSqlErr,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}
