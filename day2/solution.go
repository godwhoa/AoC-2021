package day2

import (
	"fmt"
	"io"
)

type Direction int

const (
	Up Direction = iota
	Down
	Forward
	NoOp
)

func ToDirection(str string) Direction {
	switch str {
	case "up":
		return Up
	case "down":
		return Down
	case "forward":
		return Forward
	default:
		return NoOp
	}
}

type Command struct {
	direction Direction
	steps     int
}

func WhereDidIEndUp(commands []Command) (depth int, forwardPos int) {
	for _, c := range commands {
		switch c.direction {
		case Up:
			depth -= c.steps
		case Down:
			depth += c.steps
		case Forward:
			forwardPos += c.steps
		}
	}
	return
}

func WhereDidIEndUpWithAim(commands []Command) (aim int, depth int, forwardPos int) {
	for _, c := range commands {
		switch c.direction {
		case Up:
			aim -= c.steps
		case Down:
			aim += c.steps
		case Forward:
			forwardPos += c.steps
			depth += aim * c.steps
		}
	}
	return
}

func ParseInput(input io.Reader) []Command {
	var commands []Command
	for {
		var dir string
		var steps int
		_, err := fmt.Fscanf(input, "%s %d", &dir, &steps)
		if err == io.EOF {
			break
		}
		commands = append(commands, Command{
			direction: ToDirection(dir),
			steps:     steps,
		})
	}
	return commands
}
