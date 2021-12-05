package day1

import (
	"fmt"
	"io"
	"math"
)

func CountIncreasing(input []int) int {
	count := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] < input[i+1] {
			count++
		}
	}
	return count
}

func SlidingWindowIncreasing(input []int) int {
	count := 0
	previous := math.MaxInt
	for i := 0; i < len(input)-2; i++ {
		current := input[i] + input[i+1] + input[i+2]
		if current > previous {
			count++
		}
		previous = current
	}
	return count
}

func ParseInput(input io.ReadCloser) []int {
	var numbers []int
	for {
		var number int
		_, err := fmt.Fscan(input, &number)
		if err == io.EOF {
			break
		}
		numbers = append(numbers, number)
	}
	return numbers
}
