package cardstorage

import "context"

func (cs *CardStorage) Delete(ctx context.Context, name string) error {
	query := `DELETE FROM cards WHERE card_name = $1`

	_, err := cs.pool.Exec(ctx, query, name)
	if err != nil {
		return err
	}
	return nil
}
