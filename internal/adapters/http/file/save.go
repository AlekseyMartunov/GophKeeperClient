package fileclienthttp

import (
	"GophKeeperClient/internal/entity/file"
	httpErrors "GophKeeperClient/internal/errors/http"
	"encoding/json"
	"fmt"
	"net/http"
)

func (fc *FileClientHTTP) Send(f file.File) error {
	dto := fileDTO{}
	dto.FromEntity(f)

	b, err := json.Marshal(dto)
	if err != nil {
		fc.log.Error(err)
		return err
	}

	req, err := fc.client.R().
		SetHeader("Authorization", fc.config.Token()).
		SetBody(b).
		Post(fmt.Sprintf("%s/%s", fc.config.ServerADDR(), fileURL))

	if err != nil {
		fc.log.Error(err)
		return err
	}

	if req.StatusCode() != http.StatusOK {
		fc.log.Info(string(req.Body()))

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
