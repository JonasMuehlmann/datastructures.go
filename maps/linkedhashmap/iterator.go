// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashmap

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/lists/doublylinkedlist"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[string, any] = (*Iterator[string, any])(nil)

type Iterator[TKey comparable, TValue any] struct {
	s             *Map[TKey, TValue]
	orderIterator *doublylinkedlist.Iterator[TKey]
	// Redundant but has better locality
	value TValue
	key   TKey
	index int
	size  int
}

func (m *Map[TKey, TValue]) NewIterator(s *Map[TKey, TValue], position int) *Iterator[TKey, TValue] {
	it := &Iterator[TKey, TValue]{
		s:             s,
		orderIterator: s.ordering.NewIterator(s.ordering, position),
		index:         position,
		size:          s.Size(),
	}

	it.key, _ = it.orderIterator.Get()
	it.value, _ = s.Get(it.key)

	return it
}

// IsBegin implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) IsBegin() bool {
	return it.index == -1
}

// IsEnd implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) IsEnd() bool {
	return it.index == it.size
}

// IsFirst implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}

// IsLast implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) IsLast() bool {
	return it.index == it.size-1
}

// IsValid implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) IsValid() bool {
	return it.size > 0 && !it.IsBegin() && !it.IsEnd()
}

// IsEqual implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*Iterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.index == otherThis.index

}

// DistanceTo implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) DistanceTo(other ds.OrderedIterator) int {
	otherThis, ok := other.(*Iterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.index - otherThis.index
}

// IsAfter implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) IsAfter(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*Iterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) > 0
}

// IsBefore implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) IsBefore(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*Iterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) < 0
}

// Size implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) Size() int {
	return it.size
}

// Index implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) Index() (index TKey, found bool) {
	return it.key, it.IsValid()
}

// Next implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) Next() bool {
	it.orderIterator.Next()
	it.index = utils.Min(it.index+1, it.size)

	var found bool

	it.key, found = it.orderIterator.Get()

	return found
}

// NextN implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) NextN(i int) bool {
	it.orderIterator.NextN(i)
	it.index = utils.Min(it.index+i, it.size)

	var found bool

	it.key, found = it.orderIterator.Get()

	return found
}

// Previous implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) Previous() bool {
	it.orderIterator.Previous()
	it.index = utils.Max(it.index-1, -1)

	var found bool

	it.key, found = it.orderIterator.Get()

	return found
}

// PreviousN implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) PreviousN(n int) bool {
	it.orderIterator.PreviousN(n)
	it.index = utils.Max(it.index-n, -1)

	var found bool

	it.key, found = it.orderIterator.Get()

	return found
}

// MoveBy implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) MoveBy(n int) bool {
	if n > 0 {
		return it.NextN(n)
	} else if n < 0 {
		return it.PreviousN(-n)
	}

	return it.IsValid()
}

// MoveTo implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) MoveTo(key TKey) bool {
	// PERF: This can be optimized
	for it.Next() {
		if it.key == key {
			return true
		}
	}
	for it.Previous() {
		if it.key == key {
			return true
		}
	}

	return false
}

// Get implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) Get() (value TValue, found bool) {
	return it.s.Get(it.key)
}

// GetAt implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) GetAt(key TKey) (value TValue, found bool) {
	return it.s.Get(key)
}

// Set implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) Set(value TValue) bool {
	if !it.IsValid() {
		return false
	}

	it.value = value
	it.s.Put(it.key, value)

	return true
}

// SetAt implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[TKey, TValue]) SetAt(i TKey, value TValue) bool {
	it.s.Put(it.key, value)

	return true
}
