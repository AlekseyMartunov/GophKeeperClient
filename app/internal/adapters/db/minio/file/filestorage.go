package filestorage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

const (
	location = "us-east-1"
)

type config interface {
	MinioAccessKeyID() string
	MinioSecretAccessKey() string
	MinioEndpoint() string
	MinioBucketName() string
}

type FileStorage struct {
	client *minio.Client
	config config
}

func NewFileStorage(ctx context.Context, cfg config) (*FileStorage, error) {
	minioClient, err := minio.New(cfg.MinioEndpoint(), &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKeyID(), cfg.MinioSecretAccessKey(), ""),
		Secure: false,
	})

	err = minioClient.MakeBucket(ctx, cfg.MinioBucketName(), minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, cfg.MinioBucketName())
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", cfg.MinioBucketName())
		} else {
			return nil, err
		}
	} else {
		log.Printf("Successfully created %s\n", cfg.MinioBucketName())
	}

	return &FileStorage{client: minioClient, config: cfg}, nil
}
