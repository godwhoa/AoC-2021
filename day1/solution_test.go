package day1

import (
	"os"
	"testing"
)

func TestCountIncreasing(t *testing.T) {
	f, _ := os.Open("input.txt")
	input := ParseInput(f)
	if CountIncreasing(input) != 1316 {
		t.Error("Expected 1316, got ", CountIncreasing(input))
	}
}

func TestWindowIncreasing(t *testing.T) {
	f, _ := os.Open("input.txt")
	input := ParseInput(f)
	got := SlidingWindowIncreasing(input)
	want := 1344
	if got != want {
		t.Error("Expected", want, "got", got)
	}
}
