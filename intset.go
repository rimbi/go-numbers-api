package main

import "sort"

// Intset definition
type Intset map[int]struct{}

// NewIntSet creates a new integer set
func NewIntSet() Intset {
	return make(Intset)
}

// Add adds a slice of integers to the integer set
func (set Intset) Add(ints []int) {
	for _, i := range ints {
		set[i] = struct{}{}
	}
}

// AsSortedSlice returns the contents of the integer Intset
// as a sorted slice of integers
func (set Intset) AsSortedSlice() []int {
	s := make([]int, len(set))
	i := 0
	for k := range set {
		s[i] = k
		i++
	}
	sort.Ints(s)
	return s
}
