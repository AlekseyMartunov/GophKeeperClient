package cardservice

import (
	"GophKeeperClient/internal/entity/card"
	"context"
)

type cardStorage interface {
	Save(ctx context.Context, c card.Card) error
	Get(ctx context.Context, Name string) (card.Card, error)
	GetAll(ctx context.Context) ([]card.Card, error)
	Delete(ctx context.Context, name string) error
}

type cardHTTPClient interface {
	Save(c card.Card) error
	Get(Name string) (card.Card, error)
	GetAll() ([]card.Card, error)
	Delete(name string) error
}

type encrypter interface {
	EncryptString(text, key string) (string, error)
	DecryptString(text, key string) (string, error)
}

type CardService struct {
	repo   cardStorage
	client cardHTTPClient
	crypto encrypter
}

func NewCardService(r cardStorage, c cardHTTPClient, e encrypter) *CardService {
	return &CardService{
		repo:   r,
		client: c,
		crypto: e,
	}
}

func (cs *CardService) SaveLocal(c card.Card, key string) error {
	p, err := cs.encrypt(c, key)
	if err != nil {
		return err
	}

	err = cs.repo.Save(context.Background(), p)
	if err != nil {
		return err
	}

	return nil
}

func (cs *CardService) SaveRemote(c card.Card, key string) error {
	p, err := cs.encrypt(c, key)
	if err != nil {
		return err
	}

	err = cs.client.Save(p)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CardService) GetFromLocal(name, key string) (card.Card, error) {
	p, err := cs.repo.Get(context.Background(), name)
	if err != nil {
		return p, err
	}

	decryptedCard, err := cs.decrypt(p, key)
	if err != nil {
		return p, err
	}

	return decryptedCard, nil
}

func (cs *CardService) GetFromRemote(name, key string) (card.Card, error) {
	c, err := cs.client.Get(name)
	if err != nil {
		return c, err
	}

	c, err = cs.decrypt(c, key)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (cs *CardService) GetAllFromLocal() ([]card.Card, error) {
	return cs.repo.GetAll(context.Background())
}

func (cs *CardService) GetAllFromRemote() ([]card.Card, error) {
	return cs.client.GetAll()
}

func (cs *CardService) DeleteFromLocal(name string) error {
	return cs.repo.Delete(context.Background(), name)
}

func (cs *CardService) DeleteFromRemote(name string) error {
	return cs.client.Delete(name)
}

func (cs *CardService) encrypt(c card.Card, key string) (card.Card, error) {
	number, err := cs.crypto.EncryptString(c.Number, key)
	if err != nil {
		return c, err
	}

	owner, err := cs.crypto.EncryptString(c.Owner, key)
	if err != nil {
		return c, err
	}

	cvv, err := cs.crypto.EncryptString(c.CVV, key)
	if err != nil {
		return c, err
	}

	date, err := cs.crypto.EncryptString(c.Date, key)
	if err != nil {
		return c, err
	}

	c.Number = number
	c.CVV = cvv
	c.Owner = owner
	c.Date = date
	return c, nil
}

func (cs *CardService) decrypt(c card.Card, key string) (card.Card, error) {
	number, err := cs.crypto.DecryptString(c.Number, key)
	if err != nil {
		return c, err
	}

	cvv, err := cs.crypto.DecryptString(c.CVV, key)
	if err != nil {
		return c, err
	}

	owner, err := cs.crypto.DecryptString(c.Owner, key)
	if err != nil {
		return c, err
	}

	date, err := cs.crypto.DecryptString(c.Date, key)
	if err != nil {
		return c, err
	}

	c.Number = number
	c.CVV = cvv
	c.Owner = owner
	c.Date = date

	return c, nil
}
