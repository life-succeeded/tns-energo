package inspection

import (
	"time"
	"tns-energo/database/consumer"
	"tns-energo/database/device"
	"tns-energo/database/file"
)

type Inspection struct {
	Id                  string            `json:"_id,omitempty" bson:"_id,omitempty"`
	TaskId              string            `json:"task_id" bson:"task_id"`
	BrigadeId           string            `json:"brigade_id" bson:"brigade_id"`
	ActNumber           string            `json:"act_number" bson:"act_number"`
	Resolution          int               `json:"resolution" bson:"resolution"`
	Address             string            `json:"address" bson:"address"`
	Consumer            consumer.Consumer `json:"consumer" bson:"consumer"`
	HaveAutomaton       bool              `json:"have_automaton" bson:"have_automaton"`
	AccountNumber       string            `json:"account_number" bson:"account_number"`
	IsIncompletePayment bool              `json:"is_incomplete_payment" bson:"is_incomplete_payment"`
	OtherReason         string            `json:"other_reason" bson:"other_reason"`
	MethodBy            int               `json:"method_by" bson:"method_by"`
	Method              string            `json:"method" bson:"method"`
	Device              device.Device     `json:"device" bson:"device"`
	ReasonType          int               `json:"reason_type" bson:"reason_type"`
	Reason              string            `json:"reason" bson:"reason"`
	ActCopies           int               `json:"act_copies" bson:"act_copies"`
	Images              []file.File       `json:"images" bson:"images"`
	InspectionDate      time.Time         `json:"inspection_date" bson:"inspection_date"`
	ResolutionFile      file.File         `json:"resolution_file" bson:"resolution_file"`
}
