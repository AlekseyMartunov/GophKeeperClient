package filestorage

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
)

func (fs *FileStorage) Save(ctx context.Context, bucketName, fileName string, b []byte, size int) error {
	_, err := fs.client.PutObject(
		ctx,
		bucketName,
		fileName,
		bytes.NewReader(b),
		int64(size),
		minio.PutObjectOptions{},
	)

	if err != nil {
		return err
	}

	return nil
}
