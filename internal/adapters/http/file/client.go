package fileclienthttp

import "github.com/go-resty/resty/v2"

const (
	fileURL = "file"
)

type config interface {
	ServerADDR() string
	Token() string
}

type logger interface {
	Info(message string)
	Error(err error)
}

type FileClientHTTP struct {
	client *resty.Client
	config config
	log    logger
}

func NewFileClientHTTP(c config, l logger) *FileClientHTTP {
	return &FileClientHTTP{
		client: resty.New(),
		config: c,
		log:    l,
	}
}
