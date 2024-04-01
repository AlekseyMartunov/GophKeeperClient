package filestorage

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
)

func (fs *FileStorage) GetFile(ctx context.Context, bucketName, fileName string) ([]byte, error) {
	object, err := fs.client.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	info, err := object.Stat()
	if err != nil {
		return nil, err
	}
	buff := make([]byte, info.Size)
	_, err = object.Read(buff)

	if err != nil && err != io.EOF {
		return nil, err
	}

	return buff, nil
}
