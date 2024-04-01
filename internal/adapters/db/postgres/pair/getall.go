package pairstorage

import (
	"GophKeeperClient/internal/entity/pair"
	pairerrors "GophKeeperClient/internal/errors/pair"
	"context"
)

func (ps *PairStorage) GetAll(ctx context.Context) ([]pair.Pair, error) {
	query := "SELECT pair_name, created_time FROM pairs"

	rows, err := ps.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	pairs := make([]pair.Pair, 0, 3)

	for rows.Next() {
		var p pair.Pair

		err = rows.Scan(&p.Name, &p.CreatedTime)
		if err != nil {
			return nil, err
		}

		pairs = append(pairs, p)

	}

	if rows.Err() != nil {
		return nil, err
	}

	if len(pairs) == 0 {
		return nil, pairerrors.ErrPairDoseNotExist
	}

	return pairs, nil
}
