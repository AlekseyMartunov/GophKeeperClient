package app

import (
	fileStorage "GophKeeperClient/internal/adapters/db/minio/file"
	cardStorage "GophKeeperClient/internal/adapters/db/postgres/card"
	migration "GophKeeperClient/internal/adapters/db/postgres/migrations"
	pairStorage "GophKeeperClient/internal/adapters/db/postgres/pair"
	cardClientHTTP "GophKeeperClient/internal/adapters/http/card"
	fileClientHTTP "GophKeeperClient/internal/adapters/http/file"
	pairClientHTTP "GophKeeperClient/internal/adapters/http/pair"
	tokenClientHTTP "GophKeeperClient/internal/adapters/http/token"
	userClientHTTP "GophKeeperClient/internal/adapters/http/users"
	"GophKeeperClient/internal/config"
	"GophKeeperClient/internal/encrypter"
	"GophKeeperClient/internal/gui"
	cardService "GophKeeperClient/internal/usecase/card"
	fileService "GophKeeperClient/internal/usecase/file"
	pairService "GophKeeperClient/internal/usecase/pair"
	"GophKeeperClient/pkg/logger"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Run(ctx context.Context) error {
	cfg := config.NewConfig()
	err := cfg.ParseFlags()
	if err != nil {
		return fmt.Errorf("parsing palgs error: %w", err)
	}

	err = migration.MigrationsUp(cfg)
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

	userClientHTTP := userClientHTTP.NewUserClientHTTP(cfg, log)

	pairStorage := pairStorage.NewPairStorage(pool)
	pairClient := pairClientHTTP.NewPairClientHTTP(cfg, log)
	pairService := pairService.NewPairService(pairStorage, pairClient, encryptor)

	cardStorage := cardStorage.NewCardStorage(pool)
	cardClient := cardClientHTTP.NewCardClientHTTP(cfg, log)
	cardService := cardService.NewCardService(cardStorage, cardClient, encryptor)

	fileStorage, err := fileStorage.NewFileStorage(ctx, cfg)
	if err != nil {
		return err
	}
	fileClient := fileClientHTTP.NewFileClientHTTP(cfg, log)

	fileService := fileService.NewFileService(fileStorage, fileClient, encryptor)

	gui.Run(cfg, userClientHTTP, tokenClient, cardService, pairService, fileService)

	return nil
}
