package pairclienthttp

import (
	httpErrors "GophKeeperClient/internal/errors/http"
	"fmt"
	"net/http"
)

func (pc *PairClientHTTP) Delete(pairName string) error {
	req, err := pc.client.R().
		SetHeader("Authorization", pc.config.Token()).
		Delete(fmt.Sprintf("%s/%s/%s", pc.config.ServerADDR(), pairURL, pairName))

	if err != nil {
		pc.log.Error(err)
		return err
	}

	if req.StatusCode() != http.StatusNoContent {
		pc.log.Info(string(req.Body()))

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
