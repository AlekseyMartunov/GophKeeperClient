package filestorage

import (
	"GophKeeperClient/internal/entity/file"
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"time"
)

func (fs *FileStorage) Save(ctx context.Context, f *file.File) error {
	_, err := fs.client.PutObject(
		ctx,
		fs.config.MinioBucketName(),
		f.Name,
		bytes.NewReader(f.Data),
		int64(len(f.Data)),
		minio.PutObjectOptions{
			UserMetadata: map[string]string{
				"Created_time": f.CreatedTime.Format(time.RFC3339),
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}
