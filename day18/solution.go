package day18

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

type RegularNumber struct {
	n     int
	depth int
}

func (r *RegularNumber) Copy() *RegularNumber {
	return &RegularNumber{
		n:     r.n,
		depth: r.depth,
	}
}

func (r RegularNumber) String() string {
	return fmt.Sprintf("(n=%d, d=%d)", r.n, r.depth)
}

func Explode(rns []*RegularNumber) ([]*RegularNumber, bool) {
	var changed bool
	for i := 0; i < len(rns); i++ {
		if rns[i].depth >= 4 {
			if i-1 >= 0 {
				rns[i-1].n += rns[i].n
			}
			if i < len(rns)-2 {
				rns[i+2].n += rns[i+1].n
			}
			rns[i].n = 0
			rns[i].depth--
			rns[i+1] = nil
			changed = true
			break
		}
	}
	return FilterNils(rns), changed
}

func splitRN(rn *RegularNumber) (*RegularNumber, *RegularNumber) {
	return &RegularNumber{
			n:     rn.n / 2,
			depth: rn.depth + 1,
		},
		&RegularNumber{
			n:     (rn.n / 2) + (rn.n % 2),
			depth: rn.depth + 1,
		}
}

func Split(rns []*RegularNumber) ([]*RegularNumber, bool) {
	var split []*RegularNumber
	var changed bool
	for i := 0; i < len(rns); i++ {
		if rns[i].n >= 10 && !changed {
			a, b := splitRN(rns[i])
			split = append(split, a, b)
			changed = true
		} else {
			split = append(split, rns[i])
		}
	}
	return split, changed
}

func Add(a, b []*RegularNumber) []*RegularNumber {
	var added []*RegularNumber
	for i := 0; i < len(a); i++ {
		a[i].depth += 1
		added = append(added, a[i])
	}
	for i := 0; i < len(b); i++ {
		b[i].depth += 1
		added = append(added, b[i])
	}
	return added
}

func Reduce(rns []*RegularNumber) []*RegularNumber {
	var echanged, schanged bool
	var final = rns[:]
	exploded := func() bool {
		final, echanged = Explode(final)
		return echanged
	}
	splited := func() bool {
		final, schanged = Split(final)
		return schanged
	}
	for exploded() || splited() {
	}
	return final
}

func AddAndReduceAll(input [][]*RegularNumber) []*RegularNumber {
	final := input[0]
	for i := 1; i < len(input); i++ {
		final = Reduce(Add(final, input[i]))
	}
	return final
}

func Magnitude(rns []*RegularNumber) int {
	var magnitude func(rns []*RegularNumber) []*RegularNumber
	magnitude = func(rns []*RegularNumber) []*RegularNumber {
		for i := 0; i < len(rns)-1; i++ {
			if rns[i] == nil {
				continue
			}
			if rns[i].depth == rns[i+1].depth {
				rns[i].n *= 3
				rns[i+1].n *= 2
				rns[i].n += rns[i+1].n
				rns[i].depth--
				rns[i+1] = nil
			}
		}
		return SortByDepth(FilterNils(rns))
	}
	reduced := magnitude(rns)
	for len(reduced) > 1 {
		reduced = magnitude(reduced)
	}
	return reduced[0].n
}

func HighestMagnitude(all [][]*RegularNumber) int {
	highest := 0
	for i := 0; i < len(all); i++ {
		for j := 0; j < len(all); j++ {
			if i == j {
				continue
			}
			a, b := clone(all[i]), clone(all[j])
			reduced := Reduce(Add(a, b))
			m := Magnitude(reduced)
			if m > highest {
				highest = m
			}
		}
	}
	return highest
}

func ParseLine(line string) []*RegularNumber {
	var rns []*RegularNumber
	var depth int
	var consumingNumber bool
	var digits []rune
	var index int
	terminateDigits := func() {
		if consumingNumber {
			rns = append(rns, &RegularNumber{n: toInt(digits), depth: depth})
			digits = []rune{}
			consumingNumber = false
			index++
		}
	}
	depth-- // we like 0-index'd depth
	for _, c := range line {
		switch c {
		case '[':
			depth++
		case ']':
			terminateDigits()
			depth--
		case ',':
			terminateDigits()
		default:
			if c >= '0' && c <= '9' {
				digits = append(digits, c)
				consumingNumber = true
			}
		}
	}
	return rns
}

func ParseInput(r io.ReadCloser) [][]*RegularNumber {
	var all [][]*RegularNumber
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		all = append(all, ParseLine(scanner.Text()))
	}
	return all
}

func toInt(digits []rune) int {
	n, err := strconv.Atoi(string(digits))
	if err != nil {
		panic(err)
	}
	return n
}

func FilterNils(rns []*RegularNumber) []*RegularNumber {
	var filtered []*RegularNumber
	for _, rn := range rns {
		if rn != nil {
			filtered = append(filtered, rn)
		}
	}
	return filtered
}

func clone(rns []*RegularNumber) []*RegularNumber {
	cloned := make([]*RegularNumber, len(rns))
	for i := 0; i < len(rns); i++ {
		cloned[i] = rns[i].Copy()
	}
	return cloned
}

func SortByDepth(rns []*RegularNumber) []*RegularNumber {
	sort.Slice(rns, func(i, j int) bool {
		return rns[i].depth > rns[j].depth
	})
	return rns
}
