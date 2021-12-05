package day3

import (
	"os"
	"reflect"
	"testing"
)

func TestPowerConsumption_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	defer f.Close()
	got := PowerConsumption(ParseInput(f))
	want := 4138664
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPowerConsumption_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	defer f.Close()
	got := PowerConsumption(ParseInput(f))
	want := 198
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestOxygen_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	defer f.Close()
	got := Oxygen(ParseInput(f))
	want := 23
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCO2_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	defer f.Close()
	got := CO2(ParseInput(f))
	want := 10
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPartTwo(t *testing.T) {
	f, _ := os.Open("input.txt")
	defer f.Close()
	input := ParseInput(f)
	oxygen := Oxygen(input[:])
	co2 := CO2(input)
	got := oxygen * co2
	want := 4273224
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
