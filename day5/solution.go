package day5

import (
	"fmt"
	"io"
)

type Point struct {
	X, Y int
}

type Line struct {
	From Point
	To   Point
}

func (l Line) IsVertical() bool {
	return l.From.Y == l.To.Y
}

func (l Line) IsHorizontal() bool {
	return l.From.X == l.To.X
}

func (l Line) Delta() Point {
	return Point{l.To.X - l.From.X, l.To.Y - l.From.Y}
}

// Bresenham's line drawing algorithm
func Brrrseeennhaaammm(line Line) []Point {
	var points []Point
	d := line.Delta()
	if d.X < 0 {
		d.X = -d.X
	}
	if d.Y < 0 {
		d.Y = -d.Y
	}
	sx, sy := 1, 1
	if line.From.X > line.To.X {
		sx = -1
	}
	if line.From.Y > line.To.Y {
		sy = -1
	}
	err := d.X - d.Y
	for {
		points = append(points, Point{line.From.X, line.From.Y})
		if line.From.X == line.To.X && line.From.Y == line.To.Y {
			break
		}
		e2 := 2 * err
		if e2 > -d.Y {
			err -= d.Y
			line.From.X += sx
		}

		if e2 < d.X {
			err += d.X

			line.From.Y += sy
		}
	}
	return points
}

func TotalOverlappingPoints(lines []Line, dag bool) int {
	grid := make(map[Point]int)
	for _, line := range lines {
		if !(line.IsVertical() || line.IsHorizontal()) && !dag {
			continue
		}
		for _, point := range Brrrseeennhaaammm(line) {
			grid[point]++
		}
	}
	var count int
	for _, v := range grid {
		if v > 1 {
			count++
		}
	}
	return count
}

func ParseInput(input io.ReadCloser) []Line {
	var lines []Line
	for {
		var line Line
		_, err := fmt.Fscanf(input, "%d,%d -> %d,%d\n", &line.From.X, &line.From.Y, &line.To.X, &line.To.Y)
		if err == io.EOF {
			break
		}
		lines = append(lines, line)
	}
	return lines
}
