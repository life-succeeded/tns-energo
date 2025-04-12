package config

import (
	"os"
	"tns-energo/lib/config"
)

func Parse() (Settings, error) {
	var settings Settings
	if err := config.Parse(&settings); err != nil {
		return Settings{}, err
	}

	settings.Databases.Mongo = os.Getenv("MONGO_CONNECTION_STRING")
	settings.Databases.Minio.User = os.Getenv("MINIO_USER")
	settings.Databases.Minio.Password = os.Getenv("MINIO_PASSWORD")

	return settings, nil
}
