package sets

import (
	"testing"
)

func ExampleSet_Difference() {
	s1 := New[int](0, 1, 2)
	s2 := New[int](1, 2, 3)
	d := s1.Difference(s2)
	if !d.Has(0) {
		panic("unexpected result")
	}
}

func TestStringSet(t *testing.T) {
	s := Set[string]{}
	s2 := Set[string]{}
	if len(s) != 0 {
		t.Errorf("Expected len=0: %d", len(s))
	}
	s.Insert("a", "b")
	if len(s) != 2 {
		t.Errorf("Expected len=2: %d", len(s))
	}
	s.Insert("c")
	if s.Has("d") {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.Has("a") {
		t.Errorf("Missing contents: %#v", s)
	}
	s.Delete("a")
	if s.Has("a") {
		t.Errorf("Unexpected contents: %#v", s)
	}
	s.Insert("a")
	if s.HasAll("a", "b", "d") {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.HasAll("a", "b") {
		t.Errorf("Missing contents: %#v", s)
	}
	s2.Insert("a", "b", "d")
	if s.IsSuperset(s2) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	s2.Delete("d")
	if !s.IsSuperset(s2) {
		t.Errorf("Missing contents: %#v", s)
	}
}

func TestStringSetDeleteMultiples(t *testing.T) {
	s := Set[string]{}
	s.Insert("a", "b", "c")
	if len(s) != 3 {
		t.Errorf("Expected len=3: %d", len(s))
	}

	s.Delete("a", "c")
	if len(s) != 1 {
		t.Errorf("Expected len=1: %d", len(s))
	}
	if s.Has("a") {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if s.Has("c") {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.Has("b") {
		t.Errorf("Missing contents: %#v", s)
	}

}

func TestNewStringSet(t *testing.T) {
	s := New[string]("a", "b", "c")
	if len(s) != 3 {
		t.Errorf("Expected len=3: %d", len(s))
	}
	if !s.Has("a") || !s.Has("b") || !s.Has("c") {
		t.Errorf("Unexpected contents: %#v", s)
	}
}

func TestStringSetList(t *testing.T) {
	items := []string{"z", "y", "x", "a"}
	s := New[string](items...)
	for _, item := range items {
		if !contains(s.List(), item) {
			t.Errorf("expected list to contain '%s'", item)
		}
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func TestStringSetDifference(t *testing.T) {
	a := New[string]("1", "2", "3")
	b := New[string]("1", "2", "4", "5")
	c := a.Difference(b)
	d := b.Difference(a)
	if len(c) != 1 {
		t.Errorf("Expected len=1: %d", len(c))
	}
	if !c.Has("3") {
		t.Errorf("Unexpected contents: %#v", c.List())
	}
	if len(d) != 2 {
		t.Errorf("Expected len=2: %d", len(d))
	}
	if !d.Has("4") || !d.Has("5") {
		t.Errorf("Unexpected contents: %#v", d.List())
	}
}

func TestStringSetHasAny(t *testing.T) {
	a := New[string]("1", "2", "3")

	if !a.HasAny("1", "4") {
		t.Errorf("expected true, got false")
	}

	if a.HasAny("0", "4") {
		t.Errorf("expected false, got true")
	}
}

func TestStringSetEquals(t *testing.T) {
	// Simple case (order doesn't matter)
	a := New[string]("1", "2")
	b := New[string]("2", "1")
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	// It is a set; duplicates are ignored
	b = New[string]("2", "2", "1")
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	// Edge cases around empty sets / empty strings
	a = New[string]()
	b = New[string]()
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	b = New[string]("1", "2", "3")
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	b = New[string]("1", "2", "")
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	// Check for equality after mutation
	a = New[string]()
	a.Insert("1")
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	a.Insert("2")
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	a.Insert("")
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	a.Delete("")
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}
}

func TestStringUnion(t *testing.T) {
	tests := []struct {
		s1       Set[string]
		s2       Set[string]
		expected Set[string]
	}{
		{
			New[string]("1", "2", "3", "4"),
			New[string]("3", "4", "5", "6"),
			New[string]("1", "2", "3", "4", "5", "6"),
		},
		{
			New[string]("1", "2", "3", "4"),
			New[string](),
			New[string]("1", "2", "3", "4"),
		},
		{
			New[string](),
			New[string]("1", "2", "3", "4"),
			New[string]("1", "2", "3", "4"),
		},
		{
			New[string](),
			New[string](),
			New[string](),
		},
	}

	for _, test := range tests {
		union := test.s1.Union(test.s2)
		if union.Len() != test.expected.Len() {
			t.Errorf("Expected union.Len()=%d but got %d", test.expected.Len(), union.Len())
		}

		if !union.Equal(test.expected) {
			t.Errorf("Expected union.Equal(expected) but not true.  union:%v expected:%v", union.List(), test.expected.List())
		}
	}
}

func TestStringIntersection(t *testing.T) {
	tests := []struct {
		s1       Set[string]
		s2       Set[string]
		expected Set[string]
	}{
		{
			New[string]("1", "2", "3", "4"),
			New[string]("3", "4", "5", "6"),
			New[string]("3", "4"),
		},
		{
			New[string]("1", "2", "3", "4"),
			New[string]("1", "2", "3", "4"),
			New[string]("1", "2", "3", "4"),
		},
		{
			New[string]("1", "2", "3", "4"),
			New[string](),
			New[string](),
		},
		{
			New[string](),
			New[string]("1", "2", "3", "4"),
			New[string](),
		},
		{
			New[string](),
			New[string](),
			New[string](),
		},
	}

	for _, test := range tests {
		intersection := test.s1.Intersection(test.s2)
		if intersection.Len() != test.expected.Len() {
			t.Errorf("Expected intersection.Len()=%d but got %d", test.expected.Len(), intersection.Len())
		}

		if !intersection.Equal(test.expected) {
			t.Errorf("Expected intersection.Equal(expected) but not true.  intersection:%v expected:%v", intersection.List(), test.expected.List())
		}
	}
}

func TestSet_PopAny(t *testing.T) {
	t.Run("full set", func(t *testing.T) {
		items := []string{"1", "2", "3"}
		s := New[string](items...)
		if v, _ := s.PopAny(); !contains(items, v) {
			t.Errorf("expected set to contain '%v'", v)
		}
	})
	t.Run("empty set", func(t *testing.T) {
		s := New[string]()
		_, ok := s.PopAny()
		if ok {
			t.Error("expected ok to be false because the set is empty")
		}
	})
}
