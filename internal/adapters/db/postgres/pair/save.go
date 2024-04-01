package pairstorage

import (
	"GophKeeperClient/internal/entity/pair"
	pairErrors "GophKeeperClient/internal/errors/pair"
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (ps *PairStorage) Save(ctx context.Context, p pair.Pair) error {
	query := `INSERT INTO pairs (pair_name, password, login, created_time) 
				VALUES ($1, $2, $3, $4)`

	_, err := ps.pool.Exec(ctx, query, p.Name, p.Password, p.Login, p.CreatedTime)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return pairErrors.ErrPairAlreadyExists
		}
		return err
	}
	return nil
}
