package config

type Settings struct {
	Port        int         `json:"port"`
	Databases   Databases   `json:"databases"`
	Auth        Auth        `json:"auth"`
	Templates   Templates   `json:"templates"`
	Registry    Registry    `json:"registry"`
	Inspections Inspections `json:"inspections"`
}

type Databases struct {
	Postgres string `json:"postgres"`
	Mongo    string `json:"mongo"`
	Minio    Minio  `json:"minio"`
}

type Minio struct {
	Endpoint        string `json:"endpoint"`
	ImagesBucket    string `json:"images_bucket"`
	DocumentsBucket string `json:"documents_bucket"`
	User            string `json:"user"`
	Password        string `json:"password"`
	UseSSL          bool   `json:"use_ssl"`
	Host            string `json:"host"`
}

type Auth struct {
	Secret                 string `json:"secret"`
	TokenExpiresAfterHours int    `json:"token_expires_after_hours"`
}

type Templates struct {
	Limitation string `json:"limitation"`
	Resumption string `json:"resumption"`
}

type Registry struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
}

type Inspections struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
}
