package cardstorage

import (
	"GophKeeperClient/internal/entity/card"
	cardErrors "GophKeeperClient/internal/errors/card"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

func (cs *CardStorage) Get(ctx context.Context, name string) (card.Card, error) {
	query := `SELECT card_name, card_number, owner, cvv, card_date, created_time FROM cards WHERE card_name = $1`

	row := cs.pool.QueryRow(ctx, query, name)

	var c card.Card
	err := row.Scan(&c.Name, &c.Number, &c.Owner, &c.CVV, &c.Date, &c.CreatedTime)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c, cardErrors.ErrCardDoseNotExist
		}
		return c, err
	}

	return c, nil
}
