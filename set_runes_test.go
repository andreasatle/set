package set

import "testing"

func TestRunes(t *testing.T) {
	s := NewRuneSet()

	if len(s.elem) != 0 {
		t.Errorf("Init Test failed, s.elem should contain 0 elems, but contains %d!", len(s.elem))
	}

	if !s.Empty() {
		t.Errorf("Empty Test failed, s.elem should contain 0 elems, but contains %d!", s.Size())
	}

	s.Add('A')
	s.Add('B')
	s.Add('C')
	s.Add('D')
	s.Add('E')
	s.Add('F')
	s.Add('G')
	s.Add('H')
	s.Add('I')

	if len(s.elem) != 9 {
		t.Errorf("Add Test failed, s.elem should contain 9 elems, but contains %d!", len(s.elem))
	}

	if s.Size() != 9 {
		t.Errorf("Size Test failed, s.elem should contain 9 elems, but contains %d!", s.Size())
	}

	s.Remove('D')
	s.Remove('E')
	s.Remove('F')

	if len(s.elem) != 6 {
		t.Errorf("Remove Test failed, s.elem should contain {A,B,C,G,H,I}, but contains %s!", s.ToString())
	}

	s.Add('E')

	if len(s.elem) != 7 {
		t.Errorf("Add Test failed, s.elem should contain 7 elems, but contains %d!", len(s.elem))
	}

	s.Add('E')

	if len(s.elem) != 7 {
		t.Errorf("Add Test failed, s.elem should contain 7 elems, but contains %d!", len(s.elem))
	}

	s.Remove('Q')
	if len(s.elem) != 7 {
		t.Errorf("Remove Test failed, s.elem should contain 7 elems, but contains %d!", len(s.elem))
	}

	if !s.Contains('A') {
		t.Errorf("Contains Test failed, s.elem should contain element A")
	}

	if s.Contains('F') {
		t.Errorf("Contains Test failed, s.elem should not contain element F")
	}

	s.Clear()
	if len(s.elem) != 0 {
		t.Errorf("Clear Test failed, s.elem should contain 0 elems, but contains %d!", len(s.elem))
	}

	s1 := NewRuneSet('A', 'B', 'C')
	s2 := NewRuneSet()

	s2.Add('A')
	s2.Add('B')
	s2.Add('D')
	s = s1.Union(s2)
	if s.Size() != 4 || !s.Contains('A') || !s.Contains('B') || !s.Contains('C') || !s.Contains('D') {
		t.Errorf("Union Test failed, %s U %s == %s!", s1.ToString(), s2.ToString(), s.ToString())
	}

	s = s1.Intersection(s2)
	if len(s1.elem) != 3 || len(s2.elem) != 3 || len(s.elem) != 2 || !s.Contains('A') || !s.Contains('B') {
		t.Errorf("Union Test failed, s.elem should contain 2 elems, but contains %d!", len(s.elem))
	}

	s1.Add('E')
	s = s1.Difference(s2)
	if len(s1.elem) != 4 || len(s2.elem) != 3 || len(s.elem) != 2 || !s.Contains('C') || !s.Contains('E') {
		t.Errorf("Union Test failed, s.elem should contain 2 elems, but contains %d!", len(s.elem))
	}

	str := s.ToString()
	if str != "{C,E}" && str != "{E,C}" {
		t.Errorf("ToString Test failed!, %s != %s", str, "{C,E}")
	}

	rune := s.ToSlice()
	if len(rune) != 2 {
		t.Errorf("ToString Test failed!, %s != %s", str, "{C,E}")
	}

	if s1.Equal(s2) {
		t.Errorf("Equal Test 1 failed!, %s == %s", s1.ToString(), s2.ToString())
	}

	s2 = NewRuneSet('A', 'B', 'C', 'E')
	if !s1.Equal(s2) {
		t.Errorf("Equal Test 2 failed!, %s == %s", s1.ToString(), s2.ToString())
	}
}
