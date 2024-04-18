package cardclienthttp

import "github.com/go-resty/resty/v2"

const cardURL = "card"

type config interface {
	ServerADDR() string
	Token() string
}

type logger interface {
	Info(message string)
	Error(err error)
}

type CardClientHTTP struct {
	client *resty.Client
	config config
	log    logger
}

func NewCardClientHTTP(c config, l logger) *CardClientHTTP {
	return &CardClientHTTP{
		client: resty.New(),
		config: c,
		log:    l,
	}
}
