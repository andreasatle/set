// Package set privides primitives for a simple set containing unsigned integers.
package set

import (
	"strconv"
)

// Rename map definition to a set specific name
type intSet = map[int]empty


// Definition of the Set type
type IntSet struct {
	elem intSet // Map containing the elements in the set
}

func NewIntSet(ints ...int) *IntSet {
	s := new(IntSet)
	s.elem = make(intSet)

	for _, r := range ints {
		s.Add(r)
	}
	return s
}

func (s *IntSet) Empty() bool {
	return s.Size() == 0
}

func (s *IntSet) Copy(in IntSet) {
	s.Clear()
	for key := range in.elem {
		s.Add(key)
	}
}

func FromIntSlice(slice []int) *IntSet {
	s := NewIntSet()
	for _, e := range slice {
		s.Add(e)
	}
	return s
}

func (s *IntSet) ToSlice() []int {
	r := make([]int, 0, s.Size())

	for key := range s.elem {
		r = append(r, key)
	}
	return r
}

// Create a string with sorted contents of set
func (s *IntSet) ToString() string {

	values := make([]int, 0, len(s.elem))

	for key := range s.elem {
		values = append(values, key)
	}

	//sort.Ints(values)
	str := "{"

	for i := 0; i < len(values)-1; i++ {
		str += strconv.Itoa(values[i]) + ","
	}

	if len(values) > 0 {
		str += strconv.Itoa(values[len(values)-1])
	}

	return str + "}"
}

// Remove all elements from set
func (s *IntSet) Clear() {
	s.elem = make(intSet)
}

// Get length of set
func (s *IntSet) Size() int {
	return len(s.elem)
}

// Add element to set
func (s *IntSet) Add(e int) {
	s.elem[e] = empty{}
}

// Remove element from set
func (s *IntSet) Remove(e int) {
	delete(s.elem, e)
}

// Check if the set contains the element e
func (s *IntSet) Contains(e int) bool {
	_, flag := s.elem[e]
	return flag
}

// Take (non-destructive) union of two sets
func (s *IntSet) Union(s2 *IntSet) *IntSet {
	out := new(IntSet)
	out.Clear()

	for key := range s.elem {
		out.Add(key)
	}

	for key := range s2.elem {
		out.Add(key)
	}

	return out
}

// Take (non-destructive) intersection of two sets
func (s *IntSet) Intersection(s2 *IntSet) *IntSet {
	out := new(IntSet)
	out.Clear()

	for key := range s.elem {
		if s2.Contains(key) {
			out.Add(key)
		}
	}

	return out
}

// Take (non-destructive) difference of two sets
func (s *IntSet) Difference(s2 *IntSet) *IntSet {
	out := new(IntSet)
	out.Clear()

	for key := range s.elem {
		out.Add(key)
	}

	for key := range s2.elem {
		out.Remove(key)
	}

	return out
}

func (s *IntSet) Equal(s2 *IntSet) bool {
	return s.Difference(s2).Size() == 0 && s2.Difference(s).Size() == 0
}
