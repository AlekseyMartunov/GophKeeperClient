package pairclienthttp

import (
	"GophKeeperClient/internal/entity/pair"
	httpErrors "GophKeeperClient/internal/errors/http"
	"encoding/json"
	"fmt"
	"net/http"
)

func (pc *PairClientHTTP) Get(pairName string) (pair.Pair, error) {
	req, err := pc.client.R().
		SetHeader("Authorization", pc.config.Token()).
		Get(fmt.Sprintf("%s/%s/%s", pc.config.ServerADDR(), pairURL, pairName))

	if err != nil {
		return pair.Pair{}, err
	}

	if req.StatusCode() != http.StatusOK {
		pc.log.Info(string(req.Body()))

		switch req.StatusCode() {
		case http.StatusInternalServerError:
			return pair.Pair{}, httpErrors.ErrInternalServer

		case http.StatusUnauthorized:
			return pair.Pair{}, httpErrors.ErrUnauthorized

		case http.StatusBadRequest:
			return pair.Pair{}, httpErrors.ErrInvalidRequest
		}
	}

	d := dto{}

	err = json.Unmarshal(req.Body(), &d)
	if err != nil {
		pc.log.Error(err)
		return pair.Pair{}, err
	}

	return d.toEntity(), nil

}
