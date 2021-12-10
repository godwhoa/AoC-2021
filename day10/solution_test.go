package day10

import (
	"os"
	"testing"
)

func TestScoreSyntax_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	input := ParseInput(f)
	got := TotalSyntaxErrorScore(input)
	want := 26397
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestScoreSyntax_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	input := ParseInput(f)
	got := TotalSyntaxErrorScore(input)
	want := 339411
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCompleteionScore_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	input := ParseInput(f)
	got := TotalCompletionScore(input)
	want := 288957
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCompleteionScore_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	input := ParseInput(f)
	got := TotalCompletionScore(input)
	want := 2289754624
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
