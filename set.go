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
// source code.
package set

import (
	"fmt"
	"iter"
	"maps"
	"strings"
)

type (
	// Set is a synonym of a map.
	Set[T comparable] map[T]struct{}
)

// FromIter creates a new set from an iterator. This is useful for creating
// a set from a map.
func FromIter[T comparable](values iter.Seq[T]) Set[T] {
	s := make(map[T]struct{})

	if values != nil {
		for value := range values {
			s[value] = struct{}{}
		}
	}

	return s
}

// FromSlice creates a new set from a slice or array.
func FromSlice[T comparable](values ...T) Set[T] {
	s := make(map[T]struct{})

	for _, value := range values {
		s[value] = struct{}{}
	}

	return s
}

// String returns a string representation of a set.
func (s Set[T]) String() string {
	var b strings.Builder

	b.WriteString("{")

	for k := range s {
		fmt.Fprintf(&b, "%v ", k)
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
	result := make([]T, 0, len(s))

	for value := range s {
		result = append(result, value)
	}

	return result
}

// Add adds one or more values to set.
func (s Set[T]) Add(values ...T) {
	for _, value := range values {
		s[value] = struct{}{}
	}
}

// Remove removes items from a set.
func (s Set[T]) Remove(values ...T) {
	for _, value := range values {
		delete(s, value)
	}
}

// Contains returns true if item is present in the set.
func (s Set[T]) Contains(value T) bool {
	_, c := s[value]

	return c
}

// Union returns set that consists of items that are in either of the 2 sets.
func (s Set[T]) Union(other Set[T]) Set[T] {
	result := make(map[T]struct{})

	for k := range s {
		result[k] = struct{}{}
	}

	for k := range other {
		result[k] = struct{}{}
	}

	return result
}

// UnionIter returns set that consists of items that are in either the set or
// iterable.
func (s Set[T]) UnionIter(values iter.Seq[T]) Set[T] {
	result := make(map[T]struct{})

	for k := range s {
		result[k] = struct{}{}
	}

	for k := range values {
		result[k] = struct{}{}
	}

	return result
}

// Intersection returns set that consists of items that are in both sets.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := make(map[T]struct{})

	for k := range other {
		if s.Contains(k) {
			result[k] = struct{}{}
		}
	}

	return result
}

// IntersectionIter returns set that consists of items that are in both set
// and iterable.
func (s Set[T]) IntersectionIter(values iter.Seq[T]) Set[T] {
	result := make(map[T]struct{})

	for k := range values {
		if s.Contains(k) {
			result[k] = struct{}{}
		}
	}

	return result
}

// Difference returns set that consists of items that are in first set and
// not in second set.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	result := make(map[T]struct{})

	for k := range s {
		if !other.Contains(k) {
			result[k] = struct{}{}
		}
	}

	return result
}

// SymmetricDifference returns set that consists of items that are in each set
// but not in both sets.
func (s Set[T]) SymmetricDifference(other Set[T]) Set[T] {
	result := make(map[T]struct{})

	for k := range s {
		if !other.Contains(k) {
			result[k] = struct{}{}
		}
	}

	for k := range other {
		if !s.Contains(k) {
			result[k] = struct{}{}
		}
	}

	return result
}
