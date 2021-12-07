package day7

import (
	"io"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func FuelNeed(currentPos []int, target int, progressive bool) int {
	var fuelNeeded int
	for _, pos := range currentPos {
		steps := abs(target - pos)
		if progressive {
			fuelNeeded += (steps * (steps + 1)) / 2
		} else {
			fuelNeeded += steps
		}
	}
	return fuelNeeded
}

func MinFuel(currentPos []int, low, high int, progressive bool) int {
	minNeed := math.MaxInt
	for i := low; i <= high; i++ {
		need := FuelNeed(currentPos, i, progressive)
		minNeed = min(minNeed, need)
	}
	return minNeed
}

func ParseInput(input io.ReadCloser) ([]int, int, int) {
	b, _ := ioutil.ReadAll(input)
	var row []int
	var low = math.MaxInt
	var high = math.MinInt
	for _, rnum := range strings.Split(string(b), ",") {
		num, err := strconv.Atoi(rnum)
		if err == nil {
			low = min(low, num)
			high = max(high, num)
			row = append(row, num)
		}
	}
	return row, low, high
}

// lol no generics
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
