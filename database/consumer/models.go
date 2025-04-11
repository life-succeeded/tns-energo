package consumer

type Consumer struct {
	Surname     string `json:"surname" bson:"surname"`
	Name        string `json:"name" bson:"name"`
	Patronymic  string `json:"patronymic" bson:"patronymic"`
	PhoneNumber string `json:"phone_number" bson:"phone_number"`
}
