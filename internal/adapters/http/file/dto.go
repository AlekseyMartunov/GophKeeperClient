package fileclienthttp

import (
	"GophKeeperClient/internal/entity/file"
	"time"
)

type fileDTO struct {
	Name        string    `json:"name"`
	Data        []byte    `json:"data"`
	CreatedTime time.Time `json:"created_time"`
}

func (d *fileDTO) FromEntity(f file.File) {
	d.Name = f.Name
	d.Data = f.Data
	d.CreatedTime = f.CreatedTime
}

func (d *fileDTO) ToEntity() *file.File {
	return &file.File{
		Name:        d.Name,
		Data:        d.Data,
		CreatedTime: d.CreatedTime,
	}
}
