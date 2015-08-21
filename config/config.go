package config

type Config struct {
	DbDriver string
	DbName string
	DbUser string
	DbPass string

	EmailUser string
	EmailHost string
	EmailPort string
	EmailPassword string
}

func New() *Config {
	return &Config{
		"mysql",
		"first_go",
		"root",
		"",
		"test@gmail.com",
		"smtp.gmail.com",
		"587",
		"email_password",
	}
}