package pairstorage

import "github.com/jackc/pgx/v5/pgxpool"

type PairStorage struct {
	pool *pgxpool.Pool
}

func NewPairStorage(p *pgxpool.Pool) *PairStorage {
	return &PairStorage{pool: p}
}
