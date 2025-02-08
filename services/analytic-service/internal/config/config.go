package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Database Database
	Kafka    Kafka
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

type Kafka struct {
	Brokers []string
	GroupId string
	Topic   string
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
		Kafka: Kafka{
			Brokers: viper.GetStringSlice("kafka.brokers"),
			GroupId: viper.GetString("kafka.group_id"),
			Topic:   viper.GetString("kafka.topic"),
		},
	}
	return cfg
}
