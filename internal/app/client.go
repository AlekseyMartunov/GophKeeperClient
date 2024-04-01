package app

import (
	fileStorage "GophKeeperClient/internal/adapters/db/minio/file"
	cardStorage "GophKeeperClient/internal/adapters/db/postgres/card"
	migration "GophKeeperClient/internal/adapters/db/postgres/migrations"
	pairStorage "GophKeeperClient/internal/adapters/db/postgres/pair"
	cardClientHTTP "GophKeeperClient/internal/adapters/http/card"
	pairClientHTTP "GophKeeperClient/internal/adapters/http/pair"
	tokenClientHTTP "GophKeeperClient/internal/adapters/http/token"
	"GophKeeperClient/internal/config"
	"GophKeeperClient/internal/encrypter"
	cardService "GophKeeperClient/internal/usecase/card"
	pairService "GophKeeperClient/internal/usecase/pair"
	"GophKeeperClient/pkg/logger"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(ctx context.Context) error {
	cfg := config.NewConfig()

	err := migration.MigrationsUp(cfg.PostgresDSN())
	if err != nil {
		return fmt.Errorf("migration err: %w", err)
	}

	pool, err := pgxpool.New(ctx, cfg.PostgresDSN())
	if err != nil {
		return fmt.Errorf("pool creation err: %w", err)
	}

	encryptor := encrypter.NewEncryptionManager()
	log := logger.NewLogger()
	tokenClient := tokenClientHTTP.NewTokenClientHTTP(cfg, log)

	err = tokenClient.UpdateToken("c1-4", "a", "a")
	if err != nil {
		fmt.Println(err)
	}

	pairStorage := pairStorage.NewPairStorage(pool)
	pairClient := pairClientHTTP.NewPairClientHTTP(cfg, log)
	pairService := pairService.NewPairService(pairStorage, pairClient, encryptor)

	cardStorage := cardStorage.NewCardStorage(pool)
	cardClient := cardClientHTTP.NewCardClientHTTP(cfg, log)
	cardService := cardService.NewPairService(cardStorage, cardClient, encryptor)

	fileStorage, err := fileStorage.NewFileStorage(ctx, cfg)
	if err != nil {
		return err
	}

	fmt.Println(pairService, cardService, fileStorage)

	return nil
}
