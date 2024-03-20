package app

import (
	"GophKeeperClient/internal/gui"
	"context"
)

func Run(ctx context.Context) error {
	err := gui.Run()
	if err != nil {
		return err
	}

	return nil
}
