package carderrors

import "errors"

var (
	ErrCardAlreadyExists = errors.New("card with this name already exists")
	ErrCardDoseNotExist  = errors.New("card dose not exists")
)
