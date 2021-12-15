package day15

import (
	"os"
	"testing"
)

func TestLowestRisk_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	grid := ParseInput(f)
	_, got := LowestRisk(grid)
	want := 40
	Draw(grid)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestLowestRisk_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	grid := ParseInput(f)
	_, got := LowestRisk(grid)
	Draw(grid)
	want := 366
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestLowestRisk_SampleScaled(t *testing.T) {
	f, _ := os.Open("sample.txt")
	grid := ScaleGrid(ParseInput(f), 5, 5)
	_, got := LowestRisk(grid)
	want := 315
	Draw(grid)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestLowestRisk_InputScaled(t *testing.T) {
	f, _ := os.Open("input.txt")
	grid := ScaleGrid(ParseInput(f), 5, 5)
	_, got := LowestRisk(grid)
	want := 2829
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
