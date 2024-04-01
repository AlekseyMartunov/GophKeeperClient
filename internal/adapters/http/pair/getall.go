package pairclienthttp

import (
	"GophKeeperClient/internal/entity/pair"
	httpErrors "GophKeeperClient/internal/errors/http"
	"encoding/json"
	"fmt"
	"net/http"
)

func (pc *PairClientHTTP) GetAll() ([]pair.Pair, error) {
	req, err := pc.client.R().
		SetHeader("Authorization", pc.config.Token()).
		Get(fmt.Sprintf("%s/%s", pc.config.ServerADDR(), pc.config.PairURL()))

	if err != nil {
		return nil, err
	}

	if req.StatusCode() != http.StatusOK {
		pc.log.Info(string(req.Body()))

		switch req.StatusCode() {
		case http.StatusInternalServerError:
			return nil, httpErrors.ErrInternalServer

		case http.StatusUnauthorized:
			return nil, httpErrors.ErrUnauthorized

		case http.StatusBadRequest:
			return nil, httpErrors.ErrInvalidRequest
		}
	}

	dtoPairs := make([]dto, 0, 5)
	err = json.Unmarshal(req.Body(), &dtoPairs)
	if err != nil {
		return nil, err
	}

	pairs := make([]pair.Pair, 0, 5)

	for _, p := range dtoPairs {
		pairs = append(pairs, p.toEntity())
	}

	return pairs, nil

}
