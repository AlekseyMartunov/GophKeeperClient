package main

import (
	"GophKeeperClient/internal/app"
	"context"
)

func main() {
	ctx := context.Background()

	err := app.Run(ctx)
	if err != nil {
		panic(err)
	}
}
