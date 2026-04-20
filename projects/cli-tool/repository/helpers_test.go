package repository

import (
	"testing"

	"cli-tool/models"
)

func TestGetNextID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		todos []models.Todos
		want  int64
	}{
		{"empty", nil, 1},
		{"empty slice", []models.Todos{}, 1},
		{"single", []models.Todos{{Id: 3}}, 4},
		{"max not first", []models.Todos{{Id: 2}, {Id: 10}, {Id: 5}}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := GetNextID(tt.todos); got != tt.want {
				t.Fatalf("GetNextID() = %d, want %d", got, tt.want)
			}
		})
	}
}
