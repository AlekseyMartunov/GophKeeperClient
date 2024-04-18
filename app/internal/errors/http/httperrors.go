package httperrors

import "errors"

var (
	ErrInvalidRequest = errors.New("invalid request")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrInternalServer = errors.New("internal error on server")
)
