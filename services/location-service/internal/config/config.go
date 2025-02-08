package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server Server
	DB     Database
	Kafka  Kafka
}

type Server struct {
	Port                string
	WeatherServicePort  string
	AnalyticServicePort string
	GatewayPort         string
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
	Brokers       []string
	LocationTopic string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("internal/config")

	viper.ReadInConfig()

	viper.SetDefault("POSTGRES_USER", "weather_user")
	viper.SetDefault("POSTGRES_PASSWORD", "weather_pass")
	viper.SetDefault("POSTGRES_DB", "weather_db")

	viper.AutomaticEnv()

	cfg := &Config{
		Server: Server{
			Port:                viper.GetString("server.port"),
			WeatherServicePort:  viper.GetString("server.weather_service_port"),
			AnalyticServicePort: viper.GetString("server.analytic_service_port"),
			GatewayPort:         viper.GetString("server.gateway_port"),
		},
		DB: Database{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			User:     viper.GetString("POSTGRES_USER"),
			Password: viper.GetString("POSTGRES_PASSWORD"),
			DBName:   viper.GetString("POSTGRES_DB"),
			SSLMode:  viper.GetString("database.sslmode"),
		},
		Kafka: Kafka{
			Brokers:       viper.GetStringSlice("kafka.brokers"),
			LocationTopic: viper.GetString("kafka.location_topic"),
		},
	}

	return cfg
}
