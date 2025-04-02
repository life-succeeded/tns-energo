package config

type Settings struct {
	Port      int       `json:"port"`
	Databases Databases `json:"databases"`
	Auth      Auth      `json:"auth"`
}

type Databases struct {
	Postgres string `json:"postgres"`
}

type Auth struct {
	Secret                 string `json:"secret"`
	TokenExpiresAfterHours int    `json:"token_expires_after_hours"`
}
