package repository

import "path/filepath"

// DataDir holds todos.json and setting.json. It defaults to "db". Tests may
// point it at a temporary directory so the real JSON files are not read or written.
var DataDir = "db"

func todosPath() string {
	return filepath.Join(DataDir, "todos.json")
}

func settingPath() string {
	return filepath.Join(DataDir, "setting.json")
}

func sqlitePath() string {
	return filepath.Join(DataDir, "todos.db")
}
