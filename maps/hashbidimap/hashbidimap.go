// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashbidimap implements a bidirectional map backed by two hashmaps.
//
// A bidirectional map, or hash bag, is an associative data structure in which the (key,value) pairs form a one-to-one correspondence.
// Thus the binary relation is functional in each direction: value can also act as a key to key.
// A pair (a,b) thus provides a unique coupling between 'a' and 'b' so that 'b' can be found when 'a' is used as a key and 'a' can be found when 'b' is used as a key.
//
// Elements are unordered in the map.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Bidirectional_map
package hashbidimap

import (
	"fmt"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/maps"
	"github.com/JonasMuehlmann/datastructures.go/maps/hashmap"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// TODO: Should be implement an Equaler interface?
// Assert Map implementation
var _ maps.BidiMap[string, string] = (*Map[string, string])(nil)

// Map holds the elements in two hashmaps.
type Map[TKey comparable, TValue comparable] struct {
	forwardMap      *hashmap.Map[TKey, TValue]
	inverseMap      *hashmap.Map[TValue, TKey]
	keyComparator   utils.Comparator[TKey]
	valueComparator utils.Comparator[TValue]
}

func (m *Map[TKey, TValue]) MergeWith(other *maps.Map[TKey, TValue]) bool {
	panic("Not implemented")
}

func (m *Map[TKey, TValue]) MergeWithSafe(other *maps.Map[TKey, TValue], overwriteOriginal bool) {
	panic("Not implemented")
}

// New instantiates a bidirectional map.
func New[TKey comparable, TValue comparable](keyComparator utils.Comparator[TKey], valueComparator utils.Comparator[TValue]) *Map[TKey, TValue] {
	return &Map[TKey, TValue]{
		forwardMap:      hashmap.New[TKey, TValue](),
		inverseMap:      hashmap.New[TValue, TKey](),
		keyComparator:   keyComparator,
		valueComparator: valueComparator,
	}
}

// NewFromMap instantiates a new  map containing the provided map.
func NewFromMap[TKey comparable, TValue comparable](keyComparator utils.Comparator[TKey], valueComparator utils.Comparator[TValue], map_ map[TKey]TValue) *Map[TKey, TValue] {
	m := &Map[TKey, TValue]{
		forwardMap:      hashmap.NewFromMap(map_),
		inverseMap:      hashmap.New[TValue, TKey](),
		keyComparator:   keyComparator,
		valueComparator: valueComparator,
	}

	for k, v := range map_ {
		m.inverseMap.Put(v, k)
	}

	return m
}

// NewFromIterator instantiates a new list containing the elements provided by the passed iterator.
func NewFromIterator[TKey comparable, TValue comparable](keyComparator utils.Comparator[TKey], valueComparator utils.Comparator[TValue], begin ds.ReadCompForIndexIterator[TKey, TValue]) *Map[TKey, TValue] {
	forwardMap := make(map[TKey]TValue)
	inverseMap := make(map[TValue]TKey)

	for begin.Next() {
		newKey, _ := begin.GetKey()
		newValue, _ := begin.Get()

		forwardMap[newKey] = newValue
		inverseMap[newValue] = newKey
	}

	m := &Map[TKey, TValue]{
		forwardMap:      hashmap.NewFromMap(forwardMap),
		inverseMap:      hashmap.NewFromMap(inverseMap),
		keyComparator:   keyComparator,
		valueComparator: valueComparator,
	}

	return m
}

// NewFromIterators instantiates a new list containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[TKey comparable, TValue comparable](keyComparator utils.Comparator[TKey], valueComparator utils.Comparator[TValue], begin ds.ReadCompForIndexIterator[TKey, TValue], end ds.CompIndexIterator[TKey]) *Map[TKey, TValue] {
	forwardMap := make(map[TKey]TValue)
	inverseMap := make(map[TValue]TKey)

	for !begin.IsEqual(end) && begin.Next() {
		newKey, _ := begin.GetKey()
		newValue, _ := begin.Get()

		forwardMap[newKey] = newValue
		inverseMap[newValue] = newKey
	}

	m := &Map[TKey, TValue]{
		forwardMap:      hashmap.NewFromMap(forwardMap),
		inverseMap:      hashmap.NewFromMap(inverseMap),
		keyComparator:   keyComparator,
		valueComparator: valueComparator,
	}

	return m
}

// Put inserts element into the map.
func (m *Map[TKey, TValue]) Put(key TKey, value TValue) {
	if valueByKey, ok := m.forwardMap.Get(key); ok {
		m.inverseMap.Remove(m.valueComparator, valueByKey)
	}
	if keyByValue, ok := m.inverseMap.Get(value); ok {
		m.forwardMap.Remove(m.keyComparator, keyByValue)
	}
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
func (m *Map[TKey, TValue]) Remove(keyComparator utils.Comparator[TKey], key TKey) {
	if value, found := m.forwardMap.Get(key); found {
		m.forwardMap.Remove(keyComparator, key)
		m.inverseMap.Remove(m.valueComparator, value)
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

// GetKeys returns all keys (random order).
func (m *Map[TKey, TValue]) GetKeys() []TKey {
	return m.forwardMap.GetKeys()
}

// Values returns all values (random order).
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
	str := "HashBidimap\n"
	str += fmt.Sprintf("%v", m.forwardMap)
	return str
}

//******************************************************************//
//                         Ordered iterator                         //
//******************************************************************//

// OrderedBegin returns an initialized, reversed iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (m *Map[TKey, TValue]) OrderedBegin(comparator utils.Comparator[TKey]) ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewOrderedIterator(-1, m.Size())
}

// OrderedEnd returns an initialized,reversed iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.

func (m *Map[TKey, TValue]) OrderedEnd(comparator utils.Comparator[TKey]) ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewOrderedIterator(m.Size(), m.Size())
}

// OrderedFirst returns an initialized, reversed iterator, which points to it's first element.

func (m *Map[TKey, TValue]) OrderedFirst(comparator utils.Comparator[TKey]) ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewOrderedIterator(0, m.Size())
}

// OrderedLast returns an initialized, reversed iterator, which points to it's last element.
func (m *Map[TKey, TValue]) OrderedLast(comparator utils.Comparator[TKey]) ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return m.NewOrderedIterator(m.Size()-1, m.Size())
}
