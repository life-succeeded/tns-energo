package inspection

import (
	"time"
	"tns-energo/service/consumer"
	"tns-energo/service/device"
	"tns-energo/service/file"
)

type Inspection struct {
	Id                  string            `json:"id"`
	TaskId              string            `json:"task_id"`
	BrigadeId           string            `json:"brigade_id"`
	ActNumber           string            `json:"act_number"`
	Resolution          Resolution        `json:"resolution"`
	Address             string            `json:"address"`
	Consumer            consumer.Consumer `json:"consumer"`
	HaveAutomaton       bool              `json:"have_automaton"`
	AccountNumber       string            `json:"account_number"`
	IsIncompletePayment bool              `json:"is_incomplete_payment"`
	OtherReason         string            `json:"other_reason"`
	MethodBy            MethodBy          `json:"method_by"`
	Method              string            `json:"method"`
	Device              device.Device     `json:"device"`
	ReasonType          ReasonType        `json:"reason_type"`
	Reason              string            `json:"reason"`
	ActCopies           int               `json:"act_copies"`
	Images              []file.File       `json:"images"`
	InspectionDate      time.Time         `json:"inspection_date"`
	ResolutionFile      file.File         `json:"resolution_file"`
}

type InspectUniversalRequest struct {
	TaskId              string            `json:"task_id"`
	BrigadeId           string            `json:"brigade_id"`
	ActNumber           string            `json:"act_number"`
	Resolution          Resolution        `json:"resolution"`
	Address             string            `json:"address"`
	Consumer            consumer.Consumer `json:"consumer"`
	HaveAutomaton       bool              `json:"have_automaton"`
	AccountNumber       string            `json:"account_number"`
	IsIncompletePayment bool              `json:"is_incomplete_payment"`
	OtherReason         string            `json:"other_reason"`
	MethodBy            MethodBy          `json:"method_by"`
	Method              string            `json:"method"`
	Device              device.Device     `json:"device"`
	ReasonType          ReasonType        `json:"reason_type"`
	Reason              string            `json:"reason"`
	ActCopies           int               `json:"act_copies"`
	Images              []file.File       `json:"images"`
}
