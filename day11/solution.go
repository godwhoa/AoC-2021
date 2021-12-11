package day11

import (
	"bufio"
	"fmt"
	"io"
)

func makeMatrix(n, m int) [][]int {
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, m)
	}
	return mat
}

func Adj(x int, y int, fn func(int, int)) {
	// top, bottom, left, right, topleft, topright, bottomleft, bottomright
	dxs := []int{0, 0, -1, 1, -1, 1, -1, 1}
	dys := []int{-1, 1, 0, 0, -1, -1, 1, 1}
	for i := 0; i < len(dxs); i++ {
		dx, dy := dxs[i], dys[i]
		fn(x+dx, y+dy)
	}
}

func Flash(mat, flashed [][]int, fx, fy int) {
	if mat[fy][fx] <= 9 || flashed[fy][fx] > 0 {
		return
	}
	inBounds := func(x, y int) bool {
		return x >= 0 && x < len(mat[0]) && y >= 0 && y < len(mat)
	}
	flashed[fy][fx] = 1

	Adj(fx, fy, func(x int, y int) {
		if !inBounds(x, y) {
			return
		}
		mat[y][x] += 1
		if mat[y][x] > 9 {
			Flash(mat, flashed, x, y)
		}
	})
}

func CountAndReset(mat, flashed [][]int) (count int) {
	for y := 0; y < len(mat); y++ {
		for x := 0; x < len(mat[y]); x++ {
			if flashed[y][x] == 1 {
				mat[y][x] = 0
				count++
			}
		}
	}
	return
}

func PrintMatrix(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		fmt.Println(mat[i])
	}
	fmt.Println("")
}

func Simulate(mat [][]int, steps int, firstFlash bool) (count int, allFlash int) {
	for i := 0; i < steps; i++ {
		flashed := makeMatrix(len(mat), len(mat[0]))
		for y := 0; y < len(mat); y++ {
			for x := 0; x < len(mat[y]); x++ {
				mat[y][x] += 1
				Flash(mat, flashed, x, y)
			}
		}
		flashes := CountAndReset(mat, flashed)
		if flashes == len(mat)*len(mat[0]) && firstFlash {
			allFlash = i + 1
			return
		}
		if firstFlash {
			steps++
		}
		count += flashes
	}
	return
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
