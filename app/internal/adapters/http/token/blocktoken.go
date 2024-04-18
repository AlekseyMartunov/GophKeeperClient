package tokenclienthttp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (tc *TokenClientHTTP) BlockToken(clientName string) error {
	dto := blockToken{TokenName: clientName}

	b, err := json.Marshal(dto)
	if err != nil {
		return err
	}

	req, err := tc.client.R().
		SetHeader("Authorization", tc.cfg.Token()).
		SetBody(b).
		Delete(fmt.Sprintf("%s/%s", tc.cfg.ServerADDR(), tokenURL))

	if req.StatusCode() != http.StatusOK {
		return errors.New(string(req.Body()))
	}

	return nil
}
