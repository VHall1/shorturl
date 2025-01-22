package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/vrischmann/envconfig"
)

func Load(conf interface{}) error {
	env := os.Getenv("GO_ENV")

	// only load .env in dev mode
	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			return fmt.Errorf("falied to .env: %v", err)
		}
	}

	if err := envconfig.Init(conf); err != nil {
		return fmt.Errorf("falied to load config: %v", err)
	}

	return nil
}
