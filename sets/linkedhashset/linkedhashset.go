// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linkedhashset is a set that preserves insertion-order.
//
// It is backed by a hash table to store values and doubly-linked list to store ordering.
//
// Note that insertion-order is not affected if an element is re-inserted into the set.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package linkedhashset

import (
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/lists/doublylinkedlist"
	"github.com/JonasMuehlmann/datastructures.go/sets"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Set implementation
var _ sets.Set[string] = (*Set[string])(nil)

// Set holds elements in go's native map
type Set[T comparable] struct {
	table      map[T]struct{}
	ordering   *doublylinkedlist.List[T]
	comparator utils.Comparator[T]
}

var itemExists = struct{}{}

// New instantiates a new empty set and adds the passed values, if any, to the set
func New[T comparable](values ...T) *Set[T] {
	set := &Set[T]{
		table:    make(map[T]struct{}),
		ordering: doublylinkedlist.New[T](),
	}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// NewFromMap instantiates a new  set from the provided slice.
func NewFromSlice[T comparable](slice []T) *Set[T] {
	s := &Set[T]{table: make(map[T]struct{}), ordering: doublylinkedlist.New[T]()}

	for _, item := range slice {
		s.table[item] = itemExists
		s.ordering.PushBack(item)
	}

	return s
}

// NewFromIterator instantiates a new set containing the elements provided by the passed iterator.
func NewFromIterator[T comparable](begin ds.ReadForIndexIterator[int, T]) *Set[T] {
	s := &Set[T]{table: make(map[T]struct{}), ordering: doublylinkedlist.New[T]()}

	for begin.Next() {
		newValue, _ := begin.Get()

		s.table[newValue] = itemExists
		s.ordering.PushBack(newValue)
	}

	return s
}

// NewFromIterators instantiates a new set containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[T comparable](begin ds.ReadCompForIndexIterator[int, T], end ds.CompIndexIterator[int]) *Set[T] {
	s := &Set[T]{table: make(map[T]struct{}), ordering: doublylinkedlist.New[T]()}

	for !begin.IsEqual(end) && begin.Next() {
		newValue, _ := begin.Get()

		s.table[newValue] = itemExists
		s.ordering.PushBack(newValue)
	}

	return s
}

// Add adds the items (one or more) to the set.
// Note that insertion-order is not affected if an element is re-inserted into the set.
func (set *Set[T]) Add(items ...T) {
	for _, item := range items {
		if _, contains := set.table[item]; !contains {
			set.table[item] = itemExists
			set.ordering.PushBack(item)
		}
	}
}

// Remove removes the items (one or more) from the set.
// Slow operation, worst-case O(n^2).
func (set *Set[T]) Remove(comparator utils.Comparator[T], items ...T) {
	for _, item := range items {
		if _, contains := set.table[item]; contains {
			delete(set.table, item)
			index := set.ordering.IndexOf(comparator, item)
			set.ordering.Remove(index)
		}
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		if _, contains := set.table[item]; !contains {
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
	return set.ordering.Size()
}

// Clear clears all values in the set.
func (set *Set[T]) Clear() {
	set.table = make(map[T]struct{})
	set.ordering.Clear()
}

// Values returns all items in the set.
func (set *Set[T]) GetValues() []T {
	values := make([]T, 0, set.Size())

	for it := set.First(set.comparator); !it.IsEnd(); it.Next() {
		value, _ := it.Get()
		values = append(values, value)
	}

	return values
}

// String returns a string representation of container
func (set *Set[T]) ToString() string {
	str := "LinkedHashSet\n"
	items := []string{}

	for it := set.First(set.comparator); !it.IsEnd(); it.Next() {
		value, _ := it.Get()
		items = append(items, fmt.Sprintf("%v", value))
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
		for item := range set.table {
			if _, contains := concrete.table[item]; contains {
				result.Add(item)
			}
		}
	} else {
		for item := range concrete.table {
			if _, contains := set.table[item]; contains {
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

	for item := range set.table {
		result.Add(item)
	}
	for item := range concrete.table {
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

	for item := range set.table {
		if _, contains := concrete.table[item]; !contains {
			result.Add(item)
		}
	}

	return result
}

//******************************************************************//
//                             iterator                             //
//******************************************************************//

// Begin returns an initialized, reversed iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (s *Set[T]) Begin(comparator utils.Comparator[T]) ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return s.NewIterator(-1, s.Size(), comparator)
}

// End returns an initialized,reversed iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (s *Set[T]) End(comparator utils.Comparator[T]) ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return s.NewIterator(s.Size(), s.Size(), comparator)
}

// First returns an initialized, reversed iterator, which points to it's first element.
func (s *Set[T]) First(comparator utils.Comparator[T]) ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return s.NewIterator(0, s.Size(), comparator)
}

// Last returns an initialized, reversed iterator, which points to it's last element.
func (s *Set[T]) Last(comparator utils.Comparator[T]) ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return s.NewIterator(s.Size()-1, s.Size(), comparator)
}
