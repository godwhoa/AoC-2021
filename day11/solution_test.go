package day11

import (
	"os"
	"testing"
)

func TestSimulate_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	mat := ParseInput(f)
	got, _ := Simulate(mat, 100, false)
	want := 1656
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimulateFirstFlash_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	mat := ParseInput(f)
	_, got := Simulate(mat, 10, true)
	want := 195
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimulate_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	mat := ParseInput(f)
	got, _ := Simulate(mat, 100, false)
	want := 1640
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimulateFirstFlash_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	mat := ParseInput(f)
	_, got := Simulate(mat, 10, true)
	want := 312
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
