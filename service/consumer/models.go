package consumer

type Consumer struct {
	Surname         string `json:"surname"`
	Name            string `json:"name"`
	Patronymic      string `json:"patronymic"`
	LegalEntityName string `json:"legal_entity_name"`
	PhoneNumber     string `json:"phone_number"`
}
