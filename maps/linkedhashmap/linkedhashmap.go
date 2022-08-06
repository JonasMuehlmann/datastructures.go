// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linkedhashmap is a map that preserves insertion-order.
//
// It is backed by a hash table to store values and doubly-linked list to store ordering.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package linkedhashmap

import (
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/lists/doublylinkedlist"
	"github.com/JonasMuehlmann/datastructures.go/maps"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Map implementation
var _ maps.Map[string, any] = (*Map[string, any])(nil)

// Map holds the elements in a regular hash table, and uses doubly-linked list to store key ordering.
type Map[TKey comparable, TValue any] struct {
	table    map[TKey]TValue
	ordering *doublylinkedlist.List[TKey]
}

func (m *Map[TKey, TValue]) MergeWith(other *maps.Map[TKey, TValue]) bool {
	panic("Not implemented")
}

func (m *Map[TKey, TValue]) MergeWithSafe(other *maps.Map[TKey, TValue], overwriteOriginal bool) {
	panic("Not implemented")
}

// New instantiates a linked-hash-map.
func New[TKey comparable, TValue any]() *Map[TKey, TValue] {
	return &Map[TKey, TValue]{
		table:    make(map[TKey]TValue),
		ordering: doublylinkedlist.New[TKey](),
	}
}

// NewFromMap instantiates a new  set from the provided slice.
func NewFromMap[TKey comparable, TValue any](map_ map[TKey]TValue) *Map[TKey, TValue] {
	s := &Map[TKey, TValue]{table: make(map[TKey]TValue), ordering: doublylinkedlist.New[TKey]()}

	for k, v := range map_ {
		s.table[k] = v
		s.ordering.PushBack(k)
	}

	return s
}

// NewFromIterator instantiates a new set containing the elements provided by the passed iterator.
func NewFromIterator[TKey comparable, TValue any](begin ds.ReadCompForIndexIterator[TKey, TValue]) *Map[TKey, TValue] {
	s := &Map[TKey, TValue]{table: make(map[TKey]TValue), ordering: doublylinkedlist.New[TKey]()}

	for begin.Next() {
		newKey, _ := begin.Index()
		newValue, _ := begin.Get()

		s.table[newKey] = newValue
		s.ordering.PushBack(newKey)
	}

	return s
}

// NewFromIterators instantiates a new set containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[TKey comparable, TValue any](begin ds.ReadCompForIndexIterator[TKey, TValue], end ds.CompIndexIterator[TKey]) *Map[TKey, TValue] {
	s := &Map[TKey, TValue]{table: make(map[TKey]TValue), ordering: doublylinkedlist.New[TKey]()}

	for !begin.IsEqual(end) && begin.Next() {
		newKey, _ := begin.Index()
		newValue, _ := begin.Get()

		s.table[newKey] = newValue
		s.ordering.PushBack(newKey)
	}

	return s
}

// Put inserts key-value pair into the map.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[TKey, TValue]) Put(key TKey, value TValue) {
	if _, contains := m.table[key]; !contains {
		m.ordering.PushBack(key)
	}
	m.table[key] = value
}

// Get searches the element in the map by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[TKey, TValue]) Get(key TKey) (value TValue, found bool) {

	value, found = m.table[key]
	return
}

// Remove removes the element from the map by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (m *Map[TKey, TValue]) Remove(comparator utils.Comparator[TKey], key TKey) {
	if _, contains := m.table[key]; contains {
		delete(m.table, key)
		index := m.ordering.IndexOf(comparator, key)
		m.ordering.Remove(index)
	}
}

// Empty returns true if map does not contain any elements
func (m *Map[TKey, TValue]) IsEmpty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[TKey, TValue]) Size() int {
	return m.ordering.Size()
}

// GetKeys returns all keys in-order
func (m *Map[TKey, TValue]) GetKeys() []TKey {
	return m.ordering.GetValues()
}

// Values returns all values in-order based on the key.
func (m *Map[TKey, TValue]) GetValues() []TValue {
	values := make([]TValue, m.Size())
	count := 0
	it := m.Begin()
	for it.Next() {
		value, _ := it.Get()
		values[count] = value
		count++
	}
	return values
}

// Clear removes all elements from the map.
func (m *Map[TKey, TValue]) Clear() {
	m.table = make(map[TKey]TValue)
	m.ordering.Clear()
}

// String returns a string representation of container
func (m *Map[TKey, TValue]) ToString() string {
	str := "LinkedHashMap\nmap["
	it := m.Begin()
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
func (m *Map[TKey, TValue]) Begin() ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewIterator(-1)
}

// OrderedEnd returns an initialized,reversed iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (m *Map[TKey, TValue]) End() ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewIterator(m.Size())
}

// OrderedFirst returns an initialized, reversed iterator, which points to it's first element.
func (m *Map[TKey, TValue]) First() ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewIterator(0)
}

// OrderedLast returns an initialized, reversed iterator, which points to it's last element.
func (m *Map[TKey, TValue]) Last() ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewIterator(m.Size() - 1)
}
