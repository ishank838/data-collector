package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WebPort string `default:"8080"`
}

var envFile = "local.env"

func InitConfig() (*Config, error) {
	err := godotenv.Load(envFile)

	if err != nil {
		return nil, err
	}

	webPort := os.Getenv("WEB_PORT")

	return &Config{
		WebPort: webPort,
	}, nil
}
