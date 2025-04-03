package config

type Settings struct {
	Port      int       `json:"port"`
	Databases Databases `json:"databases"`
	Auth      Auth      `json:"auth"`
}

type Databases struct {
	Postgres string `json:"postgres"`
	Minio    Minio  `json:"minio"`
}

type Minio struct {
	Endpoint        string `json:"endpoint"`
	ImagesBucket    string `json:"images_bucket"`
	DocumentsBucket string `json:"documents_bucket"`
	User            string `json:"user"`
	Password        string `json:"password"`
	UseSSL          bool   `json:"use_ssl"`
}

type Auth struct {
	Secret                 string `json:"secret"`
	TokenExpiresAfterHours int    `json:"token_expires_after_hours"`
}
