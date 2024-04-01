package cardstorage

import "github.com/jackc/pgx/v5/pgxpool"

type CardStorage struct {
	pool *pgxpool.Pool
}

func NewCardStorage(p *pgxpool.Pool) *CardStorage {
	return &CardStorage{pool: p}
}
