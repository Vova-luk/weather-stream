package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server
	DB     Database
}

type Server struct {
	Port        string
	GatewayPort string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("internal/config")

	viper.ReadInConfig()

	viper.SetDefault("database.postgres_user", "weather_user")
	viper.SetDefault("database.postgres_password", "weather_pass")
	viper.SetDefault("database.postgres_db", "weather_db")

	viper.AutomaticEnv()

	cfg := &Config{
		Server: Server{
			Port:        viper.GetString("server.port"),
			GatewayPort: viper.GetString("server.gateway_port"),
		},
		DB: Database{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			User:     viper.GetString("POSTGRES_USER"),
			Password: viper.GetString("POSTGRES_PASSWORD"),
			DBName:   viper.GetString("POSTGRES_DB"),
			SSLMode:  viper.GetString("database.sslmode"),
		},
	}

	fmt.Printf("Loaded Config: Host=%s Port=%s User=%s Password=***** DBName=%s SSLMode=%s\n",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.DBName, cfg.DB.SSLMode)

	return cfg, nil
}
