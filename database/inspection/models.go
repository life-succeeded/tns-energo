package inspection

import (
	"time"

	"github.com/shopspring/decimal"
)

type Inspection struct {
	Id                 string          `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountNumber      string          `json:"account_number" bson:"account_number"`
	ConsumerSurname    string          `json:"consumer_surname" bson:"consumer_surname"`
	ConsumerName       string          `json:"consumer_name" bson:"consumer_name"`
	ConsumerPatronymic *string         `json:"consumer_patronymic" bson:"consumer_patronymic"`
	Object             string          `json:"object" bson:"object"`
	Resolution         int             `json:"resolution" bson:"resolution"`
	Reason             string          `json:"reason" bson:"reason"`
	Method             string          `json:"method" bson:"method"`
	IsReviewRefused    bool            `json:"is_review_refused" bson:"is_review_refused"`
	ActionDate         time.Time       `json:"action_date" bson:"action_date"`
	HaveAutomaton      bool            `json:"have_automaton" bson:"have_automaton"`
	Images             []string        `json:"images" bson:"images"`
	DeviceType         string          `json:"device_type" bson:"device_type"`
	DeviceNumber       string          `json:"device_number" bson:"device_number"`
	Voltage            string          `json:"voltage" bson:"voltage"`
	Amperage           string          `json:"amperage" bson:"amperage"`
	DeviceValue        decimal.Decimal `json:"device_value" bson:"device_value"`
	InspectionDate     time.Time       `json:"inspection_date" bson:"inspection_date"`
	AccuracyClass      string          `json:"accuracy_class" bson:"accuracy_class"`
	TariffsCount       int             `json:"tariffs_count" bson:"tariffs_count"`
	DeploymentPlace    string          `json:"deployment_place" bson:"deployment_place"`
}
