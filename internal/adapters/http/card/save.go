package cardclienthttp

import (
	"GophKeeperClient/internal/entity/card"
	httpErrors "GophKeeperClient/internal/errors/http"
	"encoding/json"
	"fmt"
	"net/http"
)

func (cc *CardClientHTTP) Save(c card.Card) error {
	d := dto{}
	d.fromEntity(c)

	b, err := json.Marshal(d)
	if err != nil {
		cc.log.Error(err)
		return err
	}

	req, err := cc.client.R().
		SetHeader("Authorization", cc.config.Token()).
		SetBody(b).
		Post(fmt.Sprintf("%s/%s", cc.config.ServerADDR(), cardURL))

	if err != nil {
		cc.log.Error(err)
		return err
	}

	if req.StatusCode() != http.StatusOK {
		cc.log.Info(string(req.Body()))

		switch req.StatusCode() {
		case http.StatusInternalServerError:
			return httpErrors.ErrInternalServer

		case http.StatusUnauthorized:
			return httpErrors.ErrUnauthorized

		case http.StatusBadRequest:
			return httpErrors.ErrInvalidRequest
		}
	}

	return nil
}
