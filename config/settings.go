package config

type Settings struct {
	Port        int         `json:"port"`
	Databases   Databases   `json:"databases"`
	Templates   Templates   `json:"templates"`
	Registry    MongoConfig `json:"registry"`
	Inspections MongoConfig `json:"inspections"`
	Reports     MongoConfig `json:"reports"`
	Tasks       MongoConfig `json:"tasks"`
	Brigades    MongoConfig `json:"brigades"`
	Cron        Cron        `json:"cron"`
}

type Databases struct {
	Mongo string `json:"mongo"`
	Minio Minio  `json:"minio"`
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

type Templates struct {
	Universal string `json:"universal"`
	Control   string `json:"control"`
}

type MongoConfig struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
}

type Cron struct {
	DailyReportTime string `json:"daily_report_time"`
}
