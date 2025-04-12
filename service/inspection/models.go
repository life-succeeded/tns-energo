package inspection

import (
	"time"
	"tns-energo/service/brigade"
	"tns-energo/service/consumer"
	"tns-energo/service/device"
	"tns-energo/service/file"

	"github.com/shopspring/decimal"
)

type Inspection struct {
	Id                      string            `json:"id"`
	TaskId                  string            `json:"task_id"`
	Brigade                 brigade.Brigade   `json:"brigade"`
	Type                    Type              `json:"type"`
	ActNumber               int               `json:"act_number"`
	Resolution              Resolution        `json:"resolution"`
	Address                 string            `json:"address"`
	Consumer                consumer.Consumer `json:"consumer"`
	HaveAutomaton           bool              `json:"have_automaton"`
	AccountNumber           string            `json:"account_number"`
	IsIncompletePayment     bool              `json:"is_incomplete_payment"`
	OtherReason             string            `json:"other_reason"`
	MethodBy                MethodBy          `json:"method_by"`
	Method                  string            `json:"method"`
	Device                  device.Device     `json:"device"`
	ReasonType              ReasonType        `json:"reason_type"`
	Reason                  string            `json:"reason"`
	ActCopies               int               `json:"act_copies"`
	Images                  []file.File       `json:"images"`
	InspectionDate          time.Time         `json:"inspection_date"`
	ResolutionFile          file.File         `json:"resolution_file"`
	IsChecked               bool              `json:"is_checked"`
	IsViolationDetected     bool              `json:"is_violation_detected"`
	IsExpenseAvailable      bool              `json:"is_expense_available"`
	OtherViolation          string            `json:"other_violation"`
	IsUnauthorizedConsumers bool              `json:"is_unauthorized_consumers"`
	UnauthorizedDescription string            `json:"unauthorized_description"`
	OldDeviceValue          decimal.Decimal   `json:"old_device_value"`
	OldDeviceValueDate      time.Time         `json:"old_device_value_date"`
	UnauthorizedExplanation string            `json:"unauthorized_explanation"`
	EnergyActionDate        time.Time         `json:"energy_action_date"`
}

type InspectRequest struct {
	TaskId                  string            `json:"task_id"`
	BrigadeId               string            `json:"brigade_id"`
	Type                    Type              `json:"type"`
	Resolution              Resolution        `json:"resolution"`
	Address                 string            `json:"address"`
	Consumer                consumer.Consumer `json:"consumer"`
	HaveAutomaton           bool              `json:"have_automaton"`
	AccountNumber           string            `json:"account_number"`
	IsIncompletePayment     bool              `json:"is_incomplete_payment"`
	OtherReason             string            `json:"other_reason"`
	MethodBy                MethodBy          `json:"method_by"`
	Method                  string            `json:"method"`
	Device                  device.Device     `json:"device"`
	ReasonType              ReasonType        `json:"reason_type"`
	Reason                  string            `json:"reason"`
	ActCopies               int               `json:"act_copies"`
	Images                  []file.File       `json:"images"`
	IsChecked               bool              `json:"is_checked"`
	IsViolationDetected     bool              `json:"is_violation_detected"`
	IsExpenseAvailable      bool              `json:"is_expense_available"`
	OtherViolation          string            `json:"other_violation"`
	IsUnauthorizedConsumers bool              `json:"is_unauthorized_consumers"`
	UnauthorizedDescription string            `json:"unauthorized_description"`
	OldDeviceValue          decimal.Decimal   `json:"old_device_value"`
	OldDeviceValueDate      time.Time         `json:"old_device_value_date"`
	UnauthorizedExplanation string            `json:"unauthorized_explanation"`
	EnergyActionDate        time.Time         `json:"energy_action_date"`
}
