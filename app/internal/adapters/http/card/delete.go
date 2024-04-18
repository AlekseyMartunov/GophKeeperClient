package cardclienthttp

import (
	httpErrors "GophKeeperClient/internal/errors/http"
	"fmt"
	"net/http"
)

func (cc *CardClientHTTP) Delete(cardName string) error {
	req, err := cc.client.R().
		SetHeader("Authorization", cc.config.Token()).
		Delete(fmt.Sprintf("%s/%s/%s", cc.config.ServerADDR(), cardURL, cardName))

	if err != nil {
		cc.log.Error(err)
		return err
	}

	if req.StatusCode() != http.StatusNoContent {
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
