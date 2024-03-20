package pairservice

import (
	"GophKeeperClient/internal/entity/pair"
	"context"
)

type pairStorage interface {
	Save(ctx context.Context, p pair.Pair) error
	Get(ctx context.Context, Name string) (pair.Pair, error)
	GetAll() ([]pair.Pair, error)
}

type pairHTTPClient interface {
	Send(p pair.Pair) error
	Get(Name string) (pair.Pair, error)
	GetAll() ([]pair.Pair, error)
}

type encrypter interface {
	Encrypt(text, key string) (string, error)
	Decrypt(text, key string) (string, error)
}

type PairService struct {
	repo   pairStorage
	client pairHTTPClient
	crypto encrypter
}

func NewPairService(r pairStorage, c pairHTTPClient, e encrypter) *PairService {
	return &PairService{
		repo:   r,
		client: c,
		crypto: e,
	}
}

func (ps *PairService) SaveLocal(p pair.Pair, key string) error {
	p, err := ps.encrypt(p, key)
	if err != nil {
		return err
	}

	err = ps.repo.Save(context.Background(), p)
	if err != nil {
		return err
	}

	return nil
}

func (ps *PairService) SaveRemote(p pair.Pair, key string) error {
	p, err := ps.encrypt(p, key)
	if err != nil {
		return err
	}

	err = ps.client.Send(p)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PairService) encrypt(p pair.Pair, key string) (pair.Pair, error) {
	password, err := ps.crypto.Encrypt(p.Password, key)
	if err != nil {
		return p, err
	}

	login, err := ps.crypto.Encrypt(p.Login, key)
	if err != nil {
		return p, err
	}

	p.Password = password
	p.Login = login
	return p, nil
}

func (ps *PairService) decrypt(p pair.Pair, key string) (pair.Pair, error) {
	password, err := ps.crypto.Decrypt(p.Password, key)
	if err != nil {
		return p, err
	}

	login, err := ps.crypto.Decrypt(p.Login, key)
	if err != nil {
		return p, err
	}

	p.Password = password
	p.Login = login
	return p, nil
}
