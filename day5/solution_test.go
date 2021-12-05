package day5

import (
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestParseInput(t *testing.T) {
	f, _ := os.Open("sample.txt")
	got := ParseInput(f)[0]
	want := Line{From: Point{X: 0, Y: 9}, To: Point{X: 5, Y: 9}}
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTotalOverlappingPoints_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	got := TotalOverlappingPoints(ParseInput(f), false)
	want := 5
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTotalOverlappingPoint_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	got := TotalOverlappingPoints(ParseInput(f), false)
	want := 6841
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTotalOverlappingPoint_Dag_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	got := TotalOverlappingPoints(ParseInput(f), true)
	want := 12
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestTotalOverlappingPoint_Dag_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	got := TotalOverlappingPoints(ParseInput(f), true)
	want := 19258
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSilly(t *testing.T) {
	line := Line{From: Point{X: 1, Y: 1}, To: Point{X: 1, Y: 3}}
	got := Brrrseeennhaaammm(line)
	spew.Dump(got)
}
