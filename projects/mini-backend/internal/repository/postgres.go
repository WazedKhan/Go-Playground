package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)


func InitPostgres(connStr string) error {
	var err error
	db, err = sql.Open("pgx", connStr)
	if err != nil {
		return fmt.Errorf("failed to open postgres db, %w:", err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS kv_store(KEY TEXT PRIMARY KEY, VALUE TEXT)`)
	return err
}

func (p *postgresStorage)Get(key string) (*string, error) {
	var value string
	err := db.QueryRow(`SELECT value FROM kv_store WHERE key=$1`, key).Scan(&value)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no data found!")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query, %w", err)
	}
	return &value, nil
}

func (p *postgresStorage)Save(key, value string) (bool, error) {
	_, err := db.Exec(`INSERT INTO kv_store (key, value) VALUES($1, $2)
	ON CONFLICT(key) DO UPDATE SET value=excluded.value`, key, value)
	if err != nil {
		return false, fmt.Errorf("failed to write into database, %w", err)
	}
	return true, nil
}
