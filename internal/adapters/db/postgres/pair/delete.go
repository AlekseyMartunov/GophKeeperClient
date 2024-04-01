package pairstorage

import "context"

func (ps *PairStorage) Delete(ctx context.Context, name string) error {
	query := `DELETE FROM pairs WHERE pair_name = $1`

	_, err := ps.pool.Exec(ctx, query, name)
	if err != nil {
		return err
	}
	return nil
}
