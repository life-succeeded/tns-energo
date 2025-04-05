package inspection

type Inspection struct {
	Surname      string `bson:"surname" json:"surname"`
	Name         string `bson:"name" json:"name"`
	Patronymic   string `bson:"patronymic" json:"patronymic"`
	Position     string `bson:"position" json:"position"`
	TariffsCount int    `bson:"tariffs_count" json:"tariffs_count"`
}
