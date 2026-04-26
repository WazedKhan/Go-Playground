package repository

import (
	"path/filepath"
	"testing"

	"cli-tool/models"
)

func TestSQLiteTodoStore_roundTrip(t *testing.T) {
	dataDir := t.TempDir()
	prevDir := DataDir
	DataDir = dataDir
	t.Cleanup(func() { DataDir = prevDir })

	if err := WriteSetting(models.Setting{MaxTitleLength: 50}); err != nil {
		t.Fatal(err)
	}

	dbPath := filepath.Join(t.TempDir(), "todos.db")
	s, err := NewSQLiteTodoStore(dbPath)
	if err != nil {
		t.Fatal(err)
	}
	ss := s.(*sqliteStore)
	SetTodoStore(s)
	t.Cleanup(func() {
		SetTodoStore(nil)
		_ = ss.db.Close()
	})

	if err := AddTODO(models.Todos{Title: "alpha", Status: models.PENDING, CreatedAt: "2026-01-01"}); err != nil {
		t.Fatal(err)
	}
	got := GetTODOs()
	if len(got) != 1 || got[0].Title != "alpha" || got[0].Id != 1 {
		t.Fatalf("GetTODOs() = %#v", got)
	}

	if err := UpdateTODO([]models.Todos{{Id: 1, Title: "alpha", Status: models.DONE, CreatedAt: "2026-01-01"}}); err != nil {
		t.Fatal(err)
	}
	got = GetTODOs()
	if len(got) != 1 || got[0].Status != models.DONE {
		t.Fatalf("after UpdateTODO: %#v", got)
	}
}

func TestInit_selectsSQLiteWhenEnvSet(t *testing.T) {
	dir := t.TempDir()
	prevDir := DataDir
	DataDir = dir
	t.Setenv("CLI_TOOL_STORAGE", "sqlite")
	t.Cleanup(func() {
		DataDir = prevDir
		SetTodoStore(nil)
	})

	if err := Init(); err != nil {
		t.Fatal(err)
	}
	if ss, ok := todoStore.(*sqliteStore); ok {
		t.Cleanup(func() { _ = ss.db.Close() })
	} else {
		t.Fatal("expected *sqliteStore after Init with CLI_TOOL_STORAGE=sqlite")
	}
	if err := WriteSetting(models.Setting{MaxTitleLength: 100}); err != nil {
		t.Fatal(err)
	}
	if err := AddTODO(models.Todos{Title: "from env", Status: models.PENDING, CreatedAt: "t"}); err != nil {
		t.Fatal(err)
	}
	got := GetTODOs()
	if len(got) != 1 || got[0].Title != "from env" {
		t.Fatalf("GetTODOs() = %#v", got)
	}
}
