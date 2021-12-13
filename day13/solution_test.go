package day13

import (
	"os"
	"testing"
)

func TestDotsAferFirstFold_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	dots, folds, bounds := ParseInput(f)
	got, _ := DotsAferFolds(dots, folds[:1], bounds)
	want := 17
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestDotsAferFirstFold_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	dots, folds, bounds := ParseInput(f)
	got, _ := DotsAferFolds(dots, folds[:1], bounds)
	want := 647
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestDotsAfterFolds_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	dots, folds, bounds := ParseInput(f)
	got, drawing := DotsAferFolds(dots, folds, bounds)
	want := 93
	t.Logf("code:\n%s\n", drawing)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
