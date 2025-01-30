package config

import (
	"fmt"

	"github.com/vrischmann/envconfig"
)

func Load(conf interface{}) error {
	if err := envconfig.Init(conf); err != nil {
		return fmt.Errorf("falied to load config: %v", err)
	}

	return nil
}
