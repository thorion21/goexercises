package myset

import (
	"fmt"
	"strings"
)

// Set struct has a map of void* elements to store keys (and redundant values)
type Set struct {
	items map[string]uint
}

// New return a set with internal allocation
func New() Set {
	set := Set{}
	set.items = make(map[string]uint)

	return set
}

// String returns a view of the values formatted correctly
func (s Set) String() string {
	output := "{"
	inter := make([]string, len(s.items))
	counter := 0

	for e := range s.items {
		inter[counter] = "\"" + fmt.Sprintf("%v", e) + "\""
		counter++
	}

	output += strings.Join(inter[:], ", ")
	output += "}"

	return output
}

// NewFromSlice returns a set with values from a slice of strings
func NewFromSlice(slice []string) Set {
	set := Set{}
	set.items = make(map[string]uint)

	if slice == nil {
		return set
	}

	for _, val := range slice {
		set.Add(val)
	}

	return set
}

// IsEmpty returns true if the set does not contain any items
func (s Set) IsEmpty() bool {
	return len(s.items) == 0
}

// Has return true if the value is present in the set
func (s Set) Has(e string) bool {
	_, ok := s.items[e]
	return ok
}

// Equal returns true if the elements in the first set are in the second set and vice-versa
func Equal(s1, s2 Set) bool {
	if len(s1.items) != len(s2.items) {
		return false
	}

	for key := range s1.items {
		if !s2.Has(key) {
			return false
		}
	}

	return true
}

// Add inserts a new element in the set
func (s Set) Add(e string) {
	s.items[e]++
}

// Intersection returns a new set with the elements present in both s1 and s2
func Intersection(s1, s2 Set) Set {
	set := Set{}
	set.items = make(map[string]uint)

	if s1.items == nil || s2.items == nil {
		return set
	}

	for key := range s1.items {
		if s2.Has(key) {
			set.Add(key)
		}
	}

	return set
}

// Union merges together two sets
func Union(s1, s2 Set) Set {
	set := Set{}
	set.items = make(map[string]uint)

	for key := range s1.items {
		set.Add(key)
	}

	for key := range s2.items {
		set.Add(key)
	}

	return set
}

// Difference returns a new set with the values present in s1, but not in s2
func Difference(s1, s2 Set) Set {
	set := Set{}
	set.items = make(map[string]uint)

	for key := range s1.items {
		if !s2.Has(key) {
			set.Add(key)
		}
	}

	return set
}

// Disjoint returns true if the sets have no values in common
func Disjoint(s1, s2 Set) bool {
	for key := range s1.items {
		if s2.Has(key) {
			return false
		}
	}

	return true
}

// Subset returns true if all the elements in s1 are found in s2
func Subset(s1, s2 Set) bool {
	for key := range s1.items {
		if !s2.Has(key) {
			return false
		}
	}

	return true
}
