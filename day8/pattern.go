package day8

import "sort"

func sortString(s string) string {
	pattern := []rune(s)
	sort.Slice(pattern, func(i, j int) bool {
		return pattern[i] < pattern[j]
	})
	return string(pattern)
}

type Pattern struct {
	rs map[rune]struct{}
	s  string
}

func (p *Pattern) IsSubsetOf(other *Pattern) bool {
	for r, _ := range p.rs {
		if _, ok := other.rs[r]; !ok {
			return false
		}
	}
	return true
}

func NewPattern(s string) *Pattern {
	rs := make(map[rune]struct{})
	for _, r := range s {
		rs[r] = struct{}{}
	}
	return &Pattern{
		rs: rs,
		s:  sortString(s),
	}
}

func NewBag() *Bag {
	return &Bag{
		ptod: make(map[string]int),
		dtop: make(map[int]*Pattern),
	}
}

type Bag struct {
	ptod map[string]int
	dtop map[int]*Pattern
}

func (b *Bag) Digit(p *Pattern) int {
	return b.ptod[p.s]
}

func (b *Bag) Pattern(d int) *Pattern {
	return b.dtop[d]
}

func (b *Bag) Query(s string) int {
	return b.ptod[(sortString(s))]
}

func (b *Bag) Add(d int, p *Pattern) {
	b.dtop[d] = p
	b.ptod[p.s] = d
}
