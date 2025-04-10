package device

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
