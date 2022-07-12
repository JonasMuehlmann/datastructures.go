// Copyright (c) 2022, Jonass Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashmap implements a map backed by a hash table.
//
// Elements are unordered in the map.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Associative_array
package hashmap

import (
	"fmt"

	"github.com/JonasMuehlmann/datastructures.go/maps"
)

// Assert Map implementation.
var _ maps.Map[string, any] = (*Map[string, any])(nil)

// Map holds the elements in go's native map.
type Map[TKey comparable, TValue any] struct {
	m map[TKey]TValue
}

func (m *Map[TKey, TValue]) MergeWith(other *maps.Map[TKey, TValue]) bool {
	panic("Not implemented")
}

func (m *Map[TKey, TValue]) MergeWithSafe(other *maps.Map[TKey, TValue], overwriteOriginal bool) {
	panic("Not implemented")
}

// TODO: Implement NewFromMap() method which only copies map and not items
// New instantiates a hash map.
func New[TKey comparable, TValue any]() *Map[TKey, TValue] {
	return &Map[TKey, TValue]{m: make(map[TKey]TValue)}
}

// Put inserts element into the map.
func (m *Map[TKey, TValue]) Put(key TKey, value TValue) {
	m.m[key] = value
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map[TKey, TValue]) Get(key TKey) (value TValue, found bool) {
	value, found = m.m[key]
	return
}

// Remove removes the element from the map by key.
func (m *Map[TKey, TValue]) Remove(key TKey) {
	delete(m.m, key)
}

// Empty returns true if map does not contain any elements.
func (m *Map[TKey, TValue]) IsEmpty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map[TKey, TValue]) Size() int {
	return len(m.m)
}

// Keys returns all keys (random order).
func (m *Map[TKey, TValue]) Keys() []TKey {
	keys := make([]TKey, m.Size())
	count := 0
	for key := range m.m {
		keys[count] = key
		count++
	}
	return keys
}

// Values returns all values (random order).
func (m *Map[TKey, TValue]) GetValues() []TValue {
	values := make([]TValue, m.Size())
	count := 0
	for _, value := range m.m {
		values[count] = value
		count++
	}
	return values
}

// GetMap returns the underlying map.
func (map_ *Map[TKey, TValue]) GetMap() map[TKey]TValue {
	return map_.m
}

// Clear removes all elements from the map.
func (m *Map[TKey, TValue]) Clear() {
	m.m = make(map[TKey]TValue)
}

// String returns a string representation of container.
func (m *Map[TKey, TValue]) ToString() string {
	str := "HashMap\n"
	str += fmt.Sprintf("%v", m.m)
	return str
}
