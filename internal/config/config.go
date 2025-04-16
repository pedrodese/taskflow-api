package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (*Config, error) {
	var config Config

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	config.DBHost = viper.GetString("DB_HOST")
	config.DBUser = viper.GetString("DB_USER")
	config.DBPassword = viper.GetString("DB_PASSWORD")
	config.DBName = viper.GetString("DB_NAME")

	return &config, nil
}

func (c *Config) ConnectionURL() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=5432 sslmode=disable", c.DBUser, c.DBPassword, c.DBName, c.DBHost)
}
