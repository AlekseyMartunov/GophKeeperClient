package filestorage

import (
	"GophKeeperClient/internal/entity/file"
	"context"
	"github.com/minio/minio-go/v7"
	"time"
)

func (fs *FileStorage) GetAllNames(ctx context.Context) ([]*file.File, error) {
	var files []*file.File

	for info := range fs.client.ListObjects(ctx, fs.config.MinioBucketName(), minio.ListObjectsOptions{WithMetadata: true}) {
		if info.Err != nil {
			return nil, info.Err
		}

		obj, err := fs.client.GetObject(ctx, fs.config.MinioBucketName(), info.Key, minio.GetObjectOptions{})
		if err != nil {
			return nil, err
		}

		objInfo, err := obj.Stat()
		if err != nil {
			return nil, err
		}

		time, err := time.Parse(time.RFC3339, objInfo.UserMetadata["Created_time"])
		if err != nil {
			return nil, err
		}

		files = append(files, &file.File{
			Name:        info.Key,
			CreatedTime: time,
		})
	}

	return files, nil
}
