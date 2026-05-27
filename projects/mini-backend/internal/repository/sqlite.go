package repository

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)


var	db *sql.DB


func InitSQLIte(path string) error {
	var err error
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		return fmt.Errorf("failed to open sqlite3 db, %w:", err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS user(KEY TEXT PRIMARY KEY, VALUE TEXT)`)
	return err
}

func (s *sqliteStorage)Get(key string) (*string, error) {
	var value string
	err := db.QueryRow(`SELECT value FROM user WHERE key=?`, key).Scan(&value)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no data found!")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query, %w", err)
	}
	return &value, nil
}

func (s *sqliteStorage)Save(key, value string) (bool, error) {
	_, err := db.Exec(`INSERT INTO user (key, value) VALUES(?, ?)
	ON CONFLICT(key) DO UPDATE SET value=excluded.value`, key, value)
	if err != nil {
		return false, fmt.Errorf("failed to write into database, %w", err)
	}
	return true, nil
}
