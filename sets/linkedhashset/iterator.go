// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashset

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/lists/doublylinkedlist"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[int, string] = (*Iterator[string])(nil)

type Iterator[T comparable] struct {
	s             *Set[T]
	orderIterator *doublylinkedlist.Iterator[T]
	comparator    utils.Comparator[T]
	// Redundant but has better locality
	value T
	index int
	size  int
}

func (m *Set[T]) NewIterator(s *Set[T], position int, comparator utils.Comparator[T]) *Iterator[T] {
	it := &Iterator[T]{
		s:             s,
		orderIterator: s.ordering.NewIterator(s.ordering, position),
		comparator:    comparator,
		index:         position,
		size:          s.Size(),
	}

	it.value, _ = it.orderIterator.Get()

	return it
}

// IsBegin implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsBegin() bool {
	return it.index == -1
}

// IsEnd implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsEnd() bool {
	return it.index == it.size
}

// IsFirst implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsFirst() bool {
	return it.index == 0
}

// IsLast implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsLast() bool {
	return it.index == it.size-1
}

// IsValid implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsValid() bool {
	return it.size > 0 && !it.IsBegin() && !it.IsEnd()
}

// IsEqual implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.index == otherThis.index

}

// DistanceTo implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) DistanceTo(other ds.OrderedIterator) int {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.index - otherThis.index
}

// IsAfter implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsAfter(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) > 0
}

// IsBefore implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsBefore(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) < 0
}

// Size implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Size() int {
	return it.size
}

// Index implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Index() (index int, found bool) {
	return it.index, it.IsValid()
}

// Next implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Next() bool {
	it.orderIterator.Next()
	it.index = utils.Min(it.index+1, it.size)

	var found bool

	it.value, found = it.orderIterator.Get()

	return found
}

// NextN implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) NextN(i int) bool {
	it.orderIterator.NextN(i)
	it.index = utils.Min(it.index+i, it.size)

	var found bool

	it.value, found = it.orderIterator.Get()

	return found
}

// Previous implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Previous() bool {
	it.orderIterator.Previous()
	it.index = utils.Max(it.index-1, -1)

	var found bool

	it.value, found = it.orderIterator.Get()

	return found
}

// PreviousN implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) PreviousN(n int) bool {
	it.orderIterator.PreviousN(n)
	it.index = utils.Max(it.index-n, -1)

	var found bool

	it.value, found = it.orderIterator.Get()

	return found
}

// MoveBy implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) MoveBy(n int) bool {
	if n > 0 {
		return it.NextN(n)
	}

	return it.PreviousN(-n)
}

// MoveTo implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) MoveTo(i int) bool {
	return it.MoveBy(i - it.index)
}

// Get implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Get() (value T, found bool) {
	return it.value, it.IsValid()
}

// GetAt implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) GetAt(i int) (value T, found bool) {
	return it.orderIterator.GetAt(i)
}

// Set implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Set(value T) bool {
	if !it.IsValid() {
		return false
	}

	it.orderIterator.Set(value)

	valueToRemove, _ := it.Get()
	it.s.Remove(it.comparator, valueToRemove)
	it.s.table[value] = itemExists
	it.value = value

	return true
}

// SetAt implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) SetAt(i int, value T) bool {
	if i < 0 || i >= it.Size() {
		return false
	}

	valueToRemove, _ := it.Get()
	delete(it.s.table, valueToRemove)
	it.s.table[value] = itemExists

	return true
}
