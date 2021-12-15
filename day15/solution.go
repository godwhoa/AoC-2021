package day15

import (
	"bufio"
	"fmt"
	"io"
	"math"

	"github.com/fatih/color"
)

type Node struct {
	x, y   int
	cost   int
	chosen bool
}

func Adj(x int, y int, fn func(int, int)) {
	// top, bottom, left, right
	dxs := []int{0, 0, -1, 1}
	dys := []int{-1, 1, 0, 0}
	for i := 0; i < len(dxs); i++ {
		dx, dy := dxs[i], dys[i]
		fn(x+dx, y+dy)
	}
}

func Draw(grid [][]*Node) {
	chosenColor := color.New(color.FgHiWhite)
	regularColor := color.New(color.FgHiBlack)
	for _, row := range grid {
		for _, node := range row {
			if node.chosen {
				chosenColor.Print(node.cost)
			} else {
				regularColor.Print(node.cost)
			}
		}
		fmt.Println()
	}
}

func LowestRisk(grid [][]*Node) ([]*Node, int) {
	inBounds := func(x, y int) bool {
		return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
	}
	start := grid[0][0]
	goal := grid[len(grid)-1][len(grid[0])-1]
	openSet := NewMinHeap(start)

	gScore := make(map[*Node]int)
	gScore[start] = 0
	getgScore := func(node *Node) int {
		if score, exits := gScore[node]; exits {
			return score
		}
		return math.MaxInt
	}

	cameFrom := make(map[*Node]*Node)
	constructPath := func(current *Node) ([]*Node, int) {
		path := []*Node{current}
		risk := current.cost
		current.chosen = true
		for {
			current = cameFrom[current]
			if current == nil {
				break
			}
			current.chosen = true
			path = append(path, current)
			risk += current.cost
		}
		return path, risk - start.cost
	}

	for (*openSet).Len() > 0 {
		current := openSet.PopNode()
		if current == goal {
			return constructPath(current)
		}

		Adj(current.x, current.y, func(nx int, ny int) {
			if !inBounds(nx, ny) {
				return
			}
			neighbor := grid[ny][nx]
			tentativegScore := getgScore(current) + neighbor.cost
			if tentativegScore < getgScore(neighbor) {
				cameFrom[neighbor] = current
				gScore[neighbor] = tentativegScore
				openSet.UpsertNode(neighbor, tentativegScore+neighbor.cost)
			}
		})
	}
	return []*Node{}, 0
}

func ScaleGrid(grid [][]*Node, scaleX, scaleY int) [][]*Node {
	var scaled [][]*Node
	for sy := 0; sy < scaleY; sy++ {
		for y, row := range grid {
			var scaledRow []*Node
			for sx := 0; sx < scaleX; sx++ {
				for x, node := range row {
					node = &Node{
						x:    x + (sx * len(row)),
						y:    y + (sy * len(grid)),
						cost: node.cost + sy + sx,
					}
					if node.cost > 9 {
						node.cost -= 9
					}
					scaledRow = append(scaledRow, node)
				}
			}
			scaled = append(scaled, scaledRow)
		}
	}
	return scaled
}

func ParseInput(input io.ReadCloser) [][]*Node {
	var grid [][]*Node
	scanner := bufio.NewScanner(input)
	y := 0
	for scanner.Scan() {
		var row []*Node
		for x, c := range scanner.Text() {
			row = append(row, &Node{
				x: x, y: y,
				cost: int(c - '0'),
			})
		}
		grid = append(grid, row)
		y++
	}
	return grid
}
