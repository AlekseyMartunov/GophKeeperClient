package cardstorage

import (
	"GophKeeperClient/internal/entity/card"
	cardErrors "GophKeeperClient/internal/errors/card"
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (cs *CardStorage) Save(ctx context.Context, c card.Card) error {
	query := `INSERT INTO cards (card_name, card_number, owner, cvv, card_date, created_time) 
				VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := cs.pool.Exec(ctx, query, c.Name, c.Number, c.Owner, c.CVV, c.Date, c.CreatedTime)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return cardErrors.ErrCardAlreadyExists
		}
		return err
	}
	return nil
}
