package pairclienthttp

import (
	"GophKeeperClient/internal/entity/pair"
	httpErrors "GophKeeperClient/internal/errors/http"
	"encoding/json"
	"fmt"
	"net/http"
)

func (pc *PairClientHTTP) Save(p pair.Pair) error {
	d := dto{}
	d.fromEntity(p)

	b, err := json.Marshal(d)
	if err != nil {
		pc.log.Error(err)
		return err
	}

	req, err := pc.client.R().
		SetHeader("Authorization", pc.config.Token()).
		SetBody(b).
		Post(fmt.Sprintf("%s/%s", pc.config.ServerADDR(), pc.config.PairURL()))
	
	if err != nil {
		pc.log.Error(err)
		return err
	}

	if req.StatusCode() != http.StatusOK {
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
