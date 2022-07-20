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
}

func (m *Set[T]) NewIterator(s *Set[T], position int, comparator utils.Comparator[T]) *Iterator[T] {
	return &Iterator[T]{
		s:             s,
		comparator:    comparator,
		orderIterator: s.ordering.NewIterator(s.ordering, position),
	}
}

// IsBegin implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsBegin() bool {
	return it.orderIterator.IsBegin()
}

// IsEnd implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsEnd() bool {
	return it.orderIterator.IsEnd()
}

// IsFirst implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsFirst() bool {
	return it.orderIterator.IsFirst()
}

// IsLast implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsLast() bool {
	return it.orderIterator.IsLast()
}

// IsValid implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsValid() bool {
	return it.orderIterator.IsValid()
}

// IsEqual implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.orderIterator.IsEqual(otherThis.orderIterator)

}

// DistanceTo implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) DistanceTo(other ds.OrderedIterator) int {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.orderIterator.DistanceTo(otherThis.orderIterator)
}

// IsAfter implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsAfter(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.orderIterator.IsAfter(otherThis.orderIterator)
}

// IsBefore implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsBefore(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.orderIterator.IsBefore(otherThis.orderIterator)
}

// Size implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Size() int {
	return it.s.ordering.Size()
}

// Index implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Index() (index int, found bool) {
	return it.orderIterator.Index()
}

// Next implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Next() {
	it.orderIterator.Next()
}

// NextN implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) NextN(i int) {
	it.orderIterator.NextN(i)
}

// Previous implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Previous() {
	it.orderIterator.Previous()
}

// PreviousN implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) PreviousN(n int) {
	it.orderIterator.PreviousN(n)
}

// MoveBy implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) MoveBy(n int) {
	it.orderIterator.MoveBy(n)
}

// MoveTo implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) MoveTo(i int) bool {
	return it.orderIterator.MoveTo(i)
}

// Get implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Get() (value T, found bool) {
	return it.orderIterator.Get()
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
