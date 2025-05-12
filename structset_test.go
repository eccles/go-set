package set

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// this struct is comparable because all its fields are comparable.
type testStruct struct {
	index int
	name  string
}

// nolint: gochecknoglobals // these are constants in a test package
var (
	oneStruct                 = testStruct{1, "one"}
	twoStruct                 = testStruct{2, "two"}
	threeStruct               = testStruct{3, "three"}
	fourStruct                = testStruct{4, "four"}
	fiveStruct                = testStruct{5, "five"}
	sixStruct                 = testStruct{6, "six"}
	sevenStruct               = testStruct{7, "seven"}
	structinputSet            = FromSlice([]testStruct{oneStruct, twoStruct, twoStruct, fourStruct, fiveStruct}...)
	structinputMap            = map[testStruct]int{oneStruct: 1, twoStruct: 2, sixStruct: 6}
	structinputList           = []testStruct{twoStruct, threeStruct, twoStruct, sevenStruct}
	structoutputList1         = []testStruct{oneStruct, twoStruct, sixStruct}
	structoutputList2         = []testStruct{twoStruct, threeStruct, sevenStruct}
	structintersection        = []testStruct{twoStruct}
	structunion               = []testStruct{oneStruct, twoStruct, threeStruct, fourStruct, fiveStruct, sevenStruct}
	structdifference          = []testStruct{oneStruct, fourStruct, fiveStruct}
	structsymmetricDifference = []testStruct{oneStruct, threeStruct, fourStruct, fiveStruct, sevenStruct}
)

// TestStructCreationListAddRemove tests simple additon and removal of fields.
func TestStructCreationListAddRemove(t *testing.T) {
	s := FromSlice(structinputList...)
	assert.ElementsMatch(t, structoutputList2, s.List())

	// add something that already exists
	assert.Equal(t, true, s.Contains(threeStruct))
	s.Add(threeStruct)
	assert.Equal(t, true, s.Contains(threeStruct))

	// add something that doe not exist
	assert.Equal(t, false, s.Contains(fourStruct))
	s.Add(fourStruct)
	assert.Equal(t, true, s.Contains(fourStruct))

	// delete something that does not exist
	assert.Equal(t, false, s.Contains(fiveStruct))
	s.Remove(fiveStruct)
	assert.Equal(t, false, s.Contains(fiveStruct))

	// delete something that does exist
	assert.Equal(t, true, s.Contains(twoStruct))
	s.Remove(twoStruct)
	assert.Equal(t, false, s.Contains(twoStruct))
}

// TestStructIter tests the returned iterator.
func TestStructIter(t *testing.T) {
	s := FromIter(maps.Keys(structinputMap))
	for k := range s.Iter() {
		s.Remove(k)
	}
	assert.Len(t, s, 0)
}

// TestStructCreationMap tests creation from map using iterator.
func TestStructCreationMap(t *testing.T) {
	s := FromIter(maps.Keys(structinputMap))
	assert.ElementsMatch(t, structoutputList1, s.List())
}

// TestStructIntersection tests intersection of 2 sets.
func TestStructIntersection(t *testing.T) {
	s2 := FromSlice(structinputList...)
	assert.ElementsMatch(t, structintersection, structinputSet.Intersection(s2).List())
}

// TestStructIntersectionIter tests intersection of set and iterable.
func TestStructIntersectionIter(t *testing.T) {
	assert.ElementsMatch(
		t,
		structintersection,
		structinputSet.IntersectionIter(slices.Values(structinputList)).List(),
	)
}

// TestStructUnion tests union of 2 sets.
func TestStructUnion(t *testing.T) {
	s2 := FromSlice(structinputList...)
	assert.ElementsMatch(t, structunion, structinputSet.Union(s2).List())
}

// TestStructUnionIter tests union of a set and a iterable.
func TestStructUnionIter(t *testing.T) {
	assert.ElementsMatch(
		t,
		structunion,
		structinputSet.UnionIter(slices.Values(structinputList)).List(),
	)
}

// TestStructDifference tests difference of 2 sets.
func TestStructDifference(t *testing.T) {
	s2 := FromSlice(structinputList...)
	assert.ElementsMatch(t, structdifference, structinputSet.Difference(s2).List())
}

// TestStructSymmetricDifference tests symmetric difference of 2 sets.
func TestStructSymmetricDifference(t *testing.T) {
	s2 := FromSlice(structinputList...)
	assert.ElementsMatch(t, structsymmetricDifference, structinputSet.SymmetricDifference(s2).List())
}
