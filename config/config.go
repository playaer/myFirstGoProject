package config

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
		"http://localhost:3000",

		"mysql",
		"first_go",
		"root",
		"root",
		"?parseTime=true",

		"playaer.my.test@gmail.com",
		"smtp.gmail.com",
		"587",
		"mkmk1980",
	}
}