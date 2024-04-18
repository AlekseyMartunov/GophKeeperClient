package tokenclienthttp

import (
	"github.com/go-resty/resty/v2"
)

const tokenURL = "token"

type config interface {
	UpdateToken(t string)
	ServerADDR() string
	Token() string
}

type logger interface {
	Info(message string)
	Error(err error)
}

type tokenResponseDTO struct {
	Token string `json:"token"`
}

type TokenClientHTTP struct {
	cfg    config
	client *resty.Client
	log    logger
}

func NewTokenClientHTTP(c config, l logger) *TokenClientHTTP {
	return &TokenClientHTTP{
		cfg:    c,
		client: resty.New(),
		log:    l,
	}
}
