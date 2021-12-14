package day14

import (
	"os"
	"testing"
)

func TestSimulateTen_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	letterFreq, pairFreq, insertionPairs := ParseInput(f)
	got := Simulate(letterFreq, pairFreq, insertionPairs, 10)
	want := 1588
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimulateTen_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	letterFreq, pairFreq, insertionPairs := ParseInput(f)
	got := Simulate(letterFreq, pairFreq, insertionPairs, 10)
	want := 2975
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimulateForty_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	letterFreq, pairFreq, insertionPairs := ParseInput(f)
	got := Simulate(letterFreq, pairFreq, insertionPairs, 40)
	want := 2188189693529
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSimulateForty_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	letterFreq, pairFreq, insertionPairs := ParseInput(f)
	got := Simulate(letterFreq, pairFreq, insertionPairs, 40)
	want := 3015383850689
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
