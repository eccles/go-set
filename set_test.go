package set_test

import (
	"fmt"
	"maps"
	"slices"

	"github.com/eccles/go-set"
)

func ExampleFromIter() {
	m := map[string]int{"a": 0, "b": 1}
	s := set.FromIter(maps.Keys(m))
	fmt.Printf(
		"%d %t %t",
		len(s),
		s.Contains("a"),
		s.Contains("b"),
	)
	// Output: 2 true true
}

func ExampleFromSlice() {
	a := []string{"a", "b"}
	s := set.FromSlice(a...)
	fmt.Printf(
		"%d %t %t",
		len(s),
		s.Contains("a"),
		s.Contains("b"),
	)
	// Output: 2 true true
}

func ExampleFromSlice_int() {
	a := []int{1, 2}
	s := set.FromSlice(a...)
	fmt.Printf(
		"%d %t %t",
		len(s),
		s.Contains(1),
		s.Contains(2),
	)
	// Output: 2 true true
}

func ExampleFromSlice_struct() {
	type testStruct struct {
		name  string
		index int
	}

	oneTestStruct := testStruct{"one", 1}
	twoTestStruct := testStruct{"two", 2}

	a := []testStruct{oneTestStruct, twoTestStruct}
	s := set.FromSlice(a...)
	fmt.Printf(
		"%d %t %t",
		len(s),
		s.Contains(oneTestStruct),
		s.Contains(twoTestStruct),
	)
	// Output: 2 true true
}

func ExampleSet_List() {
	s := set.FromSlice("a")
	fmt.Printf("%v", s.List())
	// Output: [a]
}

func ExampleSet_Add() {
	s := set.FromSlice("a")
	s.Add("b")
	fmt.Printf(
		"%d %t %t",
		len(s),
		s.Contains("a"),
		s.Contains("b"),
	)
	// Output: 2 true true
}

func ExampleSet_Remove() {
	s := set.FromSlice("a", "b")
	s.Remove("b")
	fmt.Printf(
		"%d %t",
		len(s),
		s.Contains("a"),
	)
	// Output: 1 true
}

func ExampleSet_Iter() {
	s := set.FromSlice("a", "b", "c")

	t := []string{}
	for k := range s.Iter() {
		t = append(t, k)
	}

	slices.Sort(t)
	fmt.Printf(
		"%d %v",
		len(t),
		t,
	)
	// Output: 3 [a b c]
}

func ExampleSet_Union() {
	a := []string{"a", "b", "c"}
	m := []string{"c", "d", "e"}
	s := set.FromSlice(a...)
	t := set.FromSlice(m...)
	u := s.Union(t)
	fmt.Printf(
		"%d %t %t %t %t %t",
		len(u),
		u.Contains("a"),
		u.Contains("b"),
		u.Contains("c"),
		u.Contains("d"),
		u.Contains("e"),
	)
	// Output: 5 true true true true true
}

func ExampleSet_Union_int() {
	a := []int{1, 2, 3}
	m := []int{3, 4, 5}
	s := set.FromSlice(a...)
	t := set.FromSlice(m...)
	u := s.Union(t)
	fmt.Printf(
		"%d %t %t %t %t %t",
		len(u),
		u.Contains(1),
		u.Contains(2),
		u.Contains(3),
		u.Contains(4),
		u.Contains(5),
	)
	// Output: 5 true true true true true
}

func ExampleSet_Union_struct() {
	type testStruct struct {
		name  string
		index int
	}

	oneTestStruct := testStruct{"one", 1}
	twoTestStruct := testStruct{"two", 2}
	threeTestStruct := testStruct{"three", 3}
	fourTestStruct := testStruct{"four", 4}
	fiveTestStruct := testStruct{"five", 5}

	a := []testStruct{oneTestStruct, twoTestStruct, threeTestStruct}
	m := []testStruct{threeTestStruct, fourTestStruct, fiveTestStruct}
	s := set.FromSlice(a...)
	t := set.FromSlice(m...)
	u := s.Union(t)
	fmt.Printf(
		"%d %t %t %t %t %t",
		len(u),
		u.Contains(oneTestStruct),
		u.Contains(twoTestStruct),
		u.Contains(threeTestStruct),
		u.Contains(fourTestStruct),
		u.Contains(fiveTestStruct),
	)
	// Output: 5 true true true true true
}

func ExampleSet_UnionIter() {
	a := []string{"a", "b", "c"}
	m := map[string]int{"c": 0, "d": 1, "e": 2}
	s := set.FromSlice(a...)
	u := s.UnionIter(maps.Keys(m))
	fmt.Printf(
		"%d %t %t %t %t %t",
		len(u),
		u.Contains("a"),
		u.Contains("b"),
		u.Contains("c"),
		u.Contains("d"),
		u.Contains("e"),
	)
	// Output: 5 true true true true true
}

func ExampleSet_Intersection() {
	a := []string{"a", "b", "c"}
	m := []string{"c", "d", "e"}
	s := set.FromSlice(a...)
	t := set.FromSlice(m...)
	u := s.Intersection(t)
	fmt.Printf("%v", u)
	// Output: {c }
}

func ExampleSet_IntersectionIter() {
	a := []string{"a", "b", "c"}
	m := map[string]int{"c": 0, "d": 1, "e": 2}
	s := set.FromSlice(a...)
	u := s.IntersectionIter(maps.Keys(m))
	fmt.Printf("%v", u)
	// Output: {c }
}

func ExampleSet_Difference() {
	a := []string{"a", "b", "c"}
	m := []string{"c", "d", "e"}
	s := set.FromSlice(a...)
	t := set.FromSlice(m...)
	u := s.Difference(t)
	fmt.Printf(
		"%d %t %t",
		len(u),
		u.Contains("a"),
		u.Contains("b"),
	)
	// Output: 2 true true
}

func ExampleSet_SymmetricDifference() {
	a := []string{"a", "b", "c"}
	m := []string{"c", "d", "e"}
	s := set.FromSlice(a...)
	t := set.FromSlice(m...)
	u := s.SymmetricDifference(t)
	fmt.Printf(
		"%d %t %t %t %t",
		len(u),
		u.Contains("a"),
		u.Contains("b"),
		u.Contains("d"),
		u.Contains("e"),
	)
	// Output: 4 true true true true
}
