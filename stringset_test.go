package set

import (
	"maps"
	"slices"
	"testing"
)

// nolint: gochecknoglobals // these are constants in a test package
var (
	inputSet = FromSlice([]string{"a1", "a2", "a2", "a4", "a5"}...)

	inputList  = []string{"a2", "a3", "a2", "a7"}
	outputList = []string{"a2", "a3", "a7"}

	inputMap    = map[string]string{"a1": "", "a2": "", "a6": ""}
	outputList1 = []string{"a1", "a2", "a6"}

	intersection        = []string{"a2"}
	union               = []string{"a1", "a2", "a3", "a4", "a5", "a7"}
	difference          = []string{"a1", "a4", "a5"}
	symmetricDifference = []string{"a1", "a3", "a4", "a5", "a7"}
)

func assertLen[K comparable](t *testing.T, s Set[K], length int) bool {
	if len(s) != length {
		t.Errorf("Expected length %d but got %d", length, len(s))
	}
	return true
}

func assertElementsMatch[K comparable](t *testing.T, actual, expected []K) bool {
	for _, k := range actual {
		if !slices.Contains(expected, k) {
			t.Errorf("Expected %v but got %v", expected, actual)
		}
	}
	if len(actual) != len(expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
	return true
}

func assertSetContains[K comparable](t *testing.T, s Set[K], key K) bool {
	if !s.Contains(key) {
		t.Errorf("%v does not contain %v", s, key)
	}
	return true
}

func assertSetNotContains[K comparable](t *testing.T, s Set[K], key K) bool {
	if s.Contains(key) {
		t.Errorf("%v contains %v", s, key)
	}
	return true
}

// TestCreationListAddRemove tests simple additon and removal of fields.
func TestCreationListAddRemove(t *testing.T) {
	s := FromSlice(inputList...)
	assertElementsMatch(t, outputList, s.List())

	// add something that already exists
	assertSetContains(t, s, "a2")
	s.Add("a3")
	assertSetContains(t, s, "a2")

	// add something that doe not exist
	assertSetNotContains(t, s, "a4")
	s.Add("a4")
	assertSetContains(t, s, "a4")

	// delete something that does not exist
	assertSetNotContains(t, s, "a5")
	s.Remove("a5")
	assertSetNotContains(t, s, "a5")

	// delete something that does exist
	assertSetContains(t, s, "a2")
	s.Remove("a2")
	assertSetNotContains(t, s, "a2")
}

// TestIter tests the returned iterator.
func TestIter(t *testing.T) {
	s := FromIter(maps.Keys(inputMap))
	for k := range s.Iter() {
		s.Remove(k)
	}
	assertLen(t, s, 0)
}

// TestCreationMap tests creation from map using iterator.
func TestCreationMap(t *testing.T) {
	s := FromIter(maps.Keys(inputMap))
	assertElementsMatch(t, outputList1, s.List())
}

// TestIntersection tests intersection of 2 sets.
func TestIntersection(t *testing.T) {
	s2 := FromSlice(inputList...)
	assertElementsMatch(t, intersection, inputSet.Intersection(s2).List())
}

// TestIntersectionIter tests intersection of set and iterable.
func TestIntersectionIter(t *testing.T) {
	assertElementsMatch(
		t,
		intersection,
		inputSet.IntersectionIter(slices.Values(inputList)).List(),
	)
}

// TestUnion tests union of 2 sets.
func TestUnion(t *testing.T) {
	s2 := FromSlice(inputList...)
	assertElementsMatch(t, union, inputSet.Union(s2).List())
}

// TestUnionIter tests union of a set and a iterable.
func TestUnionIter(t *testing.T) {
	assertElementsMatch(
		t,
		union,
		inputSet.UnionIter(slices.Values(inputList)).List(),
	)
}

// TestDifference tests difference of 2 sets.
func TestDifference(t *testing.T) {
	s2 := FromSlice(inputList...)
	assertElementsMatch(t, difference, inputSet.Difference(s2).List())
}

// TestSymmetricDifference tests symmetric difference of 2 sets.
func TestSymmetricDifference(t *testing.T) {
	s2 := FromSlice(inputList...)
	assertElementsMatch(t, symmetricDifference, inputSet.SymmetricDifference(s2).List())
}
