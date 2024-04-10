package fileservice

import (
	"GophKeeperClient/internal/entity/file"
	"context"
)

const defaultBucketName = "bucket1"

type storage interface {
	Save(ctx context.Context, f *file.File) error
	Delete(ctx context.Context, fileName string) error
	GetAllNames(ctx context.Context) ([]*file.File, error)
	GetFile(ctx context.Context, fileName string) (*file.File, error)
}

type client interface {
	Send(f file.File) error
	Get(fileName string) (*file.File, error)
	GetAll() ([]*file.File, error)
	Delete(fileName string) error
}

type encryptor interface {
	EncryptByte(data, key []byte) ([]byte, error)
	DecryptByte(data, key []byte) ([]byte, error)
}

type FileService struct {
	repo      storage
	client    client
	encryptor encryptor
}

func NewFileService(r storage, c client, e encryptor) *FileService {
	return &FileService{
		repo:      r,
		client:    c,
		encryptor: e,
	}
}

func (fs *FileService) SaveLocal(f *file.File, key string) error {
	f, err := fs.encrypt(f, key)
	if err != nil {
		return err
	}
	err = fs.repo.Save(context.Background(), f)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FileService) SaveRemote(f *file.File, key string) error {
	f, err := fs.encrypt(f, key)
	if err != nil {
		return err
	}
	return fs.client.Send(*f)
}

func (fs *FileService) GetFromLocal(fileName, key string) (*file.File, error) {
	f, err := fs.repo.GetFile(context.Background(), fileName)
	if err != nil {
		return nil, err
	}
	f, err = fs.decrypt(f, key)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs *FileService) GetFromRemote(fileName, key string) (*file.File, error) {
	f, err := fs.client.Get(fileName)
	if err != nil {
		return nil, err
	}
	f, err = fs.decrypt(f, key)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs *FileService) GetAllFromLocal() ([]*file.File, error) {
	return fs.repo.GetAllNames(context.Background())
}

func (fs *FileService) GetAllFromRemote() ([]*file.File, error) {
	return fs.client.GetAll()
}

func (fs *FileService) DeleteFromLocal(fileName string) error {
	return fs.repo.Delete(context.Background(), fileName)
}

func (fs *FileService) DeleteFromRemote(fileName string) error {
	return fs.client.Delete(fileName)
}

func (fs *FileService) encrypt(f *file.File, key string) (*file.File, error) {
	b, err := fs.encryptor.EncryptByte(f.Data, []byte(key))
	if err != nil {
		return nil, err
	}

	return &file.File{
		Name:        f.Name,
		Data:        b,
		CreatedTime: f.CreatedTime,
	}, nil
}

func (fs *FileService) decrypt(f *file.File, key string) (*file.File, error) {
	b, err := fs.encryptor.DecryptByte(f.Data, []byte(key))

	if err != nil {
		return nil, err
	}

	return &file.File{
		Name:        f.Name,
		Data:        b,
		CreatedTime: f.CreatedTime,
	}, nil
}
