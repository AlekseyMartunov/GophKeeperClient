package cardclienthttp

import (
	"GophKeeperClient/internal/entity/card"
	"time"
)

type dto struct {
	Name        string    `json:"card_name"`
	Number      string    `json:"number"`
	CVV         string    `json:"cvv"`
	Owner       string    `json:"owner"`
	Date        string    `json:"date"`
	CreatedTime time.Time `json:"created_time"`
}

func (d *dto) fromEntity(c card.Card) {
	d.Name = c.Name
	d.Number = c.Number
	d.Owner = c.Owner
	d.CVV = c.CVV
	d.Date = c.Date
	d.CreatedTime = c.CreatedTime
}

func (d *dto) toEntity() card.Card {
	c := card.Card{}

	c.Name = d.Name
	c.CVV = d.CVV
	c.Number = d.Number
	c.Owner = d.Owner
	c.Date = d.Date
	c.CreatedTime = d.CreatedTime

	return c
}
