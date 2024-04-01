package cardclienthttp

import (
	"GophKeeperClient/internal/entity/card"
	httpErrors "GophKeeperClient/internal/errors/http"
	"encoding/json"
	"fmt"
	"net/http"
)

func (cc *CardClientHTTP) GetAll() ([]card.Card, error) {
	req, err := cc.client.R().
		SetHeader("Authorization", cc.config.Token()).
		Get(fmt.Sprintf("%s/%s", cc.config.ServerADDR(), cc.config.CardURL()))

	if err != nil {
		return nil, err
	}

	if req.StatusCode() != http.StatusOK {
		cc.log.Info(string(req.Body()))

		switch req.StatusCode() {
		case http.StatusInternalServerError:
			return nil, httpErrors.ErrInternalServer

		case http.StatusUnauthorized:
			return nil, httpErrors.ErrUnauthorized

		case http.StatusBadRequest:
			return nil, httpErrors.ErrInvalidRequest
		}
	}

	dtoCards := make([]dto, 0, 5)
	err = json.Unmarshal(req.Body(), &dtoCards)
	if err != nil {
		return nil, err
	}

	cards := make([]card.Card, 0, 5)

	for _, c := range dtoCards {
		cards = append(cards, c.toEntity())
	}

	return cards, nil

}
