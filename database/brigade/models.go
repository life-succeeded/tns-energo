package brigade

import (
	"time"
	"tns-energo/database/inspector"
)

type Brigade struct {
	Id              string              `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstInspector  inspector.Inspector `json:"first_inspector" bson:"first_inspector"`
	SecondInspector inspector.Inspector `json:"second_inspector" bson:"second_inspector"`
	CreatedAt       time.Time           `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at" bson:"updated_at"`
}
