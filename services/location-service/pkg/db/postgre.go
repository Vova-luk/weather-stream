package db

import (
	"fmt"

	"github.com/Vova-luk/weather-stream/services/location-service/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectPostgre(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.DBName,
		cfg.DB.SSLMode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
