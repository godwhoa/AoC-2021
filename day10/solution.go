package day10

import (
	"bufio"
	"io"
	"sort"
)

var pairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var scores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completionScores = map[rune]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func SyntaxErrorScore(line string) int {
	stack := new(Stack)
	for _, r := range line {
		switch r {
		case '(', '[', '{', '<':
			stack.Push(r)
		case ')', ']', '}', '>':
			if popped := stack.Pop(); popped == 0 || r != pairs[popped] {
				return scores[r]
			}
		}
	}
	return 0
}

func TotalSyntaxErrorScore(input []string) (score int) {
	for _, line := range input {
		score += SyntaxErrorScore(line)
	}
	return
}

func CompletionScore(line string) (score int) {
	stack := new(Stack)
	for _, r := range line {
		switch r {
		case '(', '[', '{', '<':
			stack.Push(r)
		case ')', ']', '}', '>':
			stack.Pop()
		}
	}
	for stack.Next() {
		score *= 5
		score += completionScores[stack.Pop()]
	}
	return
}

func TotalCompletionScore(input []string) int {
	var scores sort.IntSlice
	for _, line := range input {
		if SyntaxErrorScore(line) == 0 {
			scores = append(scores, CompletionScore(line))
		}

	}
	sort.Sort(scores)
	return scores[len(scores)/2]
}

func ParseInput(input io.ReadCloser) []string {
	var lines []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// func popFn(stack *[]rune) func() rune {
// 	return func() rune {
// 		if len(stack) == 0 {
// 			return 0
// 		}
// 		r := stack[len(stack)-1]
// 		stack = stack[:len(*stack)-1]
// 		return r
// 	}
// }

type Stack []rune

func (s *Stack) Push(e rune) {
	*s = append(*s, e)
}

func (s *Stack) Next() bool {
	return len(*s) > 0
}

func (s *Stack) Pop() rune {
	l := len(*s)
	if l == 0 {
		return 0
	}
	r := (*s)[l-1]
	*s = (*s)[:l-1]
	return r
}
