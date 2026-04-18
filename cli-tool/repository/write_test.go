package repository

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"cli-tool/models"
)

func readTodosFile(t *testing.T, dir string) []models.Todos {
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

func readSettingFile(t *testing.T, dir string) models.Setting {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(dir, "setting.json"))
	if err != nil {
		t.Fatal(err)
	}
	var s models.Setting
	if err := json.Unmarshal(raw, &s); err != nil {
		t.Fatal(err)
	}
	return s
}

func TestWriteSetting_roundTrip(t *testing.T) {
	dir := t.TempDir()
	swapDataDir(t, dir)

	if err := WriteSetting(models.Setting{MaxTitleLength: 99}); err != nil {
		t.Fatal(err)
	}
	if s := readSettingFile(t, dir); s.MaxTitleLength != 99 {
		t.Fatalf("setting = %#v", s)
	}
}

func TestUpdateMaxTitleLength_increasesOnly(t *testing.T) {
	dir := t.TempDir()
	swapDataDir(t, dir)
	if err := WriteSetting(models.Setting{MaxTitleLength: 10}); err != nil {
		t.Fatal(err)
	}
	if err := UpdateMaxTitleLength(5); err != nil {
		t.Fatal(err)
	}
	if s := readSettingFile(t, dir); s.MaxTitleLength != 10 {
		t.Fatalf("expected unchanged 10, got %#v", s)
	}

	if err := UpdateMaxTitleLength(20); err != nil {
		t.Fatal(err)
	}
	if s := readSettingFile(t, dir); s.MaxTitleLength != 20 {
		t.Fatalf("expected 20, got %#v", s)
	}
}

func TestAddTODO_assignsIDAndPersists(t *testing.T) {
	dir := t.TempDir()
	swapDataDir(t, dir)
	if err := WriteSetting(models.Setting{MaxTitleLength: 5}); err != nil {
		t.Fatal(err)
	}

	todo := models.Todos{Title: "first", Status: models.PENDING, CreatedAt: "t0"}
	if err := AddTODO(todo); err != nil {
		t.Fatal(err)
	}
	onDisk := readTodosFile(t, dir)
	if len(onDisk) != 1 || onDisk[0].Id != 1 || onDisk[0].Title != "first" {
		t.Fatalf("after first AddTODO: %#v", onDisk)
	}

	todo2 := models.Todos{Title: "second", Status: models.PENDING, CreatedAt: "t1"}
	if err := AddTODO(todo2); err != nil {
		t.Fatal(err)
	}
	onDisk = readTodosFile(t, dir)
	if len(onDisk) != 2 || onDisk[1].Id != 2 {
		t.Fatalf("after second AddTODO: %#v", onDisk)
	}
}

func TestUpdateTODO_overwritesFile(t *testing.T) {
	dir := t.TempDir()
	swapDataDir(t, dir)
	if err := WriteSetting(models.Setting{MaxTitleLength: 10}); err != nil {
		t.Fatal(err)
	}
	seed := []models.Todos{{Id: 1, Title: "x", Status: models.PENDING, CreatedAt: "t"}}
	raw, _ := json.MarshalIndent(seed, "", "  ")
	if err := os.WriteFile(filepath.Join(dir, "todos.json"), raw, 0644); err != nil {
		t.Fatal(err)
	}

	next := []models.Todos{{Id: 1, Title: "renamed", Status: models.DONE, CreatedAt: "t"}}
	if err := UpdateTODO(next); err != nil {
		t.Fatal(err)
	}
	got := readTodosFile(t, dir)
	if len(got) != 1 || got[0].Title != "renamed" || got[0].Status != models.DONE {
		t.Fatalf("UpdateTODO result: %#v", got)
	}
}
