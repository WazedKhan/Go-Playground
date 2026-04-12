package main

import "testing"

func TestPeriMeter(t *testing.T) {
	got := PeriMeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("want %0.2f but got %0.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	got := Area(12.0, 6.0)
	want := 72.0

	if got != want {
		t.Errorf("want %0.2f but got %0.2f", want, got)
	}
}
