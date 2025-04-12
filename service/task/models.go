package task

import (
	"time"
	"tns-energo/service/consumer"
)

type Task struct {
	Id            string            `json:"id"`
	BrigadeId     string            `json:"brigade_id"`
	Address       string            `json:"address"`
	VisitDate     time.Time         `json:"visit_date"`
	Status        Status            `json:"status"`
	Consumer      consumer.Consumer `json:"consumer"`
	AccountNumber string            `json:"account_number"`
	Comment       string            `json:"comment"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
}

type AddOneRequest struct {
	BrigadeId     string            `json:"brigade_id"`
	Address       string            `json:"address"`
	VisitDate     time.Time         `json:"visit_date"`
	Consumer      consumer.Consumer `json:"consumer"`
	AccountNumber string            `json:"account_number"`
	Comment       string            `json:"comment"`
}

type UpdateStatusRequest struct {
	NewStatus Status `json:"new_status"`
}
