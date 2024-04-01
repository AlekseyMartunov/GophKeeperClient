package pairservice

import (
	"GophKeeperClient/internal/entity/pair"
	"context"
)

type pairStorage interface {
	Save(ctx context.Context, p pair.Pair) error
	Get(ctx context.Context, Name string) (pair.Pair, error)
	GetAll(ctx context.Context) ([]pair.Pair, error)
	Delete(ctx context.Context, name string) error
}

type pairHTTPClient interface {
	Save(p pair.Pair) error
	Get(Name string) (pair.Pair, error)
	GetAll() ([]pair.Pair, error)
	Delete(name string) error
}

type encrypter interface {
	EncryptString(text, key string) (string, error)
	DecryptString(text, key string) (string, error)
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

	err = ps.client.Save(p)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PairService) GetFromLocal(name, key string) (pair.Pair, error) {
	p, err := ps.repo.Get(context.Background(), name)
	if err != nil {
		return p, err
	}

	decryptedPair, err := ps.decrypt(p, key)
	if err != nil {
		return p, err
	}

	return decryptedPair, nil
}

func (ps *PairService) GetFromRemote(name, key string) (pair.Pair, error) {
	p, err := ps.client.Get(name)
	if err != nil {
		return p, err
	}

	p, err = ps.decrypt(p, key)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (ps *PairService) GetAllFromLocal() ([]pair.Pair, error) {
	return ps.repo.GetAll(context.Background())
}

func (ps *PairService) GetAllFromRemote() ([]pair.Pair, error) {
	return ps.client.GetAll()
}

func (ps *PairService) DeleteFromLocal(name string) error {
	return ps.repo.Delete(context.Background(), name)
}

func (ps *PairService) DeleteFromRemote(name string) error {
	return ps.client.Delete(name)
}

func (ps *PairService) encrypt(p pair.Pair, key string) (pair.Pair, error) {
	password, err := ps.crypto.EncryptString(p.Password, key)
	if err != nil {
		return p, err
	}

	login, err := ps.crypto.EncryptString(p.Login, key)
	if err != nil {
		return p, err
	}

	p.Password = password
	p.Login = login
	return p, nil
}

func (ps *PairService) decrypt(p pair.Pair, key string) (pair.Pair, error) {
	password, err := ps.crypto.DecryptString(p.Password, key)
	if err != nil {
		return p, err
	}

	login, err := ps.crypto.DecryptString(p.Login, key)
	if err != nil {
		return p, err
	}

	p.Password = password
	p.Login = login

	return p, nil
}
