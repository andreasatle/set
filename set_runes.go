// Package set privides primitives for a simple Runeset containing unsigned integers.
package set

// Rename map definition to a Runeset specific name
type runeMap = map[rune]empty

// Definition of the RuneSet type
type RuneSet struct {
	elem runeMap // Map containing the elements in the RuneSet
}

func NewRuneSet(runes ...rune) *RuneSet {
	s := new(RuneSet)
	s.elem = make(runeMap)

	for _, r := range runes {
		s.Add(r)
	}
	return s
}

func (s *RuneSet) Empty() bool {
	return s.Size() == 0
}

func (s *RuneSet) Copy(in RuneSet) {
	s.Clear()
	for key := range in.elem {
		s.Add(key)
	}
}

func FromRuneSlice(slice []rune) *RuneSet {
	s := NewRuneSet()
	for _, e := range slice {
		s.Add(e)
	}
	return s
}

func (s *RuneSet) ToSlice() []rune {
	r := make([]rune, 0, s.Size())

	for key := range s.elem {
		r = append(r, key)
	}
	return r
}

// Create a string with sorted contents of Runeset
func (s *RuneSet) ToString() string {

	values := make([]rune, 0, len(s.elem))

	for key := range s.elem {
		values = append(values, key)
	}

	//sort.Ints(values)
	str := "{"

	for i := 0; i < len(values)-1; i++ {
		str += string(values[i]) + ","
	}

	if len(values) > 0 {
		str += string(values[len(values)-1])
	}

	return str + "}"
}

// Remove all elements from Runeset
func (s *RuneSet) Clear() {
	s.elem = make(runeMap)
}

// Get length of Runeset
func (s *RuneSet) Size() int {
	return len(s.elem)
}

// Add element to Runeset
func (s *RuneSet) Add(e rune) {
	s.elem[e] = empty{}
}

// Remove element from Runeset
func (s *RuneSet) Remove(e rune) {
	delete(s.elem, e)
}

// Check if the Runeset contains the element e
func (s *RuneSet) Contains(e rune) bool {
	_, flag := s.elem[e]
	return flag
}

// Take (non-destructive) union of two Runesets
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

// Take (non-destructive) intersection of two Runesets
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

// Take (non-destructive) difference of two Runesets
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
