package day9

import (
	"bufio"
	"io"
	"math"
	"sort"
)

func TotalRiskLevel(mat [][]int) int {
	get := func(x, y int) int {
		if x < 0 || y < 0 || x >= len(mat[0]) || y >= len(mat) {
			return math.MaxInt
		}
		return mat[y][x]
	}

	var risk int
	for y := 0; y < len(mat); y++ {
		for x := 0; x < len(mat[y]); x++ {
			if mat[y][x] < min(get(x, y-1), get(x, y+1), get(x-1, y), get(x+1, y)) {
				risk += 1 + mat[y][x]
			}
		}
	}
	return risk
}

func Adj(x int, y int, fn func(int, int)) {
	// top, bottom, left, right
	dxs := []int{0, 0, -1, 1}
	dys := []int{-1, 1, 0, 0}
	for i := 0; i < 4; i++ {
		dx, dy := dxs[i], dys[i]
		fn(x+dx, y+dy)
	}
}

func FindBasinArea(mat [][]int, x, y int) int {
	get := func(x, y int) int {
		if x < 0 || y < 0 || x >= len(mat[0]) || y >= len(mat) || mat[y][x] == 9 {
			return -1
		}
		return mat[y][x]
	}

	area := 1
	mat[y][x] = 9

	Adj(x, y, func(ax int, ay int) {
		if get(ax, ay) == -1 || get(ax, ay) < get(x, y) {
			return
		}
		area += FindBasinArea(mat, ax, ay)
	})
	return area
}

func Basins(mat [][]int) int {
	get := func(x, y int) int {
		if x < 0 || y < 0 || x >= len(mat[0]) || y >= len(mat) {
			return math.MaxInt
		}
		return mat[y][x]
	}

	var areas sort.IntSlice

	for y := 0; y < len(mat); y++ {
		for x := 0; x < len(mat[y]); x++ {
			isLowlying := true
			Adj(x, y, func(ax int, ay int) {
				if get(x, y) > get(ax, ay) {
					isLowlying = false
					return
				}
			})
			if isLowlying {
				area := FindBasinArea(mat, x, y)
				areas = append(areas, area)
			}
		}
	}
	sort.Sort(sort.Reverse(areas))
	return areas[0] * areas[1] * areas[2]
}

func ParseInput(input io.ReadCloser) [][]int {
	var col [][]int
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var row []int
		for _, c := range scanner.Text() {
			row = append(row, int(c-'0'))
		}
		col = append(col, row)
	}
	return col
}

func min(values ...int) int {
	m := math.MaxInt
	for _, v := range values {
		if v < m {
			m = v
		}
	}
	return m
}
