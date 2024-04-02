package tokenclienthttp

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

const tokenURL = "users/login"

type config interface {
	UpdateToken(t string)
	ServerADDR() string
}

type logger interface {
	Info(message string)
	Error(err error)
}

type tokenClientDTO struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
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

func (tc *TokenClientHTTP) UpdateToken(clientName, login, password string) error {
	dto := tokenClientDTO{
		Name:     clientName,
		Login:    login,
		Password: password,
	}

	b, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	req, err := tc.client.R().
		SetBody(b).
		Post(fmt.Sprintf("%s/%s", tc.cfg.ServerADDR(), tokenURL))

	if err != nil {
		tc.log.Error(err)
		return err
	}

	token := tokenResponseDTO{}

	err = json.Unmarshal(req.Body(), &token)
	if err != nil {
		return err
	}

	tc.cfg.UpdateToken(token.Token)
	return nil
}
