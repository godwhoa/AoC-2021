package day4

import (
	"os"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	f, _ := os.Open("input.txt")
	drawsGot, boardsGot := ParseInput(f)

	drawsWant := []int{63, 23, 2, 65, 55, 94, 38, 20, 22, 39, 5, 98, 9, 60, 80, 45, 99, 68, 12, 3, 6, 34, 64, 10, 70, 69, 95, 96, 83, 81, 32, 30, 42, 73, 52, 48, 92, 28, 37, 35, 54, 7, 50, 21, 74, 36, 91, 97, 13, 71, 86, 53, 46, 58, 76, 77, 14, 88, 78, 1, 33, 51, 89, 26, 27, 31, 82, 44, 61, 62, 75, 66, 11, 93, 49, 43, 85, 0, 87, 40, 24, 29, 15, 59, 16, 67, 19, 72, 57, 41, 8, 79, 56, 4, 18, 17, 84, 90, 47, 25}
	if !reflect.DeepEqual(drawsGot, drawsWant) {
		t.Errorf("draws got %v, want %v", drawsGot, drawsWant)
	}
	boardsWant := []int{25, 29, 78, 57, 69}
	if !reflect.DeepEqual(boardsGot[0][0], boardsWant) {
		t.Errorf("boards got %v, want %v", boardsGot, boardsWant)
	}
}

func TestFinalScore_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	draws, boards := ParseInput(f)
	score := FirstToWinScore(draws, boards)
	if score != 4512 {
		t.Errorf("score got %v, want 4512", score)
	}
}

func TestFirstToWinScore_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	draws, boards := ParseInput(f)
	score := FirstToWinScore(draws, boards)
	if score != 63424 {
		t.Errorf("score got %v, want 63424", score)
	}
}

func TestLastToWinScore_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	draws, boards := ParseInput(f)
	score := LastToWinScore(draws, boards)
	if score != 23541 {
		t.Errorf("score got %v, want 23541", score)
	}
}
