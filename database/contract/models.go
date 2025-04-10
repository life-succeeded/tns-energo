package contract

import "time"

type Contract struct {
	Number string    `json:"number" bson:"number"`
	Date   time.Time `json:"date" bson:"date"`
}
