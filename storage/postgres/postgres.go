package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"reservation-service/config"
)

func Conn() (*sql.DB, error) {
	cfg := config.Config{}
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_DATABASE)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}
