package cardstorage

import (
	"GophKeeperClient/internal/entity/card"
	cardErrors "GophKeeperClient/internal/errors/card"
	"context"
)

func (cs *CardStorage) GetAll(ctx context.Context) ([]card.Card, error) {
	query := "SELECT card_name, created_time FROM cards"

	rows, err := cs.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	cards := make([]card.Card, 0, 3)

	for rows.Next() {
		var c card.Card

		err = rows.Scan(&c.Name, &c.CreatedTime)
		if err != nil {
			return nil, err
		}

		cards = append(cards, c)

	}

	if rows.Err() != nil {
		return nil, err
	}

	if len(cards) == 0 {
		return nil, cardErrors.ErrCardDoseNotExist
	}

	return cards, nil
}
