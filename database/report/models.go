package report

import "time"

type Report struct {
	Id        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Type      int       `json:"type" bson:"type"`
	Name      string    `json:"name" bson:"name"`
	URL       string    `json:"url" bson:"url"`
	ForDate   time.Time `json:"for_date" bson:"for_date"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
