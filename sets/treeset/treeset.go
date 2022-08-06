// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treeset implements a tree backed by a red-black tree.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package treeset

import (
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/sets"
	rbt "github.com/JonasMuehlmann/datastructures.go/trees/redblacktree"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Set implementation
var _ sets.Set[string] = (*Set[string])(nil)

// Set holds elements in a red-black tree
type Set[T comparable] struct {
	tree *rbt.Tree[T, struct{}]
}

var itemExists = struct{}{}

// New instantiates a new empty set with the custom comparator.
func New[T comparable](comparator utils.Comparator[T], values ...T) *Set[T] {
	set := &Set[T]{tree: rbt.New[T, struct{}](comparator)}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// NewFromMap instantiates a new  set from the provided slice.
func NewFromSlice[T comparable](comparator utils.Comparator[T], slice []T) *Set[T] {
	set := &Set[T]{tree: rbt.New[T, struct{}](comparator)}

	for _, value := range slice {
		set.Add(value)
	}

	return set
}

// NewFromIterator instantiates a new set containing the elements provided by the passed iterator.
func NewFromIterator[T comparable](comparator utils.Comparator[T], begin ds.ReadCompForIndexIterator[int, T]) *Set[T] {
	set := &Set[T]{tree: rbt.New[T, struct{}](comparator)}

	for begin.Next() {
		newValue, _ := begin.Get()

		set.Add(newValue)
	}

	return set
}

// NewFromIterators instantiates a new set containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[T comparable](comparator utils.Comparator[T], begin ds.ReadCompForIndexIterator[int, T], end ds.CompIndexIterator[int]) *Set[T] {
	set := &Set[T]{tree: rbt.New[T, struct{}](comparator)}

	for !begin.IsEqual(end) && begin.Next() {
		newValue, _ := begin.Get()

		set.Add(newValue)
	}

	return set
}

// Add adds the items (one or more) to the set.
func (set *Set[T]) Add(items ...T) {
	for _, item := range items {
		set.tree.Put(item, itemExists)
	}
}

// Remove removes the items (one or more) from the set.
func (set *Set[T]) Remove(_ utils.Comparator[T], items ...T) {
	for _, item := range items {
		set.tree.Remove(item)
	}
}

// Contains checks weather items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set[T]) Contains(items ...T) bool {
	for _, item := range items {
		if _, contains := set.tree.Get(item); !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set[T]) IsEmpty() bool {
	return set.tree.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set[T]) Size() int {
	return set.tree.Size()
}

// Clear clears all values in the set.
func (set *Set[T]) Clear() {
	set.tree.Clear()
}

// Values returns all items in the set.
func (set *Set[T]) GetValues() []T {
	return set.tree.GetKeys()
}

// String returns a string representation of container
func (set *Set[T]) ToString() string {
	str := "TreeSet\n"
	items := []string{}
	for _, v := range set.tree.GetKeys() {
		items = append(items, fmt.Sprintf("%v", v))
	}
	str += strings.Join(items, ", ")
	return str
}

// Intersection returns the intersection between two sets.
// The new set consists of all elements that are both in "set" and "other".
// The two sets should have the same comparators, otherwise the result is empty set.
// Ref: https://en.wikipedia.org/wiki/Intersection_(set_theory)
func (set *Set[T]) MakeIntersectionWith(other sets.Set[T]) sets.Set[T] {
	result := New(set.tree.Comparator)
	concrete := other.(*Set[T])

	// Iterate over smaller set (optimization)
	if set.Size() <= other.Size() {
		it := set.OrderedBegin(set.tree.Comparator)
		for it.Next() {
			value, _ := it.Get()
			if other.Contains(value) {
				result.Add(value)
			}
		}
	} else {
		it := concrete.OrderedBegin(concrete.tree.Comparator)
		for it.Next() {
			value, _ := it.Get()
			if set.Contains(value) {
				result.Add(value)
			}
		}
	}

	return result
}

// Union returns the union of two sets.
// The new set consists of all elements that are in "set" or "other" (possibly both).
// The two sets should have the same comparators, otherwise the result is empty set.
// Ref: https://en.wikipedia.org/wiki/Union_(set_theory)
func (set *Set[T]) MakeUnionWith(other sets.Set[T]) sets.Set[T] {
	result := New(set.tree.Comparator)
	concrete := other.(*Set[T])

	it := set.OrderedBegin(set.tree.Comparator)
	for it.Next() {
		value, _ := it.Get()
		result.Add(value)
	}

	it = concrete.OrderedBegin(concrete.tree.Comparator)
	for it.Next() {
		value, _ := it.Get()
		result.Add(value)
	}

	return result
}

// Difference returns the difference between two sets.
// The two sets should have the same comparators, otherwise the result is empty set.
// The new set consists of all elements that are in "set" but not in "other".
// Ref: https://proofwiki.org/wiki/Definition:Set_Difference
func (set *Set[T]) MakeDifferenceWith(other sets.Set[T]) sets.Set[T] {
	result := New(set.tree.Comparator)
	concrete := other.(*Set[T])

	it := set.OrderedBegin(set.tree.Comparator)
	for it.Next() {
		value, _ := it.Get()
		if !concrete.Contains(value) {
			result.Add(value)
		}
	}

	return result
}

//******************************************************************//
//                             iterator                             //
//******************************************************************//

// Begin returns an initialized, reversed iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (s *Set[T]) OrderedBegin(comparator utils.Comparator[T]) ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return s.NewOrderedIterator(-1)
}

// End returns an initialized,reversed iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (s *Set[T]) OrderedEnd(comparator utils.Comparator[T]) ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return s.NewOrderedIterator(s.Size())
}

// First returns an initialized, reversed iterator, which points to it's first element.
func (s *Set[T]) OrderedFirst(comparator utils.Comparator[T]) ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return s.NewOrderedIterator(0)
}

// Last returns an initialized, reversed iterator, which points to it's last element.
func (s *Set[T]) OrderedLast(comparator utils.Comparator[T]) ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return s.NewOrderedIterator(s.Size() - 1)
}
