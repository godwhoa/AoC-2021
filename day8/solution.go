package day8

import (
	"bufio"
	"io"
	"strings"
)

func DeduceFive(bag *Bag, patterns []*Pattern) {
	for _, pattern := range patterns {
		if bag.Pattern(1).IsSubsetOf(pattern) {
			bag.Add(3, pattern)
		} else if pattern.IsSubsetOf(bag.Pattern(6)) {
			bag.Add(5, pattern)
		} else {
			bag.Add(2, pattern)
		}
	}
}

func DeduceSix(bag *Bag, patterns []*Pattern) {
	for _, pattern := range patterns {
		if !bag.Pattern(1).IsSubsetOf(pattern) {
			bag.Add(6, pattern)
		} else if bag.Pattern(4).IsSubsetOf(pattern) {
			bag.Add(9, pattern)
		} else {
			bag.Add(0, pattern)
		}
	}
}

func DeducePatterns(rawPatterns []string) *Bag {
	mapping := NewBag()
	sixes := make([]*Pattern, 0, 3)
	fives := make([]*Pattern, 0, 3)
	for _, raw := range rawPatterns {
		pattern := NewPattern(raw)
		switch len(raw) {
		case 2:
			mapping.Add(1, pattern)
		case 4:
			mapping.Add(4, pattern)
		case 7:
			mapping.Add(8, pattern)
		case 3:
			mapping.Add(7, pattern)
		case 6:
			sixes = append(sixes, pattern)
		case 5:
			fives = append(fives, pattern)
		}
	}

	DeduceSix(mapping, sixes)
	DeduceFive(mapping, fives)
	return mapping
}

func SumOfDisplayOutputs(input io.ReadCloser) (sum int) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")
		displayDigits := strings.Split(parts[1], " ")
		mappings := DeducePatterns(strings.Split(parts[0], " "))

		num := 0
		for _, rawDigit := range displayDigits {
			digit := mappings.Query(rawDigit)
			num = num*10 + digit
		}
		sum += num
	}
	return
}

func CountSpecialDigits(input io.ReadCloser) (count int) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		display := strings.Split(scanner.Text(), "|")[1]
		digits := strings.Split(display, " ")
		for _, digit := range digits {
			segments := len(digit)
			switch segments {
			case 2, 4, 3, 7:
				count++
			}
		}
	}
	return
}
