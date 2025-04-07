package inspection

import (
	"time"

	"github.com/shopspring/decimal"
)

type Inspection struct {
	Id                 string          `json:"id"`
	AccountNumber      string          `json:"account_number"`
	ConsumerSurname    string          `json:"consumer_surname"`
	ConsumerName       string          `json:"consumer_name"`
	ConsumerPatronymic *string         `json:"consumer_patronymic"`
	Object             string          `json:"object"`
	Resolution         Resolution      `json:"resolution"`
	Reason             string          `json:"reason"`
	Method             string          `json:"method"`
	IsReviewRefused    bool            `json:"is_review_refused"`
	ActionDate         time.Time       `json:"action_date"`
	HaveAutomaton      bool            `json:"have_automaton"`
	Images             []string        `json:"images"`
	DeviceType         string          `json:"device_type"`
	DeviceNumber       string          `json:"device_number"`
	Voltage            string          `json:"voltage"`
	Amperage           string          `json:"amperage"`
	DeviceValue        decimal.Decimal `json:"device_value"`
	InspectionDate     time.Time       `json:"inspection_date"`
	AccuracyClass      string          `json:"accuracy_class"`
	TariffsCount       int             `json:"tariffs_count"`
	DeploymentPlace    string          `json:"deployment_place"`
}
