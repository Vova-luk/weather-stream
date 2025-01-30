package db

import (
	"fmt"

	"github.com/Vova-luk/weather-stream/location-service/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func ConnectPosrge(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.DBName,
		cfg.DB.SSLMode,
	)

	fmt.Println("Viper DB Config:", viper.AllSettings())
	fmt.Printf("Loaded DB Config: Host=%s Port=%s User=%s Password=%s DBName=%s SSLMode=%s\n",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName, cfg.DB.SSLMode)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
