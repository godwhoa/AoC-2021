package day12

import (
	"os"
	"testing"
)

func TestFindPaths_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	input := ParseInput(f)
	paths := FindPaths(input, "start", "end")
	got := len(paths)
	want := 10
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFindPathsTwice_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	input := ParseInput(f)
	paths := FindPathsCanVisitTwice(input, "start", "end")
	got := len(paths)
	want := 36
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFindPathsTwice_Sample_1(t *testing.T) {
	f, _ := os.Open("sample_1.txt")
	input := ParseInput(f)
	paths := FindPathsCanVisitTwice(input, "start", "end")
	got := len(paths)
	want := 103
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFindPathsTwice_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	input := ParseInput(f)
	paths := FindPathsCanVisitTwice(input, "start", "end")
	got := len(paths)
	want := 149385
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFindPaths_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	input := ParseInput(f)
	paths := FindPaths(input, "start", "end")
	got := len(paths)
	want := 5254
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
