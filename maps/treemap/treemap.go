// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treemap implements a map backed by red-black tree.
//
// Elements are ordered by key in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package treemap

import (
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/maps"
	rbt "github.com/JonasMuehlmann/datastructures.go/trees/redblacktree"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Map implementation
var _ maps.Map[string, any] = (*Map[string, any])(nil)

// Map holds the elements in a red-black tree
type Map[TKey comparable, TValue any] struct {
	tree *rbt.Tree[TKey, TValue]
}

// New instantiates a tree map with the custom comparator.
func New[TKey comparable, TValue any](comparator utils.Comparator[TKey]) *Map[TKey, TValue] {
	return &Map[TKey, TValue]{tree: rbt.New[TKey, TValue](comparator)}
}

// NewFromMap instantiates a new tree containing the provided map.
func NewFromMap[TKey comparable, TValue any](comparator utils.Comparator[TKey], map_ map[TKey]TValue) *Map[TKey, TValue] {
	tree := New[TKey, TValue](comparator)

	for k, v := range map_ {
		tree.Put(k, v)
	}

	return tree
}

// NewFromIterator instantiates a new tree containing the elements provided by the passed iterator.
func NewFromIterator[TKey comparable, TValue any](comparator utils.Comparator[TKey], begin ds.ReadCompForIndexIterator[TKey, TValue]) *Map[TKey, TValue] {
	tree := New[TKey, TValue](comparator)

	for begin.Next() {
		newKey, _ := begin.Index()
		newValue, _ := begin.Get()

		tree.Put(newKey, newValue)
	}

	return tree
}

// NewFromIterators instantiates a new tree containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[TKey comparable, TValue any](comparator utils.Comparator[TKey], begin ds.ReadCompForIndexIterator[TKey, TValue], end ds.CompIndexIterator[TKey]) *Map[TKey, TValue] {
	tree := New[TKey, TValue](comparator)

	for !begin.IsEqual(end) && begin.Next() {
		newKey, _ := begin.Index()
		newValue, _ := begin.Get()

		tree.Put(newKey, newValue)
	}

	return tree
}

func (m *Map[TKey, TValue]) MergeWith(other *maps.Map[TKey, TValue]) bool {
	panic("Not implemented")
}

func (m *Map[TKey, TValue]) MergeWithSafe(other *maps.Map[TKey, TValue], overwriteOriginal bool) {
	panic("Not implemented")
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[TKey, TValue]) Put(key TKey, value TValue) {
	m.tree.Put(key, value)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[TKey, TValue]) Get(key TKey) (value TValue, found bool) {
	return m.tree.Get(key)
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[TKey, TValue]) Remove(comparator utils.Comparator[TKey], key TKey) {
	m.tree.Remove(key)
}

// Empty returns true if map does not contain any elements
func (m *Map[TKey, TValue]) IsEmpty() bool {
	return m.tree.IsEmpty()
}

// Size returns number of elements in the map.
func (m *Map[TKey, TValue]) Size() int {
	return m.tree.Size()
}

// GetKeys returns all keys in-order
func (m *Map[TKey, TValue]) GetKeys() []TKey {
	return m.tree.GetKeys()
}

// Values returns all values in-order based on the key.
func (m *Map[TKey, TValue]) GetValues() []TValue {
	return m.tree.GetValues()
}

// Clear removes all elements from the map.
func (m *Map[TKey, TValue]) Clear() {
	m.tree.Clear()
}

// Min returns the minimum key and its value from the tree map.
// Returns nil, nil if map is empty.
func (m *Map[TKey, TValue]) Min() (key TKey, value TValue) {
	if node := m.tree.Left(); node != nil {
		return node.Key, node.Value
	}
	return
}

// Max returns the maximum key and its value from the tree map.
// Returns nil, nil if map is empty.
func (m *Map[TKey, TValue]) Max() (key TKey, value TValue) {
	if node := m.tree.Right(); node != nil {
		return node.Key, node.Value
	}
	return
}

// Floor finds the floor key-value pair for the input key.
// In case that no floor is found, then both returned values will be nil.
// It's generally enough to check the first value (key) for nil, which determines if floor was found.
//
// Floor key is defined as the largest key that is smaller than or equal to the given key.
// A floor key may not be found, either because the map is empty, or because
// all keys in the map are larger than the given key.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[TKey, TValue]) Floor(key TKey) (foundkey TKey, foundvalue TValue) {
	node, found := m.tree.Floor(key)
	if found {
		return node.Key, node.Value
	}
	return
}

// Ceiling finds the ceiling key-value pair for the input key.
// In case that no ceiling is found, then both returned values will be nil.
// It's generally enough to check the first value (key) for nil, which determines if ceiling was found.
//
// Ceiling key is defined as the smallest key that is larger than or equal to the given key.
// A ceiling key may not be found, either because the map is empty, or because
// all keys in the map are smaller than the given key.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[TKey, TValue]) Ceiling(key TKey) (foundkey TKey, foundvalue TValue) {
	node, found := m.tree.Ceiling(key)
	if found {
		return node.Key, node.Value
	}
	return
}

// String returns a string representation of container
func (m *Map[TKey, TValue]) ToString() string {
	str := "TreeMap\nmap["
	it := m.OrderedBegin(m.tree.Comparator)
	for it.Next() {
		key, _ := it.Index()
		value, _ := it.Get()

		str += fmt.Sprintf("%v:%v ", key, value)
	}
	return strings.TrimRight(str, " ") + "]"

}

//******************************************************************//
//                         Ordered iterator                         //
//******************************************************************//

// OrderedBegin returns an initialized, reversed iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (m *Map[TKey, TValue]) OrderedBegin(comparator utils.Comparator[TKey]) ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewOrderedIterator(m, -1)
}

// OrderedEnd returns an initialized,reversed iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (m *Map[TKey, TValue]) OrderedEnd(comparator utils.Comparator[TKey]) ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewOrderedIterator(m, m.Size())
}

// OrderedFirst returns an initialized, reversed iterator, which points to it's first element.
func (m *Map[TKey, TValue]) OrderedFirst(comparator utils.Comparator[TKey]) ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewOrderedIterator(m, 0)
}

// OrderedLast returns an initialized, reversed iterator, which points to it's last element.
func (m *Map[TKey, TValue]) OrderedLast(comparator utils.Comparator[TKey]) ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewOrderedIterator(m, m.Size()-1)
}
