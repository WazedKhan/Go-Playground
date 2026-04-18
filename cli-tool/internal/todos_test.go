package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"cli-tool/models"
	"cli-tool/repository"
)

func swapDataDir(t *testing.T, dir string) {
	t.Helper()
	prev := repository.DataDir
	repository.DataDir = dir
	t.Cleanup(func() { repository.DataDir = prev })
}

func writeTodosJSON(t *testing.T, dir string, todos []models.Todos) {
	t.Helper()
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "todos.json")
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatal(err)
	}
}

func writeSettingJSON(t *testing.T, dir string, maxLen int) {
	t.Helper()
	s := models.Setting{MaxTitleLength: maxLen}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "setting.json")
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatal(err)
	}
}

func readTodosFromDisk(t *testing.T, dir string) []models.Todos {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(dir, "todos.json"))
	if err != nil {
		t.Fatal(err)
	}
	var out []models.Todos
	if err := json.Unmarshal(raw, &out); err != nil {
		t.Fatal(err)
	}
	return out
}

func TestGetTodos_usesIsolatedDataDir(t *testing.T) {
	dir := t.TempDir()
	swapDataDir(t, dir)
	seed := []models.Todos{
		{Id: 1, Title: "a", Status: models.PENDING, CreatedAt: "2026-01-01 00:00:00"},
	}
	writeTodosJSON(t, dir, seed)

	got, err := GetTodos()
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 1 || got[0].Id != 1 || got[0].Title != "a" {
		t.Fatalf("GetTodos() = %#v, want one seeded todo", got)
	}
}

func TestGetFilteredTodos(t *testing.T) {
	dir := t.TempDir()
	swapDataDir(t, dir)
	writeTodosJSON(t, dir, []models.Todos{
		{Id: 1, Title: "open", Status: models.PENDING, CreatedAt: "t1"},
		{Id: 2, Title: "closed", Status: models.DONE, CreatedAt: "t2"},
		{Id: 3, Title: "also open", Status: models.PENDING, CreatedAt: "t3"},
	})

	pending, err := GetFilteredTodos(models.PENDING)
	if err != nil {
		t.Fatal(err)
	}
	if len(pending) != 2 {
		t.Fatalf("len(pending) = %d, want 2", len(pending))
	}

	done, err := GetFilteredTodos(models.DONE)
	if err != nil {
		t.Fatal(err)
	}
	if len(done) != 1 || done[0].Id != 2 {
		t.Fatalf("filtered DONE = %#v", done)
	}
}

func TestMarkTodoDone_writesToTempDirOnly(t *testing.T) {
	dir := t.TempDir()
	swapDataDir(t, dir)
	writeSettingJSON(t, dir, 10)
	writeTodosJSON(t, dir, []models.Todos{
		{Id: 1, Title: "x", Status: models.PENDING, CreatedAt: "t"},
		{Id: 2, Title: "y", Status: models.PENDING, CreatedAt: "t"},
	})

	if err := MarkTodoDone(1); err != nil {
		t.Fatal(err)
	}
	onDisk := readTodosFromDisk(t, dir)
	if onDisk[0].Status != models.DONE || onDisk[1].Status != models.PENDING {
		t.Fatalf("after MarkTodoDone(1): %#v", onDisk)
	}
}

func TestDeleteTodo(t *testing.T) {
	dir := t.TempDir()
	swapDataDir(t, dir)
	writeSettingJSON(t, dir, 10)
	writeTodosJSON(t, dir, []models.Todos{
		{Id: 1, Title: "keep", Status: models.PENDING, CreatedAt: "t"},
		{Id: 2, Title: "drop", Status: models.PENDING, CreatedAt: "t"},
	})

	if err := DeleteTodo(2); err != nil {
		t.Fatal(err)
	}
	onDisk := readTodosFromDisk(t, dir)
	if len(onDisk) != 1 || onDisk[0].Id != 1 {
		t.Fatalf("after DeleteTodo(2): %#v", onDisk)
	}
}

func TestEditTodo(t *testing.T) {
	dir := t.TempDir()
	swapDataDir(t, dir)
	writeSettingJSON(t, dir, 100)
	writeTodosJSON(t, dir, []models.Todos{
		{Id: 5, Title: "old", Status: models.PENDING, CreatedAt: "t"},
	})

	if err := EditTodo(5, "new title"); err != nil {
		t.Fatal(err)
	}
	onDisk := readTodosFromDisk(t, dir)
	if len(onDisk) != 1 || onDisk[0].Title != "new title" {
		t.Fatalf("after EditTodo: %#v", onDisk)
	}
}
