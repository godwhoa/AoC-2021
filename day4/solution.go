package day4

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ArrayInSet(array []int, set map[int]struct{}) bool {
	for _, i := range array {
		if _, ok := set[i]; !ok {
			return false
		}
	}
	return true
}

func SumUnmarkedInBoard(marked map[int]struct{}, board [][]int) int {
	sum := 0
	for _, row := range board {
		for _, i := range row {
			if _, ok := marked[i]; !ok {
				sum += i
			}
		}
	}
	return sum
}

func AnyStrike(draws map[int]struct{}, board [][]int) bool {
	columns := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		columns[i] = make([]int, len(board[0]))
	}

	for i, row := range board {
		if ArrayInSet(row, draws) {
			return true
		}
		for j, value := range row {
			columns[j][i] = value
		}
	}

	for _, column := range columns {
		if ArrayInSet(column, draws) {
			return true
		}
	}

	return false
}

func FirstToWinScore(draws []int, boards [][][]int) int {
	set := make(map[int]struct{})
	for _, draw := range draws {
		set[draw] = struct{}{}
		for _, board := range boards {
			if AnyStrike(set, board) {
				return SumUnmarkedInBoard(set, board) * draw
			}
		}
	}
	return 0
}

func LastToWinScore(draws []int, boards [][][]int) int {
	won := make(map[int]struct{})
	score := 0
	set := make(map[int]struct{})
	for _, draw := range draws {
		set[draw] = struct{}{}
		for index, board := range boards {
			_, ok := won[index]
			if AnyStrike(set, board) && !ok {
				won[index] = struct{}{}
				score = SumUnmarkedInBoard(set, board) * draw
			}
		}
	}
	return score
}

func ParseInput(input io.ReadCloser) ([]int, [][][]int) {
	var rows [][]int
	scanner := bufio.NewScanner(input)

	scanner.Scan()
	draws := ToIntSlice(strings.Split(scanner.Text(), ","))

	for scanner.Scan() {
		line := scanner.Text()
		textRow := strings.Split(line, " ")
		if len(textRow) == 0 {
			continue
		}
		slice := ToIntSlice(textRow)
		if len(slice) > 0 {
			rows = append(rows, slice)
		}
	}
	boards := make([][][]int, (len(rows) / 5))
	i := 0
	for j := 0; j < len(rows); j++ {
		boards[i] = append(boards[i], rows[j])
		if (j+1)%5 == 0 {
			i++
		}
	}
	return draws, boards
}

func ToIntSlice(ss []string) []int {
	var row []int
	for _, rnum := range ss {
		num, err := strconv.Atoi(rnum)
		if err == nil {
			row = append(row, num)
		}
	}
	return row
}
