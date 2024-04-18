package pairclienthttp

import (
	"GophKeeperClient/internal/entity/pair"
	"time"
)

type dto struct {
	Name        string    `json:"name"`
	Login       string    `json:"login"`
	Password    string    `json:"password"`
	CreatedTime time.Time `json:"created_time"`
}

func (d *dto) fromEntity(p pair.Pair) {
	d.Name = p.Name
	d.Login = p.Login
	d.Password = p.Password
	d.CreatedTime = p.CreatedTime
}

func (d *dto) toEntity() pair.Pair {
	p := pair.Pair{}

	p.Name = d.Name
	p.Login = d.Login
	p.Password = d.Password
	p.CreatedTime = d.CreatedTime

	return p
}
