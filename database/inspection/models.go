package inspection

import (
	"time"
)

type Inspection struct {
	Id                  string    `json:"_id,omitempty" bson:"_id,omitempty"`
	InspectorId         int       `json:"inspector_id" bson:"inspector_id"`
	AccountNumber       string    `json:"account_number" bson:"account_number"`
	Consumer            Consumer  `json:"consumer" bson:"consumer"`
	Object              string    `json:"object" bson:"object"`
	Resolution          int       `json:"resolution" bson:"resolution"`
	Reason              string    `json:"reason" bson:"reason"`
	Method              string    `json:"method" bson:"method"`
	IsReviewRefused     bool      `json:"is_review_refused" bson:"is_review_refused"`
	ActionDate          time.Time `json:"action_date" bson:"action_date"`
	HaveAutomaton       bool      `json:"have_automaton" bson:"have_automaton"`
	AutomatonSealNumber string    `json:"automaton_seal_number" bson:"automaton_seal_number"`
	Images              []Image   `json:"images" bson:"images"`
	Device              Device    `json:"device" bson:"device"`
	InspectionDate      time.Time `json:"inspection_date" bson:"inspection_date"`
	ResolutionFileName  string    `json:"resolution_file_name" bson:"resolution_file_name"`
	ActNumber           string    `json:"act_number" bson:"act_number"`
	Contract            Contract  `json:"contract" bson:"contract"`
	SealNumber          string    `json:"seal_number" bson:"seal_number"`
	CreatedAt           time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" bson:"updated_at"`
}

type Consumer struct {
	Surname         string `json:"surname" bson:"surname"`
	Name            string `json:"name" bson:"name"`
	Patronymic      string `json:"patronymic" bson:"patronymic"`
	LegalEntityName string `json:"legal_entity_name" bson:"legal_entity_name"`
}

type Device struct {
	Type                string `json:"type" bson:"type"`
	Number              string `json:"number" bson:"number"`
	Voltage             string `json:"voltage" bson:"voltage"`
	Amperage            string `json:"amperage" bson:"amperage"`
	AccuracyClass       string `json:"accuracy_class" bson:"accuracy_class"`
	TariffsCount        int    `json:"tariffs_count" bson:"tariffs_count"`
	DeploymentPlace     string `json:"deployment_place" bson:"deployment_place"`
	ValencyBeforeDot    string `json:"valency_before_dot" bson:"valency_before_dot"`
	ValencyAfterDot     string `json:"valency_after_dot" bson:"valency_after_dot"`
	ManufactureYear     int    `json:"manufacture_year" bson:"manufacture_year"`
	VerificationQuarter int    `json:"verification_quarter" bson:"verification_quarter"`
	VerificationYear    int    `json:"verification_year" bson:"verification_year"`
	Value               string `json:"value" bson:"value"`
}

type Contract struct {
	Number string    `json:"number" bson:"number"`
	Date   time.Time `json:"date" bson:"date"`
}

type Image struct {
	Name string `json:"name" bson:"name"`
	URL  string `json:"url" bson:"url"`
}
