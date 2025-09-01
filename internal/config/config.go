package config

import "os"

type (
	Config struct {
		Api APIConfig
		DB  DBConfig
	}

	APIConfig struct {
		Port string
	}

	DBConfig struct {
		Dsn string
	}
)

func New() Config {
	return Config{
		Api: APIConfig{
			Port: os.Getenv("API_PORT"),
		},
		DB: DBConfig{
			Dsn: os.Getenv("DATABASE_URL"),
		},
	}
}
