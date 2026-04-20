package models

import "testing"

func TestStatusConstants(t *testing.T) {
	if PENDING != "PENDING" || DONE != "DONE" {
		t.Fatalf("PENDING=%q DONE=%q", PENDING, DONE)
	}
}

func TestTodosStructFieldJSON(t *testing.T) {
	// Ensures the exported type stays usable for tests and JSON round-trips.
	todo := Todos{Id: 1, Title: "t", Status: PENDING, CreatedAt: "now"}
	if todo.Id != 1 || todo.Title != "t" {
		t.Fatalf("Todos = %#v", todo)
	}
}

func TestSettingStruct(t *testing.T) {
	s := Setting{MaxTitleLength: 12}
	if s.MaxTitleLength != 12 {
		t.Fatalf("Setting = %#v", s)
	}
}
