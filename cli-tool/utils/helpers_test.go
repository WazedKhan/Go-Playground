package utils

import "testing"

func TestExtractQuotedTitle(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{"ok", `todo edit 1 "hello world"`, "hello world", false},
		{"inner quotes", `x "a \" b"`, `a \" b`, false},
		{"no open", `todo edit 1 noquotes`, "", true},
		{"only one quote", `"only`, "", true},
		{"empty quoted segment", `todo ""`, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := ExtractQuotedTitle(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Fatalf("ExtractQuotedTitle(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
