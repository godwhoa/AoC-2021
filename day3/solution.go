package day3

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func ToInt(s string) int {
	i, _ := strconv.ParseInt(s, 2, 64)
	return int(i)
}

func runesToInt(rs []rune) int {
	return ToInt(string(rs))
}

func PowerConsumption(mat [][]rune) int {
	bits := len(mat[0])
	var zeros = make([]int, bits)
	var ones = make([]int, bits)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == '0' {
				zeros[j]++
			} else {
				ones[j]++
			}
		}
	}

	var gammaS strings.Builder
	var epsilonS strings.Builder
	for i := 0; i < len(zeros); i++ {
		if zeros[i] < ones[i] {
			gammaS.WriteRune('1')
			epsilonS.WriteRune('0')
		} else {
			gammaS.WriteRune('0')
			epsilonS.WriteRune('1')
		}
	}
	return ToInt(gammaS.String()) * ToInt(epsilonS.String())
}

func MostCommonAtN(mat [][]rune, n int) rune {
	zeros, ones := 0, 0
	for i := 0; i < len(mat); i++ {
		if mat[i][n] == '0' {
			zeros++
		} else {
			ones++
		}
	}
	if zeros > ones {
		return '0'
	}
	return '1'
}

func Invert(r rune) rune {
	if r == '0' {
		return '1'
	}
	return '0'
}

func Oxygen(mat [][]rune) int {
	n := 0
	for len(mat) > 1 {
		common := MostCommonAtN(mat, n)
		mat = FilterMat(mat, func(row []rune) bool {
			return row[n] == common
		})
		n++
	}
	return runesToInt(mat[0])
}

func CO2(mat [][]rune) int {
	n := 0
	for len(mat) > 1 {
		leastCommon := MostCommonAtN(mat, n)
		mat = FilterMat(mat, func(row []rune) bool {
			return row[n] != leastCommon
		})
		n++
	}
	return runesToInt(mat[0])
}

func FilterMat(mat [][]rune, filter func([]rune) bool) [][]rune {
	var newmat [][]rune
	for i := 0; i < len(mat); i++ {
		if filter(mat[i]) {
			newmat = append(newmat, mat[i])
		}
	}
	return newmat
}

func ParseInput(input io.ReadCloser) [][]rune {
	var col [][]rune
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var row []rune
		for _, c := range scanner.Text() {
			row = append(row, c)
		}
		col = append(col, row)
	}
	return col
}
