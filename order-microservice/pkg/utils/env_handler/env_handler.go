package env_handler

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnvironment() {
	if _, err := os.Stat(".env.dev"); err == nil {
		err := godotenv.Load(".env.dev")
		if err != nil {
			panic("Error loading .env.dev file")
		}
	}
}
