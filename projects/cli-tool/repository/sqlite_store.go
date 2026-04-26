package repository

import (
	"database/sql"
	"fmt"

	"cli-tool/models"

	_ "modernc.org/sqlite"
)

type sqliteStore struct {
	db *sql.DB
}

// NewSQLiteTodoStore opens (or creates) a SQLite database at dbPath and returns a TodoStore.
func NewSQLiteTodoStore(dbPath string) (TodoStore, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("sqlite open: %w", err)
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("sqlite ping: %w", err)
	}
	if _, err := db.Exec(`
CREATE TABLE IF NOT EXISTS todos (
	id INTEGER PRIMARY KEY,
	title TEXT NOT NULL,
	status TEXT NOT NULL,
	created_at TEXT NOT NULL
)`); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("sqlite schema: %w", err)
	}
	return &sqliteStore{db: db}, nil
}

func (s *sqliteStore) List() []models.Todos {
	rows, err := s.db.Query(`SELECT id, title, status, created_at FROM todos ORDER BY id ASC`)
	if err != nil {
		fmt.Println("error querying sqlite,", err)
		return []models.Todos{}
	}
	defer rows.Close()

	var out []models.Todos
	for rows.Next() {
		var t models.Todos
		if err := rows.Scan(&t.Id, &t.Title, &t.Status, &t.CreatedAt); err != nil {
			fmt.Println("error scanning sqlite row,", err)
			return []models.Todos{}
		}
		out = append(out, t)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("sqlite rows error,", err)
		return []models.Todos{}
	}
	return out
}

func (s *sqliteStore) Add(todo models.Todos) error {
	existing := s.List()
	todo.Id = GetNextID(existing)

	if err := UpdateMaxTitleLength(int64(len(todo.Title))); err != nil {
		return err
	}

	_, err := s.db.Exec(
		`INSERT INTO todos (id, title, status, created_at) VALUES (?, ?, ?, ?)`,
		todo.Id, todo.Title, todo.Status, todo.CreatedAt,
	)
	if err != nil {
		fmt.Println("sqlite insert error,", err)
		return err
	}
	fmt.Println("TODO added successfully! ID:", todo.Id)
	return nil
}

func (s *sqliteStore) ReplaceAll(todos []models.Todos) error {
	tx, err := s.db.Begin()
	if err != nil {
		fmt.Println("sqlite begin tx error,", err)
		return err
	}
	if _, err := tx.Exec(`DELETE FROM todos`); err != nil {
		_ = tx.Rollback()
		fmt.Println("sqlite delete error,", err)
		return err
	}
	stmt, err := tx.Prepare(`INSERT INTO todos (id, title, status, created_at) VALUES (?, ?, ?, ?)`)
	if err != nil {
		_ = tx.Rollback()
		fmt.Println("sqlite prepare error,", err)
		return err
	}
	for _, t := range todos {
		if _, err := stmt.Exec(t.Id, t.Title, t.Status, t.CreatedAt); err != nil {
			_ = stmt.Close()
			_ = tx.Rollback()
			fmt.Println("sqlite insert error,", err)
			return err
		}
	}
	if err := stmt.Close(); err != nil {
		_ = tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		fmt.Println("sqlite commit error,", err)
		return err
	}
	fmt.Println("TODO updated successfully!")
	return nil
}
