package gui

import (
	"GophKeeperClient/internal/entity/card"
	"GophKeeperClient/internal/entity/file"
	"GophKeeperClient/internal/entity/pair"
)

type config interface {
	GetClientUserLogin() string
}

type userClientHTTP interface {
	RegisterUser(login, password string) error
}

type tokenClientHTTP interface {
	UpdateToken(clientName, login, password string) error
	BlockToken(clientName string) error
	GetAll() ([]string, error)
}

type cardService interface {
	SaveLocal(c card.Card, key string) error
	SaveRemote(c card.Card, key string) error
	GetFromLocal(name, key string) (card.Card, error)
	GetFromRemote(name, key string) (card.Card, error)
	GetAllFromLocal() ([]card.Card, error)
	GetAllFromRemote() ([]card.Card, error)
	DeleteFromLocal(name string) error
	DeleteFromRemote(name string) error
}

type passwordService interface {
	SaveLocal(p pair.Pair, key string) error
	SaveRemote(p pair.Pair, key string) error
	GetFromLocal(name, key string) (pair.Pair, error)
	GetFromRemote(name, key string) (pair.Pair, error)
	GetAllFromLocal() ([]pair.Pair, error)
	GetAllFromRemote() ([]pair.Pair, error)
	DeleteFromLocal(name string) error
	DeleteFromRemote(name string) error
}

type fileService interface {
	SaveLocal(f *file.File) error
	SaveRemote(f *file.File) error
	GetFromLocal(fileName string) (*file.File, error)
	GetFromRemote(fileName string) (*file.File, error)
	GetAllFromLocal() ([]*file.File, error)
	GetAllFromRemote() ([]*file.File, error)
	DeleteFromLocal(fileName string) error
	DeleteFromRemote(fileName string) error
}
