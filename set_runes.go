// Package set provides primitives for a simple Runeset containing unsigned integers.
package set

// Rename map definition to a Runeset specific name
type runeSet = map[rune]empty

// Definition of the RuneSet type
type RuneSet struct {
	elem runeSet // Map containing the elements in the RuneSet
}

// Create a new instance of RuneSet
func NewRuneSet(runes ...rune) *RuneSet {
	s := new(RuneSet)
	s.elem = make(runeSet)

	for _, r := range runes {
		s.Add(r)
	}
	return s
}

// Check if the set is empty
func (s *RuneSet) Empty() bool {
	return s.Size() == 0
}

// Create a new copy of RuneSet
func (s *RuneSet) Copy(in RuneSet) {
	s.Clear()
	for key := range in.elem {
		s.Add(key)
	}
}

// Create a RuneSet from a slice of runes
func FromRuneSlice(slice []rune) *RuneSet {
	s := NewRuneSet()
	for _, e := range slice {
		s.Add(e)
	}
	return s
}

// Convert the Set to a slice of runes
func (s *RuneSet) ToSlice() []rune {
	r := make([]rune, 0, s.Size())

	for key := range s.elem {
		r = append(r, key)
	}
	return r
}

// Create a string from the Runeset
func (s *RuneSet) ToString() string {

	values := make([]rune, 0, len(s.elem))

	for key := range s.elem {
		values = append(values, key)
	}

	str := "{"

	for i := 0; i < len(values)-1; i++ {
		str += string(values[i]) + ","
	}

	if len(values) > 0 {
		str += string(values[len(values)-1])
	}

	return str + "}"
}

// Remove all elements from RuneSet
func (s *RuneSet) Clear() {
	s.elem = make(runeSet)
}

// Get length of RuneSet
func (s *RuneSet) Size() int {
	return len(s.elem)
}

// Add element to Runeset
func (s *RuneSet) Add(e rune) {
	s.elem[e] = empty{}
}

// Remove element from RuneSet
func (s *RuneSet) Remove(e rune) {
	delete(s.elem, e)
}

// Check if the Runeset contains the element e
func (s *RuneSet) Contains(e rune) bool {
	_, flag := s.elem[e]
	return flag
}

// Take (non-destructive) union of two RuneSets
func (s *RuneSet) Union(s2 *RuneSet) *RuneSet {
	out := new(RuneSet)
	out.Clear()

	for key := range s.elem {
		out.Add(key)
	}

	for key := range s2.elem {
		out.Add(key)
	}

	return out
}

// Take (non-destructive) intersection of two RuneSets
func (s *RuneSet) Intersection(s2 *RuneSet) *RuneSet {
	out := new(RuneSet)
	out.Clear()

	for key := range s.elem {
		if s2.Contains(key) {
			out.Add(key)
		}
	}

	return out
}

// Take (non-destructive) difference of two RuneSets
func (s *RuneSet) Difference(s2 *RuneSet) *RuneSet {
	out := new(RuneSet)
	out.Clear()

	for key := range s.elem {
		out.Add(key)
	}

	for key := range s2.elem {
		out.Remove(key)
	}

	return out
}

func (s *RuneSet) Equal(s2 *RuneSet) bool {
	return s.Difference(s2).Size() == 0 && s2.Difference(s).Size() == 0
}
