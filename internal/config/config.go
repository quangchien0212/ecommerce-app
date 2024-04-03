package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	config *viper.Viper
}

func NewConfig() *Config {
	c := new(Config)
	c.config = viper.New()
	c.config.SetConfigFile(".env")
	err := c.config.ReadInConfig()
	if err != nil {
    log.Fatalf("Error while reading config file %s", err)
  }

	return c
}

func (c *Config) Get() *viper.Viper {
	if c.config == nil {
		log.Fatal("config not initialized")
	}

	return c.config
}
