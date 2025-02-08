package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Database Database
}

type Server struct {
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("internal/config")

	viper.SetDefault("POSTGRES_USER", "weather_user")
	viper.SetDefault("POSTGRES_PASSWORD", "weather_pass")
	viper.SetDefault("POSTGRES_DB", "weather_db")

	cfg := &Config{
		Server: Server{
			Port: viper.GetString("server.port"),
		},
		Database: Database{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			User:     viper.GetString("POSTGRES_USER"),
			Password: viper.GetString("PORTGRES_PASSWORD"),
			DBName:   viper.GetString("POSTGRES_DB"),
			SSLMode:  viper.GetString("database.sslmode"),
		},
	}
	return cfg
}
