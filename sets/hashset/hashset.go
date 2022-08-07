// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package hashset

import (
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/sets"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Set implementation
var _ sets.Set[string] = (*Set[string])(nil)

// Set holds elements in go's native map
type Set[T comparable] struct {
	items map[T]struct{}
}

var itemExists = struct{}{}

// New instantiates a new empty set and adds the passed values, if any, to the set
func New[T comparable](values ...T) *Set[T] {
	set := &Set[T]{items: make(map[T]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// NewFromMap instantiates a new  set from the provided slice.
func NewFromSlice[T comparable](slice []T) *Set[T] {
	s := &Set[T]{items: make(map[T]struct{})}

	for _, item := range slice {
		s.items[item] = itemExists
	}

	return s
}

// NewFromIterator instantiates a new set containing the elements provided by the passed iterator.
func NewFromIterator[T comparable](beign ds.ReadCompForIndexIterator[int, T]) *Set[T] {
	s := &Set[T]{items: make(map[T]struct{})}

	for beign.Next() {
		newValue, _ := beign.Get()

		s.items[newValue] = itemExists
	}

	return s
}

// NewFromIterators instantiates a new set containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[T comparable](begin ds.ReadCompForIndexIterator[int, T], end ds.CompIndexIterator[int]) *Set[T] {
	s := &Set[T]{items: make(map[T]struct{})}

	for !begin.IsEqual(end) && begin.Next() {
		newValue, _ := begin.Get()

		s.items[newValue] = itemExists
	}

	return s
}

// Add adds the items (one or more) to the set.
func (set *Set[T]) Add(items ...T) {
	for _, item := range items {
		set.items[item] = itemExists
	}
}

// Remove removes the items (one or more) from the set.
func (set *Set[T]) Remove(_ utils.Comparator[T], items ...T) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set[T]) IsEmpty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set[T]) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *Set[T]) Clear() {
	set.items = make(map[T]struct{})
}

// Values returns all items in the set.
func (set *Set[T]) GetValues() []T {
	values := make([]T, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// String returns a string representation of container
func (set *Set[T]) ToString() string {
	str := "HashSet\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}

// Intersection returns the intersection between two sets.
// The new set consists of all elements that are both in "set" and "other".
// Ref: https://en.wikipedia.org/wiki/Intersection_(set_theory)
func (set *Set[T]) MakeIntersectionWith(other sets.Set[T]) sets.Set[T] {
	result := New[T]()
	concrete := other.(*Set[T])

	// Iterate over smaller set (optimization)
	if set.Size() <= other.Size() {
		for item := range set.items {
			if _, contains := concrete.items[item]; contains {
				result.Add(item)
			}
		}
	} else {
		for item := range concrete.items {
			if _, contains := set.items[item]; contains {
				result.Add(item)
			}
		}
	}

	return result
}

// Union returns the union of two sets.
// The new set consists of all elements that are in "set" or "other" (possibly both).
// Ref: https://en.wikipedia.org/wiki/Union_(set_theory)
func (set *Set[T]) MakeUnionWith(other sets.Set[T]) sets.Set[T] {
	result := New[T]()
	concrete := other.(*Set[T])

	for item := range set.items {
		result.Add(item)
	}
	for item := range concrete.items {
		result.Add(item)
	}

	return result
}

// Difference returns the difference between two sets.
// The new set consists of all elements that are in "set" but not in "other".
// Ref: https://proofwiki.org/wiki/Definition:Set_Difference
func (set *Set[T]) MakeDifferenceWith(other sets.Set[T]) sets.Set[T] {
	result := New[T]()
	concrete := other.(*Set[T])

	for item := range set.items {
		if _, contains := concrete.items[item]; !contains {
			result.Add(item)
		}
	}

	return result
}

//******************************************************************//
//                         Ordered iterator                         //
//******************************************************************//

// OrderedBegin returns an initialized, reversed iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (s *Set[T]) OrderedBegin(comparator utils.Comparator[T]) ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return s.NewOrderedIterator(-1, s.Size(), comparator)
}

// OrderedEnd returns an initialized,reversed iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (s *Set[T]) OrderedEnd(comparator utils.Comparator[T]) ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return s.NewOrderedIterator(len(s.items), s.Size(), comparator)
}

// OrderedFirst returns an initialized, reversed iterator, which points to it's first element.
func (s *Set[T]) OrderedFirst(comparator utils.Comparator[T]) ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return s.NewOrderedIterator(0, s.Size(), comparator)
}

// OrderedLast returns an initialized, reversed iterator, which points to it's last element.
func (s *Set[T]) OrderedLast(comparator utils.Comparator[T]) ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return s.NewOrderedIterator(len(s.items)-1, s.Size(), comparator)
}
