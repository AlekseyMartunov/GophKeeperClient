package fileclienthttp

import (
	"GophKeeperClient/internal/entity/file"
	httpErrors "GophKeeperClient/internal/errors/http"
	"encoding/json"
	"fmt"
	"net/http"
)

func (fc *FileClientHTTP) Get(fileName string) (*file.File, error) {
	req, err := fc.client.R().
		SetHeader("Authorization", fc.config.Token()).
		Get(fmt.Sprintf("%s/%s/%s", fc.config.ServerADDR(), fileURL, fileName))

	if err != nil {
		fc.log.Error(err)
		return nil, err
	}

	if req.StatusCode() != http.StatusOK {
		fc.log.Info(string(req.Body()))

		switch req.StatusCode() {
		case http.StatusInternalServerError:
			return nil, httpErrors.ErrInternalServer

		case http.StatusUnauthorized:
			return nil, httpErrors.ErrUnauthorized

		case http.StatusBadRequest:
			return nil, httpErrors.ErrInvalidRequest
		}
	}
	dto := fileDTO{}
	err = json.Unmarshal(req.Body(), &dto)
	if err != nil {
		return nil, err
	}

	return dto.ToEntity(), nil

}
