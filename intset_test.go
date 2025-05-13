package set

import (
	"maps"
	"slices"
	"testing"
)

// nolint: gochecknoglobals // these are constants in a test package
var (
	inTinputSet            = FromSlice([]int{1, 2, 2, 4, 5}...)
	inTinputMap            = map[int]int{1: 0, 2: 0, 6: 0}
	inTinputList           = []int{2, 3, 2, 7}
	inToutputList1         = []int{1, 2, 6}
	inToutputList2         = []int{2, 3, 7}
	inTintersection        = []int{2}
	inTunion               = []int{1, 2, 3, 4, 5, 7}
	inTdifference          = []int{1, 4, 5}
	inTsymmetricDifference = []int{1, 3, 4, 5, 7}
)

// TestIntCreationListAddRemove tests simple additon and removal of fields.
func TestIntCreationListAddRemove(t *testing.T) {
	s := FromSlice(inTinputList...)
	assertElementsMatch(t, inToutputList2, s.List())

	// add something that already exists
	assertSetContains(t, s, 3)
	s.Add(3)
	assertSetContains(t, s, 3)

	// add something that doe not exist
	assertSetNotContains(t, s, 4)
	s.Add(4)
	assertSetContains(t, s, 4)

	// delete something that does not exist
	assertSetNotContains(t, s, 5)
	s.Remove(5)
	assertSetNotContains(t, s, 5)

	// delete something that does exist
	assertSetContains(t, s, 2)
	s.Remove(2)
	assertSetNotContains(t, s, 2)
}

// TestIntIter tests the returned iterator.
func TestIntIter(t *testing.T) {
	s := FromIter(maps.Keys(inTinputMap))
	for k := range s.Iter() {
		s.Remove(k)
	}
	assertLen(t, s, 0)
}

// TestIntCreationMap tests creation from map using iterator.
func TestIntCreationMap(t *testing.T) {
	s := FromIter(maps.Keys(inTinputMap))
	assertElementsMatch(t, inToutputList1, s.List())
}

// TestIntIntersection tests intersection of 2 sets.
func TestIntIntersection(t *testing.T) {
	s2 := FromSlice(inTinputList...)
	assertElementsMatch(t, inTintersection, inTinputSet.Intersection(s2).List())
}

// TestIntIntersectionIter tests intersection of set and iterable.
func TestIntIntersectionIter(t *testing.T) {
	assertElementsMatch(
		t,
		inTintersection,
		inTinputSet.IntersectionIter(slices.Values(inTinputList)).List(),
	)
}

// TestIntUnion tests union of 2 sets.
func TestIntUnion(t *testing.T) {
	s2 := FromSlice(inTinputList...)
	assertElementsMatch(t, inTunion, inTinputSet.Union(s2).List())
}

// TestIntUnionIter tests union of a set and a iterable.
func TestIntUnionIter(t *testing.T) {
	assertElementsMatch(
		t,
		inTunion,
		inTinputSet.UnionIter(slices.Values(inTinputList)).List(),
	)
}

// TestIntDifference tests difference of 2 sets.
func TestIntDifference(t *testing.T) {
	s2 := FromSlice(inTinputList...)
	assertElementsMatch(t, inTdifference, inTinputSet.Difference(s2).List())
}

// TestIntSymmetricDifference tests symmetric difference of 2 sets.
func TestIntSymmetricDifference(t *testing.T) {
	s2 := FromSlice(inTinputList...)
	assertElementsMatch(t, inTsymmetricDifference, inTinputSet.SymmetricDifference(s2).List())
}
