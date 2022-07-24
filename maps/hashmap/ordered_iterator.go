// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashmap

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[string, any] = (*OrderedIterator[string, any])(nil)

type OrderedIterator[TKey comparable, TValue any] struct {
	m          *Map[TKey, TValue]
	keys       []TKey
	index      int
	comparator utils.Comparator[TKey]
	// Redundant but has better locality
	key   TKey
	value TValue
	size  int
}

func (m *Map[TKey, TValue]) NewOrderedIterator(m_ *Map[TKey, TValue], position int, comparator utils.Comparator[TKey]) *OrderedIterator[TKey, TValue] {
	keys := m_.GetKeys()
	utils.Sort(keys, comparator)

	it := &OrderedIterator[TKey, TValue]{
		m:          m_,
		keys:       keys,
		index:      position,
		comparator: comparator,
		size:       m_.Size(),
	}

	if position > 0 && position < len(keys) {
		it.MoveTo(it.keys[position])
	}

	return it
}

// IsBegin implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsBegin() bool {
	return it.index == -1
}

// IsEnd implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsEnd() bool {
	return len(it.keys) == 0 || it.index == len(it.keys)
}

// IsFirst implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}

// IsLast implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsLast() bool {
	return it.index == len(it.keys)-1
}

// IsValid implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsValid() bool {
	return len(it.keys) > 0 && it.index >= 0 && it.index < len(it.keys)
}

// IsEqual implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*OrderedIterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) == 0
}

// DistanceTo implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) DistanceTo(other ds.OrderedIterator) int {
	otherThis, ok := other.(*OrderedIterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.index - otherThis.index
}

// IsAfter implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsAfter(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*OrderedIterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) > 0
}

// IsBefore implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsBefore(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*OrderedIterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) < 0
}

// Size implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) Size() int {
	return len(it.keys)
}

// Index implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) Index() (key TKey, found bool) {
	if !it.IsValid() {
		found = false
		return
	}

	key = it.keys[it.index]
	found = true

	return
}

// Next implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) Next() bool {
	it.index = utils.Min(it.index+1, it.size)

	if !it.IsValid() {
		return false
	}

	it.key = it.keys[it.index]
	it.value = it.m.m[it.key]

	return true
}

// NextN implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) NextN(i int) bool {
	it.index = utils.Min(it.index+i, it.size)

	if !it.IsValid() {
		return false
	}

	it.key = it.keys[it.index]
	it.value = it.m.m[it.key]

	return true
}

// Previous implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) Previous() bool {
	it.index = utils.Max(it.index-1, -1)

	if !it.IsValid() {
		return false
	}

	it.key = it.keys[it.index]
	it.value = it.m.m[it.key]

	return true
}

// PreviousN implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) PreviousN(n int) bool {
	it.index = utils.Max(it.index-n, -1)

	if !it.IsValid() {
		return false
	}

	it.key = it.keys[it.index]
	it.value = it.m.m[it.key]

	return true
}

// MoveBy implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) MoveBy(n int) bool {
	if n > 0 {
		return it.NextN(n)
	}

	return it.PreviousN(-n)
}

// MoveTo implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) MoveTo(k TKey) bool {
	for i, key := range it.keys {
		if key == k {
			it.index = i
			it.key = key
			it.value = it.m.m[key]

			return true
		}
	}

	return false
}

// Get implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) Get() (value TValue, found bool) {
	if !it.IsValid() {
		return
	}

	return it.m.Get(it.keys[it.index])
}

// GetAt implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) GetAt(i TKey) (value TValue, found bool) {
	if !it.IsValid() {
		return
	}

	return it.m.Get(i)
}

// Set implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) Set(value TValue) bool {
	if !it.IsValid() {
		return false
	}

	it.m.Put(it.keys[it.index], value)

	return true
}

// SetAt implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) SetAt(i TKey, value TValue) bool {
	it.m.Put(i, value)

	return true
}
