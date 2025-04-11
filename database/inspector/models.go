package inspector

type Inspector struct {
	Surname    string `json:"surname" bson:"surname"`
	Name       string `json:"name" bson:"name"`
	Patronymic string `json:"patronymic" bson:"patronymic"`
}
