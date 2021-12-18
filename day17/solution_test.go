package main

import (
	"testing"
)

func TestBruteForceHighestY(t *testing.T) {
	tests := []struct {
		name   string
		target Target
		want   int
	}{
		{
			"Sample",
			Target{20, 30, -10, -5},
			45,
		},
		{
			"Input",
			Target{241, 273, -97, -63},
			4656,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BruteForceHighestY(tt.target); got != tt.want {
				t.Errorf("BruteForceHighestY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBruteForceDistinctVectors(t *testing.T) {
	tests := []struct {
		name      string
		target    Target
		wantCount int
	}{
		{
			"Sample",
			Target{20, 30, -10, -5},
			112,
		},
		{
			"Input",
			Target{241, 273, -97, -63},
			1908,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := BruteForceDistinctVectors(tt.target); gotCount != tt.wantCount {
				t.Errorf("BruteForceDistinctVectors() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
