package filestorage

import (
	"GophKeeperClient/internal/entity/file"
	fileserrors "GophKeeperClient/internal/errors/files"
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"time"
)

const minioErrorMessage = "The specified key does not exist."

func (fs *FileStorage) GetFile(ctx context.Context, fileName string) (*file.File, error) {
	object, err := fs.client.GetObject(ctx, fs.config.MinioBucketName(), fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	info, err := object.Stat()
	if err != nil {
		if err.Error() == minioErrorMessage {
			return nil, fileserrors.ErrFailDoseNotExist
		}
		return nil, err
	}
	buff := make([]byte, info.Size)
	_, err = object.Read(buff)

	if err != nil && err != io.EOF {
		return nil, err
	}

	time, err := time.Parse(time.RFC3339, info.UserMetadata["Created_time"])
	if err != nil {
		return nil, err
	}

	f := file.File{
		Name:        info.Key,
		Data:        buff,
		CreatedTime: time,
	}

	return &f, nil
}
