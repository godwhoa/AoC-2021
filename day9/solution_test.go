package day9

import (
	"os"
	"testing"
)

func TestTotalRiskLevel_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	mat := ParseInput(f)
	got := TotalRiskLevel(mat)
	want := 15
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestBasins_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	mat := ParseInput(f)
	got := Basins(mat)
	want := 1134
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestBasins_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	mat := ParseInput(f)
	got := Basins(mat)
	want := 736920
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestTotalRiskLevel_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	mat := ParseInput(f)
	got := TotalRiskLevel(mat)
	want := 539
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
