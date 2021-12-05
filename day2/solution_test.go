package day2

import (
	"os"
	"testing"
)

func TestWhereDidIEndUp_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	commands := ParseInput(f)
	depth, forwardPos := WhereDidIEndUp(commands)
	if depth != 10 {
		t.Errorf("Expected 10, got %d", depth)
	}
	if forwardPos != 15 {
		t.Errorf("Expected 15, got %d", forwardPos)
	}
}

func TestWhereDidIEndUpWithAim_Sample(t *testing.T) {
	f, _ := os.Open("sample.txt")
	commands := ParseInput(f)
	_, depth, forwardPos := WhereDidIEndUpWithAim(commands)
	if depth*forwardPos != 900 {
		t.Errorf("Expected 900, got %d", depth*forwardPos)
	}
}

func TestWhereDidIEndUpWithAim_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	commands := ParseInput(f)
	_, depth, forwardPos := WhereDidIEndUpWithAim(commands)
	t.Logf("Depth: %d, forwardPos: %d Final: %d", depth, forwardPos, depth*forwardPos)
	if depth*forwardPos != 1741971043 {
		t.Errorf("Expected 1741971043, got %d", depth*forwardPos)
	}
}

func TestWhereDidIEndUp_Input(t *testing.T) {
	f, _ := os.Open("input.txt")
	commands := ParseInput(f)
	depth, forwardPos := WhereDidIEndUp(commands)
	t.Logf("Depth: %d, forwardPos: %d Final: %d", depth, forwardPos, depth*forwardPos)
	if depth*forwardPos != 1746616 {
		t.Errorf("Expected 1746616, got %d", depth*forwardPos)
	}
}
