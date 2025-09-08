// .*@mycompany\.com MY COMPANY 2025
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package set implements a Set using an underlying map as an aliased type.
//
// Additionally the Set implements methods from the python set type including
// unions, intersections, differences and symmetric differences.
//
// Data can be fed to a set from an array or slice and additionally from an
// iterator. The implementation avoids allocations where possible.
//
// Unlike python, sets can only consist of comparable types. This eliminates the
// possibility of a 'set of sets'. The API is stable.
//
// Similarly to map, sets are not goroutine safe.
//
// This is not production code but simply a demonstration of generics and iterators.
// Do not import.
//
// [Python Set]: https://docs.python.org/3/library/stdtypes.html#set-types-set-frozenset
// Frozenset is NOT supported although it would be possible.
//
// This implementation using generics has a few problems. Map keys or sets are
// 'comparable' (obeys the == and != operations) whereas slices are 'cmp.Ordered'
// (obeys <, <=, ==, >=, > operations) and this introduces some basic incompatibilities.
// Attempting to sort a set using slices.Sort from the go slices package will not work
// as this requires the cmp.Ordered constraint.
//
// A better solution is to implement Set as a stripped down version of Map in the go
// source code perhaps.
//
// For a deep discussion of generics see the blog by Axel Wagner
//
//	https://go.dev/blog/generic-interfaces
//
// For now resist the temptation to convert this package into something more complicated.
package set

import (
	"fmt"
	"iter"
	"maps"
	"slices"
	"strings"
)

type (
	// Set is a synonym of a map with no values.
	Set[T comparable] map[T]struct{}
)

// FromIter creates a new set from an iterator. This is useful for creating
// a set from a map.
func FromIter[T comparable](items iter.Seq[T]) Set[T] {
	s := make(Set[T])

	if items != nil {
		for item := range items {
			s[item] = struct{}{}
		}
	}

	return s
}

// FromSlice creates a new set from a slice or array.
func FromSlice[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))

	for _, item := range items {
		s[item] = struct{}{}
	}

	return s
}

// Equal reports whether two sets contain the same items.
func Equal[S Set[T], T comparable](s1 S, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			return false
		}
	}
	return true
}

// String returns a string representation of a set.
func (s Set[T]) String() string {
	if s == nil {
		return "{}"
	}
	var b strings.Builder

	b.WriteString("{")
	first := true
	for item := range s {
		if first {
			fmt.Fprintf(&b, "%v", item)
			first = false
		} else {
			fmt.Fprintf(&b, " %v", item)
		}
	}

	b.WriteString("}")

	return b.String()
}

// Iter returns an iterator over the set.
func (s Set[T]) Iter() iter.Seq[T] {
	return maps.Keys(s)
}

// List returns the set as the original array.
// Order is not preserved.
func (s Set[T]) List() []T {
	return slices.Collect(s.Iter())
}

// Add adds one or more items to set.
func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

// Remove removes items from a set.
func (s Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s, item)
	}
}

// Contains returns true if item is present in the set.
func (s Set[T]) Contains(item T) bool {
	_, exists := s[item]
	return exists
}

// Equal returns true sets are equal.
func (s Set[T]) Equal(other Set[T]) bool {
	return Equal(s, other)
}

// copy creates a shallow copy of the set.
func (s Set[T]) copy() Set[T] {
	return maps.Clone(s)
}

// Sub returns true if other is a subset of set.
func (s Set[T]) Sub(other Set[T]) bool {
	if len(s) < len(other) {
		return false
	}
	for k := range other {
		if _, ok := s[k]; !ok {
			return false
		}
	}
	return true
}

// Union returns set that consists of items that are in either of the 2 sets.
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := s.copy()

	for item := range other {
		result[item] = struct{}{}
	}

	return result
}

// UnionIter returns set that consists of items that are in either the set or
// iterable.
func (s Set[T]) UnionIter(items iter.Seq[T]) Set[T] {
	result := s.copy()

	for item := range items {
		result[item] = struct{}{}
	}

	return result
}

// Intersection returns set that consists of items that are in both sets.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := make(Set[T])

	for item := range other {
		if s.Contains(item) {
			result[item] = struct{}{}
		}
	}

	return result
}

// IntersectionIter returns set that consists of items that are in both set
// and iterable.
func (s Set[T]) IntersectionIter(items iter.Seq[T]) Set[T] {
	result := make(Set[T])

	for item := range items {
		if s.Contains(item) {
			result[item] = struct{}{}
		}
	}

	return result
}

// Difference returns set that consists of items that are in first set and
// not in second set.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := make(Set[T])

	for item := range s {
		if !other.Contains(item) {
			result[item] = struct{}{}
		}
	}

	return result
}

// SymmetricDifference returns set that consists of items that are in each set
// but not in both sets.
func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	result := make(Set[T])

	for item := range s {
		if !other.Contains(item) {
			result[item] = struct{}{}
		}
	}

	for item := range other {
		if !s.Contains(item) {
			result[item] = struct{}{}
		}
	}

	return result
}
