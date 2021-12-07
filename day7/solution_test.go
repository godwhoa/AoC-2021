package day7

import (
	"os"
	"testing"
)

func TestMinFuelFixedStep_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	input, low, high := ParseInput(f)
	got := MinFuel(input, low, high, false)
	want := 37
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestMinFuelFixedStep_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	input, low, high := ParseInput(f)
	got := MinFuel(input, low, high, false)
	want := 353800
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestMinFuelProgressiveStep_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	input, low, high := ParseInput(f)
	got := MinFuel(input, low, high, true)
	want := 168
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestMinFuelProgressiveStep_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	input, low, high := ParseInput(f)
	got := MinFuel(input, low, high, true)
	want := 98119739
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
