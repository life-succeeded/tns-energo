package inspection

import (
	"time"
	"tns-energo/database/brigade"
	"tns-energo/database/consumer"
	"tns-energo/database/device"
	"tns-energo/database/file"

	"github.com/shopspring/decimal"
)

type Inspection struct {
	Id                      string            `json:"_id,omitempty" bson:"_id,omitempty"`
	TaskId                  string            `json:"task_id" bson:"task_id"`
	Brigade                 brigade.Brigade   `json:"brigade" bson:"brigade"`
	Type                    int               `json:"type" bson:"type"`
	ActNumber               int               `json:"act_number" bson:"act_number"`
	Resolution              int               `json:"resolution" bson:"resolution"`
	Address                 string            `json:"address" bson:"address"`
	Consumer                consumer.Consumer `json:"consumer" bson:"consumer"`
	HaveAutomaton           bool              `json:"have_automaton" bson:"have_automaton"`
	AccountNumber           string            `json:"account_number" bson:"account_number"`
	IsIncompletePayment     bool              `json:"is_incomplete_payment" bson:"is_incomplete_payment"`
	OtherReason             string            `json:"other_reason" bson:"other_reason"`
	MethodBy                int               `json:"method_by" bson:"method_by"`
	Method                  string            `json:"method" bson:"method"`
	Device                  device.Device     `json:"device" bson:"device"`
	ReasonType              int               `json:"reason_type" bson:"reason_type"`
	Reason                  string            `json:"reason" bson:"reason"`
	Images                  []file.File       `json:"images" bson:"images"`
	InspectionDate          time.Time         `json:"inspection_date" bson:"inspection_date"`
	ResolutionFile          file.File         `json:"resolution_file" bson:"resolution_file"`
	IsChecked               bool              `json:"is_checked" bson:"is_checked"`
	IsViolationDetected     bool              `json:"is_violation_detected" bson:"is_violation_detected"`
	IsExpenseAvailable      bool              `json:"is_expense_available" bson:"is_expense_available"`
	OtherViolation          string            `json:"other_violation" bson:"other_violation"`
	IsUnauthorizedConsumers bool              `json:"is_unauthorized_consumers" bson:"is_unauthorized_consumers"`
	UnauthorizedDescription string            `json:"unauthorized_description" bson:"unauthorized_description"`
	OldDeviceValue          decimal.Decimal   `json:"old_device_value" bson:"old_device_value"`
	OldDeviceValueDate      time.Time         `json:"old_device_value_date" bson:"old_device_value_date"`
	UnauthorizedExplanation string            `json:"unauthorized_explanation" bson:"unauthorized_explanation"`
	EnergyActionDate        time.Time         `json:"energy_action_date" bson:"energy_action_date"`
}
