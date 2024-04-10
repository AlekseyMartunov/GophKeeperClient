package fileclienthttp

import (
	"GophKeeperClient/internal/entity/file"
	httpErrors "GophKeeperClient/internal/errors/http"
	"encoding/json"
	"fmt"
	"net/http"
)

func (fc *FileClientHTTP) GetAll() ([]*file.File, error) {
	req, err := fc.client.R().
		SetHeader("Authorization", fc.config.Token()).
		Get(fmt.Sprintf("%s/%s", fc.config.ServerADDR(), fileURL))

	if err != nil {
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

	dtoArr := make([]fileDTO, 0, 10)

	err = json.Unmarshal(req.Body(), &dtoArr)
	if err != nil {
		return nil, err
	}

	filesArr := make([]*file.File, 0, len(dtoArr))
	for _, dto := range dtoArr {
		filesArr = append(filesArr, dto.ToEntity())
	}

	return filesArr, err
}
