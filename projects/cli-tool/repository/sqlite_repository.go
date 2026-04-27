package repository

import (
	"database/sql"
	"fmt"

	"cli-tool/models"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteStore struct {
	db *sql.DB
}

func NewSqliteStore(db *sql.DB) TodoStore {
	return &sqliteStore{db: db}
}

func (s *sqliteStore) Init() error {
	_, err := s.db.Exec(`CREATE TABLE IF NOT EXISTS todos (
        id          INTEGER PRIMARY KEY AUTOINCREMENT,
        title       TEXT NOT NULL,
        status      TEXT NOT NULL DEFAULT 'PENDING',
        created_at  TEXT NOT NULL
    )`)
	return err
}

func InitSqliteStore(db *sql.DB) (TodoStore, error) {
	store := &sqliteStore{db: db}
	if err := store.Init(); err != nil {
		return nil, fmt.Errorf("failed to init sqlite store: %w", err)
	}
	return store, nil
}

func (s *sqliteStore) GetTodos() ([]models.Todos, error) {
	rows, err := s.db.Query("SELECT id, title, status, created_at FROM todos")
	if err != nil {
		return nil, fmt.Errorf("error querying todo: %w", err)
	}
	defer rows.Close()

	var todos []models.Todos
	for rows.Next() {
		var todo models.Todos
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Status, &todo.CreatedAt); err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (s *sqliteStore) AddTodo(todo models.Todos) error {
	_, err := s.db.Exec(
		"INSERT INTO todos (title, status, created_at) VALUES (?, ?, ?)",
		todo.Title, todo.Status, todo.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("error inserting todo: %w", err)
	}
	fmt.Println("TODO added successfully! Title:", todo.Title)
	return nil
}

func (s *sqliteStore) ReplaceAll(todos []models.Todos) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	if _, err := tx.Exec("DELETE FROM todos"); err != nil {
		tx.Rollback()
		return fmt.Errorf("error deleting todos: %w", err)
	}

	for _, todo := range todos {
		_, err := tx.Exec(
			"INSERT INTO todos (title, status, created_at) VALUES (?, ?, ?)",
			todo.Title, todo.Status, todo.CreatedAt,
		)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("error inserting todo: %w", err)
		}
	}

	return tx.Commit()
}
