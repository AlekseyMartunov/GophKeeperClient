package tokenclienthttp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

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
	if req.StatusCode() != http.StatusOK {
		return errors.New(string(req.Body()))
	}

	err = json.Unmarshal(req.Body(), &token)
	if err != nil {
		return err
	}

	tc.cfg.UpdateToken(token.Token)

	return nil
}
