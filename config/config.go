package config

import "os"

type Config struct {
	SiteDomain string

	DbDriver string
	DbName string
	DbUser string
	DbPass string
	DbParams string

	EmailUser string
	EmailHost string
	EmailPort string
	EmailPassword string
}

func New() *Config {
	return &Config{
		os.Getenv("APP_DOMAIN"),

		"mysql",
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_USER"),
		os.Getenv("APP_DB_PASS"),
		"?parseTime=true",

		os.Getenv("APP_EMAIL_USER"),
		os.Getenv("APP_EMAIL_HOST"),
		os.Getenv("APP_EMAIL_PORT"),
		os.Getenv("APP_EMAIL_PASS"),
	}
}