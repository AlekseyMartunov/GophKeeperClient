package cardclienthttp

import (
	"GophKeeperClient/internal/entity/card"
	httpErrors "GophKeeperClient/internal/errors/http"
	"encoding/json"
	"fmt"
	"net/http"
)

func (cc *CardClientHTTP) Get(cardName string) (card.Card, error) {
	req, err := cc.client.R().
		SetHeader("Authorization", cc.config.Token()).
		Get(fmt.Sprintf("%s/%s/%s", cc.config.ServerADDR(), cc.config.CardURL(), cardName))

	if err != nil {
		return card.Card{}, err
	}

	if req.StatusCode() != http.StatusOK {
		cc.log.Info(string(req.Body()))

		switch req.StatusCode() {
		case http.StatusInternalServerError:
			return card.Card{}, httpErrors.ErrInternalServer

		case http.StatusUnauthorized:
			return card.Card{}, httpErrors.ErrUnauthorized

		case http.StatusBadRequest:
			return card.Card{}, httpErrors.ErrInvalidRequest
		}
	}

	d := dto{}

	err = json.Unmarshal(req.Body(), &d)
	if err != nil {
		cc.log.Error(err)
		return card.Card{}, err
	}

	return d.toEntity(), nil

}
