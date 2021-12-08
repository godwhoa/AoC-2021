package day8

import (
	"os"
	"testing"
)

func TestCountSpecialDigits(t *testing.T) {
	tests := []struct {
		name      string
		filepath  string
		wantCount int
	}{
		{"Sample Part 1", "sample.txt", 26},
		{"Input Part 1", "input.txt", 255},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, _ := os.Open(tt.filepath)
			defer f.Close()
			if gotCount := CountSpecialDigits(f); gotCount != tt.wantCount {
				t.Errorf("CountSpecialDigits() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func EqualToBag(mappings map[string]int, bag *Bag) bool {
	for raw, digit := range mappings {
		if bag.Digit(NewPattern(raw)) != digit {
			return false
		}
	}
	return true
}

func TestSumOfDisplayOutputs_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	got := SumOfDisplayOutputs(f)
	want := 61229
	if got != want {
		t.Errorf("SumOfDisplayOutputs() = %v, want %v", got, want)
	}
}

func TestSumOfDisplayOutputs_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	got := SumOfDisplayOutputs(f)
	want := 982158
	if got != want {
		t.Errorf("SumOfDisplayOutputs() = %v, want %v", got, want)
	}
}

func TestDeducePatterns(t *testing.T) {
	want := map[string]int{
		"acedgfb": 8,
		"cdfbe":   5,
		"gcdfa":   2,
		"fbcad":   3,
		"dab":     7,
		"cefabd":  9,
		"cdfgeb":  6,
		"eafb":    4,
		"cagedb":  0,
		"ab":      1,
	}
	rawPatterns := []string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"}
	if got := DeducePatterns(rawPatterns); !EqualToBag(want, got) {
		t.Errorf("DeducePatterns() = %v, want %v", got.ptod, want)
	}
}
