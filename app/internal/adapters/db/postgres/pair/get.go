package pairstorage

import (
	"GophKeeperClient/internal/entity/pair"
	pairerrors "GophKeeperClient/internal/errors/pair"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

func (ps *PairStorage) Get(ctx context.Context, name string) (pair.Pair, error) {
	query := `SELECT pair_name, password, login, created_time FROM pairs WHERE pair_name = $1`

	row := ps.pool.QueryRow(ctx, query, name)

	var p pair.Pair
	err := row.Scan(&p.Name, &p.Password, &p.Login, &p.CreatedTime)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return p, pairerrors.ErrPairDoseNotExist
		}
		return p, err
	}

	return p, nil
}
