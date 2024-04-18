package filestorage

import (
	"context"
	"github.com/minio/minio-go/v7"
)

func (fs *FileStorage) Delete(ctx context.Context, fileName string) error {
	err := fs.client.RemoveObject(ctx, fs.config.MinioBucketName(), fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}

	return nil
}
