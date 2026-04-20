package repository

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"cli-tool/models"
)

func swapDataDir(t *testing.T, dir string) {
	t.Helper()
	prev := DataDir
	DataDir = dir
	t.Cleanup(func() { DataDir = prev })
}

func TestGetTODOs_missingFile_returnsEmpty(t *testing.T) {
	dir := t.TempDir()
	swapDataDir(t, dir)

	got := GetTODOs()
	if len(got) != 0 {
		t.Fatalf("GetTODOs() = %#v, want empty", got)
	}
}

func TestGetTODOs_invalidJSON_returnsEmpty(t *testing.T) {
	dir := t.TempDir()
	swapDataDir(t, dir)
	path := filepath.Join(dir, "todos.json")
	if err := os.WriteFile(path, []byte("{not json"), 0644); err != nil {
		t.Fatal(err)
	}

	got := GetTODOs()
	if len(got) != 0 {
		t.Fatalf("GetTODOs() = %#v, want empty", got)
	}
}

func TestGetTODOs_validFile(t *testing.T) {
	dir := t.TempDir()
	swapDataDir(t, dir)
	seed := []models.Todos{
		{Id: 1, Title: "a", Status: models.PENDING, CreatedAt: "t"},
	}
	raw, err := json.MarshalIndent(seed, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "todos.json"), raw, 0644); err != nil {
		t.Fatal(err)
	}

	got := GetTODOs()
	if len(got) != 1 || got[0].Title != "a" {
		t.Fatalf("GetTODOs() = %#v", got)
	}
}

func TestGetMaxTitleLength_missingFile_returnsZero(t *testing.T) {
	dir := t.TempDir()
	swapDataDir(t, dir)

	if n := GetMaxTitleLength(); n != 0 {
		t.Fatalf("GetMaxTitleLength() = %d, want 0", n)
	}
}

func TestGetMaxTitleLength_invalidJSON_returnsZero(t *testing.T) {
	dir := t.TempDir()
	swapDataDir(t, dir)
	if err := os.WriteFile(filepath.Join(dir, "setting.json"), []byte("x"), 0644); err != nil {
		t.Fatal(err)
	}

	if n := GetMaxTitleLength(); n != 0 {
		t.Fatalf("GetMaxTitleLength() = %d, want 0", n)
	}
}

func TestGetMaxTitleLength_validFile(t *testing.T) {
	dir := t.TempDir()
	swapDataDir(t, dir)
	s := models.Setting{MaxTitleLength: 42}
	raw, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "setting.json"), raw, 0644); err != nil {
		t.Fatal(err)
	}

	if n := GetMaxTitleLength(); n != 42 {
		t.Fatalf("GetMaxTitleLength() = %d, want 42", n)
	}
}
