package report

import (
	"time"
	"tns-energo/database/file"
)

type Report struct {
	Id        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Type      int       `json:"type" bson:"type"`
	File      file.File `json:"file" bson:"file"`
	ForDate   time.Time `json:"for_date" bson:"for_date"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
