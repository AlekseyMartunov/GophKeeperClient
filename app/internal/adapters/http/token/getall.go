package tokenclienthttp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (tc *TokenClientHTTP) GetAll() ([]string, error) {
	req, err := tc.client.R().
		SetHeader("Authorization", tc.cfg.Token()).
		Get(fmt.Sprintf("%s/%s", tc.cfg.ServerADDR(), tokenURL))

	if err != nil {
		return nil, err
	}

	if req.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%s:%s:%s", req.String(), tc.cfg.Token(), "token"))
	}

	dto := allClients{}

	err = json.Unmarshal(req.Body(), &dto)
	if err != nil {
		return nil, err
	}

	return dto.Clients, nil

}
