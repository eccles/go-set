package set

import (
	"slices"
	"testing"
)

// b.Loop is used but there may be some drawbacks:
//
// https://github.com/golang/go/issues/73137

func BenchmarkFromSlice(b *testing.B) {
	for b.Loop() {
		FromSlice("a2", "a3", "a2", "a7")
	}
}

func BenchmarkAddRemove(b *testing.B) {
	s := FromSlice("a2", "a3", "a2", "a7")
	for b.Loop() {
		s.Add("a3", "a4")
		s.Remove("a5", "a6")
	}
}

func BenchmarkRemove(b *testing.B) {
	s := FromSlice("a2", "a3", "a2", "a7")
	for b.Loop() {
		s.Remove("a2")
	}
}

func BenchmarkCreation(b *testing.B) {
	s := FromSlice("a2", "a3", "a2", "a7")
	for b.Loop() {
		s.List()
	}
}

func BenchmarkIntersection(b *testing.B) {
	s1 := FromSlice("a1", "a2", "a2", "a4", "a5")
	s2 := FromSlice("a2", "a3", "a2", "a7")
	for b.Loop() {
		s1.Intersection(s2)
	}
}

func BenchmarkIntersectionIter(b *testing.B) {
	s1 := FromSlice("a1", "a2", "a2", "a4", "a5")
	s2 := slices.Values([]string{"a2", "a3", "a2", "a7"})
	for b.Loop() {
		s1.IntersectionIter(s2)
	}
}

func BenchmarkUnion(b *testing.B) {
	s1 := FromSlice([]string{"a1", "a2", "a2", "a4", "a5"}...)
	s2 := FromSlice("a2", "a3", "a2", "a7")
	for b.Loop() {
		s1.Union(s2)
	}
}

func BenchmarkUnionIter(b *testing.B) {
	s1 := FromSlice("a1", "a2", "a2", "a4", "a5")
	s2 := slices.Values([]string{"a2", "a3", "a2", "a7"})
	for b.Loop() {
		s1.UnionIter(s2)
	}
}

func BenchmarkDifference(b *testing.B) {
	s1 := FromSlice("a1", "a2", "a2", "a4", "a5")
	s2 := FromSlice("a2", "a3", "a2", "a7")
	for b.Loop() {
		s1.Difference(s2)
	}
}

func BenchmarkSymmetricDifferenceString(b *testing.B) {
	s1 := FromSlice("a1", "a2", "a2", "a4", "a5")
	s2 := FromSlice("a2", "a3", "a2", "a7")
	for b.Loop() {
		s1.SymmetricDifference(s2)
	}
}
