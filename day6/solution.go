package day6

import (
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func Simulate(fishes map[int]int, days int) int {
	for day := 1; day <= days; day++ {
		draft := cloneMap(fishes)
		for age, count := range fishes {
			draft[age] -= count
			if age-1 == -1 {
				draft[6] += count
				draft[8] += count
			} else {
				draft[age-1] += count
			}
		}
		fishes = draft
	}
	total := 0
	for _, count := range fishes {
		total += count
	}
	return total
}

func cloneMap(m map[int]int) map[int]int {
	clone := make(map[int]int)
	for k, v := range m {
		clone[k] = v
	}
	return clone
}

func ParseInput(input io.ReadCloser) map[int]int {
	fishes := make(map[int]int)
	b, _ := ioutil.ReadAll(input)
	for _, c := range strings.Split(string(b), ",") {
		age, err := strconv.Atoi(c)
		if err == nil {
			fishes[age]++
		}
	}
	return fishes
}
