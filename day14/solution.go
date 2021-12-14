package day14

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

func DiffMostAndLeast(freq map[rune]int) int {
	most, least := -1, math.MaxInt
	for _, count := range freq {
		most = max(most, count)
		least = min(least, count)
	}
	return most - least
}

func Simulate(
	letterFreq map[rune]int,
	pairFreq map[[2]rune]int,
	rules map[[2]rune]rune,
	steps int,
) int {
	for i := 0; i < steps; i++ {
		newPairFreq := make(map[[2]rune]int)
		for pair, count := range pairFreq {
			insert := rules[pair]
			pairOne := [2]rune{pair[0], insert}
			pairTwo := [2]rune{insert, pair[1]}
			newPairFreq[pairOne] += count
			newPairFreq[pairTwo] += count
			letterFreq[insert] += count
		}
		pairFreq = newPairFreq
	}
	return DiffMostAndLeast(letterFreq)
}

func ParseInput(input io.ReadCloser) (map[rune]int, map[[2]rune]int, map[[2]rune]rune) {
	scanner := bufio.NewScanner(input)
	scanner.Scan()
	template := []rune(scanner.Text())

	scanner.Scan() // skip empty line

	rules := make(map[[2]rune]rune)
	for scanner.Scan() {
		var pair [2]rune
		var to rune
		fmt.Sscanf(scanner.Text(), "%c%c -> %c", &pair[0], &pair[1], &to)
		rules[pair] = to
	}

	pairFreq := make(map[[2]rune]int)
	for i := 0; i < len(template)-1; i++ {
		pair := [2]rune{template[i], template[i+1]}
		pairFreq[pair]++
	}

	letterFreq := make(map[rune]int)
	for _, c := range template {
		letterFreq[c] += 1
	}
	return letterFreq, pairFreq, rules
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
