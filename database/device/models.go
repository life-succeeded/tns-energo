package device

import "time"

type Device struct {
	Type                string    `json:"type,omitempty" bson:"type,omitempty"`
	Number              string    `json:"number,omitempty" bson:"number,omitempty"`
	Voltage             string    `json:"voltage,omitempty" bson:"voltage,omitempty"`
	Amperage            string    `json:"amperage,omitempty" bson:"amperage,omitempty"`
	AccuracyClass       string    `json:"accuracy_class,omitempty" bson:"accuracy_class,omitempty"`
	TariffsCount        int       `json:"tariffs_count,omitempty" bson:"tariffs_count,omitempty"`
	DeploymentPlace     string    `json:"deployment_place,omitempty" bson:"deployment_place,omitempty"`
	ValencyBeforeDot    string    `json:"valency_before_dot,omitempty" bson:"valency_before_dot,omitempty"`
	ValencyAfterDot     string    `json:"valency_after_dot,omitempty" bson:"valency_after_dot,omitempty"`
	ManufactureYear     int       `json:"manufacture_year,omitempty" bson:"manufacture_year,omitempty"`
	VerificationQuarter int       `json:"verification_quarter,omitempty" bson:"verification_quarter,omitempty"`
	VerificationYear    int       `json:"verification_year,omitempty" bson:"verification_year,omitempty"`
	DeploymentDate      time.Time `json:"deployment_date,omitempty" bson:"deployment_date,omitempty"`
	Value               string    `json:"value,omitempty" bson:"value,omitempty"`
}
