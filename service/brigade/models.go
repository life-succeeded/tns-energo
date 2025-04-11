package brigade

import (
	"time"
	"tns-energo/service/inspector"
)

type Brigade struct {
	Id              string              `json:"id"`
	FirstInspector  inspector.Inspector `json:"first_inspector"`
	SecondInspector inspector.Inspector `json:"second_inspector"`
	CreatedAt       time.Time           `json:"created_at"`
}

type CreateRequest struct {
	FirstInspector  inspector.Inspector `json:"first_inspector"`
	SecondInspector inspector.Inspector `json:"second_inspector"`
}
