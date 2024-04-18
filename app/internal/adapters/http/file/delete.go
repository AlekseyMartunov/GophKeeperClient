package fileclienthttp

import (
	httpErrors "GophKeeperClient/internal/errors/http"
	"fmt"
	"net/http"
)

func (fc *FileClientHTTP) Delete(fileName string) error {
	req, err := fc.client.R().
		SetHeader("Authorization", fc.config.Token()).
		Delete(fmt.Sprintf("%s/%s/%s", fc.config.ServerADDR(), fileURL, fileName))

	if err != nil {
		fc.log.Error(err)
		return err
	}

	if req.StatusCode() != http.StatusNoContent {
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
