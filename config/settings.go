package config

type Settings struct {
	Port      int       `json:"port"`
	Databases Databases `json:"databases"`
}

type Databases struct {
	Postgres string `json:"postgres"`
}
