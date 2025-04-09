package inspection

import (
	"time"
	"tns-energo/service/image"

	"github.com/shopspring/decimal"
)

type Inspection struct {
	Id                  string         `json:"id"`
	InspectorId         int            `json:"inspector_id"`
	AccountNumber       string         `json:"account_number"`
	Consumer            Consumer       `json:"consumer"`
	Object              string         `json:"object"`
	Resolution          Resolution     `json:"resolution"`
	Reason              string         `json:"reason"`
	Method              string         `json:"method"`
	IsReviewRefused     bool           `json:"is_review_refused"`
	ActionDate          time.Time      `json:"action_date"`
	HaveAutomaton       bool           `json:"have_automaton"`
	AutomatonSealNumber string         `json:"automaton_seal_number"`
	Images              []image.Image  `json:"images"`
	Device              Device         `json:"device"`
	InspectionDate      time.Time      `json:"inspection_date"`
	ResolutionFile      ResolutionFile `json:"resolution_file"`
	ActNumber           string         `json:"act_number"`
	Contract            Contract       `json:"contract"`
	SealNumber          string         `json:"seal_number"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
}

type Consumer struct {
	Surname         string `json:"surname"`
	Name            string `json:"name"`
	Patronymic      string `json:"patronymic"`
	LegalEntityName string `json:"legal_entity_name"`
}

type Device struct {
	Type                string          `json:"type"`
	Number              string          `json:"number"`
	Voltage             string          `json:"voltage"`
	Amperage            string          `json:"amperage"`
	AccuracyClass       string          `json:"accuracy_class"`
	TariffsCount        int             `json:"tariffs_count"`
	DeploymentPlace     string          `json:"deployment_place"`
	ValencyBeforeDot    string          `json:"valency_before_dot"`
	ValencyAfterDot     string          `json:"valency_after_dot"`
	ManufactureYear     int             `json:"manufacture_year"`
	VerificationQuarter int             `json:"verification_quarter"`
	VerificationYear    int             `json:"verification_year"`
	Value               decimal.Decimal `json:"value"`
}

type Contract struct {
	Number string    `json:"number"`
	Date   time.Time `json:"date"`
}

type ResolutionFile struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
