package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerPort  string
	DatabaseURL string
	LogLevel    string
	Secret      string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	return &Config{
		ServerPort:  viper.GetString("SERVER_PORT"),
		DatabaseURL: viper.GetString("DATABASE_DB_URL"),
		LogLevel:    viper.GetString("LOG_LEVEL"),
		Secret:      viper.GetString("SECRET"),
	}
}
