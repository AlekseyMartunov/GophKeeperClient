package filestorage

import (
	"context"
	"github.com/minio/minio-go/v7"
)

func (fs *FileStorage) GetAllNames(ctx context.Context, bucketName string) ([]string, error) {
	var files []string

	for info := range fs.client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{WithMetadata: true}) {
		if info.Err != nil {
			return nil, info.Err
		}
		files = append(files, info.Key)
	}

	return files, nil
}
