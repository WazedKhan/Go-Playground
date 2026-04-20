package internal

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

func captureStdout(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	_ = w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	_ = r.Close()
	return buf.String()
}

func TestHandleCommands_validationMessages(t *testing.T) {
	out := captureStdout(func() {
		HandleCommands("")
		HandleCommands("one")
		HandleCommands("todo")
		HandleCommands("other todo list")
	})
	if !strings.Contains(out, "Invalid command") {
		t.Fatalf("missing invalid hint:\n%s", out)
	}
	if !strings.Contains(out, "Unknown command") {
		t.Fatalf("missing unknown hint:\n%s", out)
	}
}

func TestHandleCommands_help(t *testing.T) {
	_ = captureStdout(func() {
		HandleCommands("todo help")
		HandleCommands("todo -h")
	})
}

func TestHandleCommands_add_success(t *testing.T) {
	dir := t.TempDir()
	prev := repository.DataDir
	repository.DataDir = dir
	t.Cleanup(func() { repository.DataDir = prev })
	_ = os.WriteFile(filepath.Join(dir, "setting.json"), []byte(`{"maxTitleLength":100}`), 0644)

	out := captureStdout(func() {
		HandleCommands(`todo add "Ship feature"`)
	})
	if strings.Contains(out, "Error creating todo") {
		t.Fatalf("unexpected error output:\n%s", out)
	}
	if !strings.Contains(out, "TODO added successfully") {
		t.Fatalf("expected success message:\n%s", out)
	}
	raw, err := os.ReadFile(filepath.Join(dir, "todos.json"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(raw), "Ship feature") {
		t.Fatalf("todos.json: %s", raw)
	}
}

func TestHandleCommands_list(t *testing.T) {
	dir := t.TempDir()
	prev := repository.DataDir
	repository.DataDir = dir
	t.Cleanup(func() { repository.DataDir = prev })

	raw, err := json.MarshalIndent([]models.Todos{
		{Id: 1, Title: "a", Status: models.PENDING, CreatedAt: "t"},
	}, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "todos.json"), raw, 0644); err != nil {
		t.Fatal(err)
	}
	setRaw, _ := json.MarshalIndent(models.Setting{MaxTitleLength: 10}, "", "  ")
	if err := os.WriteFile(filepath.Join(dir, "setting.json"), setRaw, 0644); err != nil {
		t.Fatal(err)
	}

	out := captureStdout(func() { HandleCommands("todo list") })
	if !strings.Contains(out, "a") {
		t.Fatalf("expected list output, got:\n%s", out)
	}
}

func TestHandleCommands_doneAndDelete_invalidID(t *testing.T) {
	dir := t.TempDir()
	prev := repository.DataDir
	repository.DataDir = dir
	t.Cleanup(func() { repository.DataDir = prev })
	_ = os.WriteFile(filepath.Join(dir, "setting.json"), []byte(`{"maxTitleLength":10}`), 0644)

	out := captureStdout(func() {
		HandleCommands("todo done x")
		HandleCommands("todo delete y")
	})
	if !strings.Contains(out, "Invalid ID") {
		t.Fatalf("expected invalid ID messages:\n%s", out)
	}
}

func TestHandleCommands_filterList(t *testing.T) {
	dir := t.TempDir()
	prev := repository.DataDir
	repository.DataDir = dir
	t.Cleanup(func() { repository.DataDir = prev })

	seed := []models.Todos{
		{Id: 1, Title: "o", Status: models.PENDING, CreatedAt: "t"},
		{Id: 2, Title: "d", Status: models.DONE, CreatedAt: "t"},
	}
	raw, _ := json.MarshalIndent(seed, "", "  ")
	_ = os.WriteFile(filepath.Join(dir, "todos.json"), raw, 0644)
	_ = os.WriteFile(filepath.Join(dir, "setting.json"), []byte(`{"maxTitleLength":10}`), 0644)

	out := captureStdout(func() {
		HandleCommands("todo list --filter=" + models.PENDING)
	})
	if !strings.Contains(out, "o") || strings.Contains(out, "  2  ") {
		t.Fatalf("filter output:\n%s", out)
	}
}

func TestHandleCommands_edit_invalidID(t *testing.T) {
	out := captureStdout(func() {
		HandleCommands(`todo edit x "title"`)
	})
	if !strings.Contains(out, "Invalid ID") {
		t.Fatalf("output:\n%s", out)
	}
}

func TestHandleCommands_edit_missingQuotes(t *testing.T) {
	dir := t.TempDir()
	prev := repository.DataDir
	repository.DataDir = dir
	t.Cleanup(func() { repository.DataDir = prev })
	_ = os.WriteFile(filepath.Join(dir, "setting.json"), []byte(`{"maxTitleLength":10}`), 0644)

	out := captureStdout(func() {
		HandleCommands("todo edit 1 no-quotes")
	})
	if !strings.Contains(out, "Title must be in quotes") {
		t.Fatalf("output:\n%s", out)
	}
}

func TestHandleCommands_edit_success(t *testing.T) {
	dir := t.TempDir()
	prev := repository.DataDir
	repository.DataDir = dir
	t.Cleanup(func() { repository.DataDir = prev })

	seed := []models.Todos{{Id: 3, Title: "old", Status: models.PENDING, CreatedAt: "t"}}
	raw, _ := json.MarshalIndent(seed, "", "  ")
	_ = os.WriteFile(filepath.Join(dir, "todos.json"), raw, 0644)
	_ = os.WriteFile(filepath.Join(dir, "setting.json"), []byte(`{"maxTitleLength":50}`), 0644)

	_ = captureStdout(func() {
		HandleCommands(`todo edit 3 "new title"`)
	})
	got, err := os.ReadFile(filepath.Join(dir, "todos.json"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), "new title") {
		t.Fatalf("file after edit: %s", got)
	}
}

func TestHandleCommands_defaultBranch(t *testing.T) {
	out := captureStdout(func() {
		HandleCommands("todo not-a-real-subcommand")
	})
	if !strings.Contains(out, "Unknown command") {
		t.Fatalf("output:\n%s", out)
	}
}
