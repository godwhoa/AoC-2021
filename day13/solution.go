package day13

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func makeGrid(bounds Point) [][]int {
	grid := make([][]int, bounds.y+1)
	for y := 0; y <= bounds.y; y++ {
		grid[y] = make([]int, bounds.x+1)
	}
	return grid
}

func FoldUp(grid [][]int, offset int) [][]int {
	for x := 0; x < len(grid[offset]); x++ {
		for y := len(grid) - 1; y > offset; y-- {
			if grid[offset*2-y][x] == 0 {
				grid[offset*2-y][x] = grid[y][x]
			}
		}
	}
	return grid[:offset]
}

func FoldLeft(grid [][]int, offset int) [][]int {
	for y := 0; y < len(grid); y++ {
		for x := len(grid[y]) - 1; x > offset; x-- {
			if grid[y][offset*2-x] == 0 {
				grid[y][offset*2-x] = grid[y][x]
			}
		}
		grid[y] = grid[y][:offset]
	}
	return grid
}

func CountDots(grid [][]int) (count int, code string) {
	var drawing strings.Builder
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 1 {
				drawing.WriteRune('#')
				count++
			} else {
				drawing.WriteRune(' ')
			}
		}
		drawing.WriteRune('\n')
	}
	code = drawing.String()
	return
}

func DotsAferFolds(dots, folds []Point, bounds Point) (int, string) {
	grid := makeGrid(bounds)
	for _, dot := range dots {
		grid[dot.y][dot.x] = 1
	}
	for _, fold := range folds {
		if fold.y > 0 {
			grid = FoldUp(grid, fold.y)
		}
		if fold.x > 0 {
			grid = FoldLeft(grid, fold.x)
		}
	}
	return CountDots(grid)
}

func ParseInput(input io.ReadCloser) ([]Point, []Point, Point) {
	var dots []Point
	var folds []Point
	var bounds Point
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		raw := scanner.Text()
		if strings.Contains(raw, "y=") {
			folds = append(folds, Point{y: toInt(strings.Split(raw, "=")[1])})
			continue
		}
		if strings.Contains(raw, "x=") {
			folds = append(folds, Point{x: toInt(strings.Split(raw, "=")[1])})
			continue
		}
		if strings.Contains(raw, ",") {
			parts := strings.Split(raw, ",")
			point := Point{
				x: toInt(parts[0]),
				y: toInt(parts[1]),
			}
			bounds.x, bounds.y = max(bounds.x, point.x), max(bounds.y, point.y)
			dots = append(dots, point)
		}

	}
	return dots, folds, bounds
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func toInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
