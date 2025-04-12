package device

import (
	"tns-energo/service/seal"

	"github.com/shopspring/decimal"
)

type Device struct {
	Type            string          `json:"type,omitempty"`
	Number          string          `json:"number,omitempty"`
	DeploymentPlace DeploymentPlace `json:"deployment_place,omitempty"`
	OtherPlace      string          `json:"other_place,omitempty"`
	Seals           []seal.Seal     `json:"seals,omitempty"`
	Value           decimal.Decimal `json:"value,omitempty"`
	Consumption     decimal.Decimal `json:"consumption,omitempty"`
}
