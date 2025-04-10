package inspection

import (
	"time"
	"tns-energo/database/consumer"
	"tns-energo/database/contract"
	"tns-energo/database/device"
	"tns-energo/database/file"
)

type Inspection struct {
	Id                  string            `json:"_id,omitempty" bson:"_id,omitempty"`
	InspectorId         int               `json:"inspector_id" bson:"inspector_id"`
	TaskId              string            `json:"task_id" bson:"task_id"`
	AccountNumber       string            `json:"account_number" bson:"account_number"`
	Consumer            consumer.Consumer `json:"consumer" bson:"consumer"`
	Address             string            `json:"address" bson:"address"`
	Resolution          int               `json:"resolution" bson:"resolution"`
	Reason              string            `json:"reason" bson:"reason"`
	Method              string            `json:"method" bson:"method"`
	IsReviewRefused     bool              `json:"is_review_refused" bson:"is_review_refused"`
	ActionDate          time.Time         `json:"action_date" bson:"action_date"`
	HaveAutomaton       bool              `json:"have_automaton" bson:"have_automaton"`
	AutomatonSealNumber string            `json:"automaton_seal_number" bson:"automaton_seal_number"`
	Images              []file.File       `json:"images" bson:"images"`
	Device              device.Device     `json:"device" bson:"device"`
	InspectionDate      time.Time         `json:"inspection_date" bson:"inspection_date"`
	ResolutionFile      file.File         `json:"resolution_file" bson:"resolution_file"`
	ActNumber           string            `json:"act_number" bson:"act_number"`
	Contract            contract.Contract `json:"contract" bson:"contract"`
	SealNumber          string            `json:"seal_number" bson:"seal_number"`
	CreatedAt           time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at" bson:"updated_at"`
}
