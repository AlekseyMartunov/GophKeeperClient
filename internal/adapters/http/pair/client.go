package pairclienthttp

import "github.com/go-resty/resty/v2"

const pairURL = "pairs"

type config interface {
	ServerADDR() string
	Token() string
}

type logger interface {
	Info(message string)
	Error(err error)
}

type PairClientHTTP struct {
	client *resty.Client
	config config
	log    logger
}

func NewPairClientHTTP(c config, l logger) *PairClientHTTP {
	return &PairClientHTTP{
		client: resty.New(),
		config: c,
		log:    l,
	}
}
