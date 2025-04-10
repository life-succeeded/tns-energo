package contract

import "time"

type Contract struct {
	Number string    `json:"number"`
	Date   time.Time `json:"date"`
}
