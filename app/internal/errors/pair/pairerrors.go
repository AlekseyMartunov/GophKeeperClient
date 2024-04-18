package pairerrors

import "errors"

var (
	ErrPairAlreadyExists = errors.New("pair with this name already exists")
	ErrPairDoseNotExist  = errors.New("pairs dose not exists")
)
