// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treebidimap implements a bidirectional map backed by two red-black tree.
//
// This structure guarantees that the map will be in both ascending key and value order.
//
// Other than key and value ordering, the goal with this structure is to avoid duplication of elements, which can be significant if contained elements are large.
//
// A bidirectional map, or hash bag, is an associative data structure in which the (key,value) pairs form a one-to-one correspondence.
// Thus the binary relation is functional in each direction: value can also act as a key to key.
// A pair (a,b) thus provides a unique coupling between 'a' and 'b' so that 'b' can be found when 'a' is used as a key and 'a' can be found when 'b' is used as a key.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Bidirectional_map
package treebidimap

import (
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/maps"
	"github.com/JonasMuehlmann/datastructures.go/trees/redblacktree"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Map implementation
var _ maps.BidiMap[string, string] = (*Map[string, string])(nil)

// Map holds the elements in two red-black trees.
type Map[TKey comparable, TValue comparable] struct {
	forwardMap      redblacktree.Tree[TKey, TValue]
	inverseMap      redblacktree.Tree[TValue, TKey]
	keyComparator   utils.Comparator[TKey]
	valueComparator utils.Comparator[TValue]
}

type data[TKey comparable, TValue comparable] struct {
	key   TKey
	value TValue
}

// New instantiates a bidirectional map.
func New[TKey comparable, TValue comparable](keyComparator utils.Comparator[TKey], valueComparator utils.Comparator[TValue]) *Map[TKey, TValue] {
	return &Map[TKey, TValue]{
		forwardMap:      *redblacktree.New[TKey, TValue](keyComparator),
		inverseMap:      *redblacktree.New[TValue, TKey](valueComparator),
		keyComparator:   keyComparator,
		valueComparator: valueComparator,
	}
}

// NewFromMap instantiates a new tree containing the provided map.
func NewFromMap[TKey comparable, TValue comparable](keyComparator utils.Comparator[TKey], valueComparator utils.Comparator[TValue], map_ map[TKey]TValue) *Map[TKey, TValue] {
	tree := New[TKey, TValue](keyComparator, valueComparator)

	for k, v := range map_ {
		tree.Put(k, v)
	}

	return tree
}

// NewFromIterator instantiates a new tree containing the elements provided by the passed iterator.
func NewFromIterator[TKey comparable, TValue comparable](keyComparator utils.Comparator[TKey], valueComparator utils.Comparator[TValue], begin ds.ReadCompForIndexIterator[TKey, TValue]) *Map[TKey, TValue] {
	tree := New[TKey, TValue](keyComparator, valueComparator)

	for begin.Next() {
		newKey, _ := begin.Index()
		newValue, _ := begin.Get()

		tree.Put(newKey, newValue)
	}

	return tree
}

// NewFromIterators instantiates a new tree containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[TKey comparable, TValue comparable](keyComparator utils.Comparator[TKey], valueComparator utils.Comparator[TValue], begin ds.ReadCompForIndexIterator[TKey, TValue], end ds.CompIndexIterator[TKey]) *Map[TKey, TValue] {
	tree := New[TKey, TValue](keyComparator, valueComparator)

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

// Put inserts element into the map.
func (m *Map[TKey, TValue]) Put(key TKey, value TValue) {
	m.forwardMap.Put(key, value)
	m.inverseMap.Put(value, key)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[TKey, TValue]) Get(key TKey) (value TValue, found bool) {
	return m.forwardMap.Get(key)
}

// GetKey searches the element in the map by value and returns its key or nil if value is not found in map.
// Second return parameter is true if value was found, otherwise false.
func (m *Map[TKey, TValue]) GetKey(value TValue) (key TKey, found bool) {
	return m.inverseMap.Get(value)
}

// Remove removes the element from the map by key.
func (m *Map[TKey, TValue]) Remove(comparator utils.Comparator[TKey], key TKey) {
	if d, found := m.forwardMap.Get(key); found {
		m.forwardMap.Remove(key)
		m.inverseMap.Remove(d)
	}
}

// Empty returns true if map does not contain any elements
func (m *Map[TKey, TValue]) IsEmpty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[TKey, TValue]) Size() int {
	return m.forwardMap.Size()
}

// GetKeys returns all keys (ordered).
func (m *Map[TKey, TValue]) GetKeys() []TKey {
	return m.forwardMap.GetKeys()
}

// Values returns all values (ordered).
func (m *Map[TKey, TValue]) GetValues() []TValue {
	return m.inverseMap.GetKeys()
}

// Clear removes all elements from the map.
func (m *Map[TKey, TValue]) Clear() {
	m.forwardMap.Clear()
	m.inverseMap.Clear()
}

// String returns a string representation of container
func (m *Map[TKey, TValue]) ToString() string {
	str := "TreeBidiMap\nmap["
	it := m.OrderedBegin(m.keyComparator)
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
	return m.NewOrderedIterator(-1)
}

// OrderedEnd returns an initialized,reversed iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (m *Map[TKey, TValue]) OrderedEnd(comparator utils.Comparator[TKey]) ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewOrderedIterator(m.Size())
}

// OrderedFirst returns an initialized, reversed iterator, which points to it's first element.
func (m *Map[TKey, TValue]) OrderedFirst(comparator utils.Comparator[TKey]) ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewOrderedIterator(0)
}

// OrderedLast returns an initialized, reversed iterator, which points to it's last element.
func (m *Map[TKey, TValue]) OrderedLast(comparator utils.Comparator[TKey]) ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewOrderedIterator(m.Size() - 1)
}
