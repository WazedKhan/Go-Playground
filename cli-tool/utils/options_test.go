package utils

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestAvailableCommands_printsEachCommand(t *testing.T) {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w

	AvailableCommands()

	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	os.Stdout = old

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatal(err)
	}
	_ = r.Close()

	out := buf.String()
	for _, line := range []string{
		HelpCommand,
		QuitCommand,
		AddCommand,
		ListCommand,
		DeleteCommand,
		DoneCommand,
	} {
		if !strings.Contains(out, line) {
			t.Fatalf("output missing %q:\n%s", line, out)
		}
	}
}
