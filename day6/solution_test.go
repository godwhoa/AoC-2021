package day6

import (
	"os"
	"reflect"
	"testing"
)

func TestSimulate_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	input := ParseInput(f)
	got := Simulate(input, 18)
	want := 26
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSimulate_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	input := ParseInput(f)
	got := Simulate(input, 80)
	want := 379414
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSimulate_Part2_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	input := ParseInput(f)
	got := Simulate(input, 256)
	want := 26984457539
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSimulate_Part2_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	input := ParseInput(f)
	got := Simulate(input, 256)
	want := 1705008653296
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
