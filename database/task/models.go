package task

import (
	"time"
	"tns-energo/database/consumer"
)

type Task struct {
	Id            string            `json:"_id,omitempty" bson:"_id,omitempty"`
	InspectorId   int               `json:"inspector_id" bson:"inspector_id"`
	Address       string            `json:"address" bson:"address"`
	VisitDate     time.Time         `json:"visit_date" bson:"visit_date"`
	Status        int               `json:"status" bson:"status"`
	Consumer      consumer.Consumer `json:"consumer" bson:"consumer"`
	AccountNumber string            `json:"account_number" bson:"account_number"`
	Comment       string            `json:"comment" bson:"comment"`
	CreatedAt     time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at" bson:"updated_at"`
}
