package userclienthttp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

const userURL = "users/register"

type config interface {
	ServerADDR() string
	ClientUserLogin(name string)
}

type logger interface {
	Info(message string)
	Error(err error)
}

type UserClientHTTP struct {
	cfg    config
	client *resty.Client
	log    logger
}

type userDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func NewUserClientHTTP(c config, l logger) *UserClientHTTP {
	return &UserClientHTTP{
		cfg:    c,
		client: resty.New(),
		log:    l,
	}
}

func (uc *UserClientHTTP) RegisterUser(login, password string) error {
	u := userDTO{
		Login:    login,
		Password: password,
	}

	b, err := json.Marshal(u)
	if err != nil {
		return err
	}

	req, err := uc.client.R().
		SetBody(b).
		Post(fmt.Sprintf("%s/%s", uc.cfg.ServerADDR(), userURL))

	if err != nil {
		uc.log.Error(err)
		return err
	}

	if req.StatusCode() != http.StatusOK {
		return errors.New(string(req.Body()))
	}
	uc.cfg.ClientUserLogin(login)

	return nil
}
