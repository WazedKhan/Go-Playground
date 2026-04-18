package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"cli-tool/models"
	"cli-tool/repository"
)

func TestDisplayTodos_empty(t *testing.T) {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	DisplayTodos(nil)

	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	_ = r.Close()

	if !bytes.Contains(buf.Bytes(), []byte("No TODOs found")) {
		t.Fatalf("output: %s", buf.String())
	}
}

func TestDisplayTodos_nonEmpty_usesMaxTitleLength(t *testing.T) {
	dir := t.TempDir()
	prev := repository.DataDir
	repository.DataDir = dir
	t.Cleanup(func() { repository.DataDir = prev })

	setting, err := json.MarshalIndent(models.Setting{MaxTitleLength: 8}, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "setting.json"), setting, 0644); err != nil {
		t.Fatal(err)
	}

	data := []models.Todos{
		{Id: 1, Title: "short", Status: models.PENDING, CreatedAt: "2026-01-01"},
	}

	old := os.Stdout
	pr, pw, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = pw

	DisplayTodos(data)

	if err := pw.Close(); err != nil {
		t.Fatal(err)
	}
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, pr)
	_ = pr.Close()

	out := buf.String()
	for _, needle := range []string{"1 TODO", "short", "PENDING", "2026-01-01"} {
		if !strings.Contains(out, needle) {
			t.Fatalf("output missing %q:\n%s", needle, out)
		}
	}
}
