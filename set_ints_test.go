package set

import "testing"

func TestInts(t *testing.T) {
	s := NewIntSet()

	if len(s.elem) != 0 {
		t.Errorf("Init Test failed, s.elem should contain 0 elems, but contains %d!", len(s.elem))
	}

	if !s.Empty() {
		t.Errorf("Size Test failed, s.elem should contain 0 elems, but contains %d!", s.Size())
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Add(4)
	s.Add(5)
	s.Add(6)
	s.Add(7)
	s.Add(8)
	s.Add(9)

	if len(s.elem) != 9 {
		t.Errorf("Add Test failed, s.elem should contain 9 elems, but contains %d!", len(s.elem))
	}

	if s.Size() != 9 {
		t.Errorf("Size Test failed, s.elem should contain 9 elems, but contains %d!", s.Size())
	}

	s.Remove(4)
	s.Remove(5)
	s.Remove(6)

	if len(s.elem) != 6 {
		t.Errorf("Remove Test failed, s.elem should contain {A,B,C,G,H,I}, but contains %s!", s.ToString())
	}

	s.Add(5)

	if len(s.elem) != 7 {
		t.Errorf("Add Test failed, s.elem should contain 7 elems, but contains %d!", len(s.elem))
	}

	s.Add(5)

	if len(s.elem) != 7 {
		t.Errorf("Add Test failed, s.elem should contain 7 elems, but contains %d!", len(s.elem))
	}

	s.Remove(4711)
	if len(s.elem) != 7 {
		t.Errorf("Remove Test failed, s.elem should contain 7 elems, but contains %d!", len(s.elem))
	}

	if !s.Contains(1) {
		t.Errorf("Contains Test failed, s.elem should contain element A")
	}

	if s.Contains(6) {
		t.Errorf("Contains Test failed, s.elem should not contain element F")
	}

	s.Clear()
	if len(s.elem) != 0 {
		t.Errorf("Clear Test failed, s.elem should contain 0 elems, but contains %d!", len(s.elem))
	}

	s1 := NewIntSet(1, 2, 3)
	s2 := NewIntSet()

	s2.Add(1)
	s2.Add(2)
	s2.Add(4)
	s = s1.Union(s2)
	if s.Size() != 4 || !s.Contains(1) || !s.Contains(2) || !s.Contains(3) || !s.Contains(4) {
		t.Errorf("Union Test failed, %s U %s == %s!", s1.ToString(), s2.ToString(), s.ToString())
	}

	s = s1.Intersection(s2)
	if len(s1.elem) != 3 || len(s2.elem) != 3 || len(s.elem) != 2 || !s.Contains(1) || !s.Contains(2) {
		t.Errorf("Union Test failed, s.elem should contain 2 elems, but contains %d!", len(s.elem))
	}

	s1.Add(5)
	s = s1.Difference(s2)
	if len(s1.elem) != 4 || len(s2.elem) != 3 || len(s.elem) != 2 || !s.Contains(3) || !s.Contains(5) {
		t.Errorf("Union Test failed, s.elem should contain 2 elems, but contains %d!", len(s.elem))
	}

	str := s.ToString()
	if str != "{3,5}" && str != "{5,3}" {
		t.Errorf("ToString Test failed!, %s != %s", str, "{3,5}")
	}

	myInts := s.ToSlice()
	if len(myInts) != 2 {
		t.Errorf("ToString Test failed!, %s != %s", str, "{3,5}")
	}

	if s1.Equal(s2) {
		t.Errorf("Equal Test 1 failed!, %s == %s", s1.ToString(), s2.ToString())
	}

	s2 = NewIntSet(1, 2, 3, 5)
	if !s1.Equal(s2) {
		t.Errorf("Equal Test 2 failed!, %s == %s", s1.ToString(), s2.ToString())
	}
}
