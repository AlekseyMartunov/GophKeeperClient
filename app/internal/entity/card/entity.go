package card

import "time"

type Card struct {
	Name        string
	Number      string
	CVV         string
	Owner       string
	Date        string
	CreatedTime time.Time
}
