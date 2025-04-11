package device

import (
	"tns-energo/database/seal"

	"github.com/shopspring/decimal"
)

type Device struct {
	Type            string          `json:"type,omitempty" bson:"type,omitempty"`
	Number          string          `json:"number,omitempty" bson:"number,omitempty"`
	DeploymentPlace int             `json:"deployment_place,omitempty" bson:"deployment_place,omitempty"`
	OtherPlace      string          `json:"other_place,omitempty" bson:"other_place,omitempty"`
	Seals           []seal.Seal     `json:"seals,omitempty" bson:"seals,omitempty"`
	Value           decimal.Decimal `json:"value,omitempty" bson:"value,omitempty"`
}
