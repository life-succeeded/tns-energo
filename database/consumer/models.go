package consumer

type Consumer struct {
	Surname         string `json:"surname" bson:"surname"`
	Name            string `json:"name" bson:"name"`
	Patronymic      string `json:"patronymic" bson:"patronymic"`
	LegalEntityName string `json:"legal_entity_name" bson:"legal_entity_name"`
	PhoneNumber     string `json:"phone_number" bson:"phone_number"`
}
