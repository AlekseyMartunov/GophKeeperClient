package fileserrors

import "errors"

var (
	ErrFailAlreadyExists = errors.New("file with this name already exists")
	ErrFailDoseNotExist  = errors.New("file dose not exists")
)
