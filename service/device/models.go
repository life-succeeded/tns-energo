package device

import "github.com/shopspring/decimal"

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
