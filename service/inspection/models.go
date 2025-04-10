package inspection

import (
	"time"
	"tns-energo/service/consumer"
	"tns-energo/service/contract"
	"tns-energo/service/device"
	"tns-energo/service/file"
)

type Inspection struct {
	Id                  string            `json:"id"`
	InspectorId         int               `json:"inspector_id"`
	TaskId              string            `json:"task_id"`
	AccountNumber       string            `json:"account_number"`
	Consumer            consumer.Consumer `json:"consumer"`
	Address             string            `json:"address"`
	Resolution          Resolution        `json:"resolution"`
	Reason              string            `json:"reason"`
	Method              string            `json:"method"`
	IsReviewRefused     bool              `json:"is_review_refused"`
	ActionDate          time.Time         `json:"action_date"`
	HaveAutomaton       bool              `json:"have_automaton"`
	AutomatonSealNumber string            `json:"automaton_seal_number"`
	Images              []file.File       `json:"images"`
	Device              device.Device     `json:"device"`
	InspectionDate      time.Time         `json:"inspection_date"`
	ResolutionFile      file.File         `json:"resolution_file"`
	ActNumber           string            `json:"act_number"`
	Contract            contract.Contract `json:"contract"`
	SealNumber          string            `json:"seal_number"`
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
}

type InspectRequest struct {
	TaskId              string            `json:"task_id"`
	AccountNumber       string            `json:"account_number"`
	Consumer            consumer.Consumer `json:"consumer"`
	Address             string            `json:"address"`
	Resolution          Resolution        `json:"resolution"`
	Reason              string            `json:"reason"`
	Method              string            `json:"method"`
	IsReviewRefused     bool              `json:"is_review_refused"`
	ActionDate          time.Time         `json:"action_date"`
	HaveAutomaton       bool              `json:"have_automaton"`
	AutomatonSealNumber string            `json:"automaton_seal_number"`
	Images              []file.File       `json:"images"`
	Device              device.Device     `json:"device"`
	InspectionDate      time.Time         `json:"inspection_date"`
	ActNumber           string            `json:"act_number"`
	Contract            contract.Contract `json:"contract"`
	SealNumber          string            `json:"seal_number"`
}
