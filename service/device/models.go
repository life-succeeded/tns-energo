package device

import (
	"time"

	"github.com/shopspring/decimal"
)

type Device struct {
	Type                string          `json:"type,omitempty"`
	Number              string          `json:"number,omitempty"`
	Voltage             string          `json:"voltage,omitempty"`
	Amperage            string          `json:"amperage,omitempty"`
	AccuracyClass       string          `json:"accuracy_class,omitempty"`
	TariffsCount        int             `json:"tariffs_count,omitempty"`
	DeploymentPlace     string          `json:"deployment_place,omitempty"`
	ValencyBeforeDot    string          `json:"valency_before_dot,omitempty"`
	ValencyAfterDot     string          `json:"valency_after_dot,omitempty"`
	ManufactureYear     int             `json:"manufacture_year,omitempty"`
	VerificationQuarter int             `json:"verification_quarter,omitempty"`
	VerificationYear    int             `json:"verification_year,omitempty"`
	DeploymentDate      time.Time       `json:"deployment_date,omitempty"`
	Value               decimal.Decimal `json:"value,omitempty"`
}
