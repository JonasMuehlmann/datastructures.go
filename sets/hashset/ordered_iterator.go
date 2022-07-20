// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashset

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[int, string] = (*OrderedIterator[string])(nil)

type OrderedIterator[T comparable] struct {
	s          *Set[T]
	values     []T
	index      int
	comparator utils.Comparator[T]
}

func (m *Set[T]) NewOrderedIterator(s *Set[T], position int, comparator utils.Comparator[T]) *OrderedIterator[T] {
	keys := s.GetValues()
	utils.Sort(keys, comparator)

	return &OrderedIterator[T]{
		s:          s,
		values:     keys,
		index:      position,
		comparator: comparator,
	}
}

// IsBegin implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) IsBegin() bool {
	return it.index == -1
}

// IsEnd implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) IsEnd() bool {
	return len(it.values) == 0 || it.index == len(it.values)
}

// IsFirst implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) IsFirst() bool {
	return it.index == 0
}

// IsLast implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) IsLast() bool {
	return it.index == len(it.values)-1
}

// IsValid implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) IsValid() bool {
	return len(it.values) > 0 && it.index >= 0 && it.index < len(it.values)
}

// IsEqual implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) == 0
}

// DistanceTo implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) DistanceTo(other ds.OrderedIterator) int {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.index - otherThis.index
}

// IsAfter implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) IsAfter(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) > 0
}

// IsBefore implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) IsBefore(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) < 0
}

// Size implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) Size() int {
	return len(it.values)
}

// Index implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) Index() (index int, found bool) {
	if !it.IsValid() {
		found = false
		return
	}

	index = it.index
	found = true

	return
}

// Next implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) Next() {
	it.index++
}

// NextN implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) NextN(i int) {
	it.index += i
}

// Previous implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) Previous() {
	it.index--
}

// PreviousN implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) PreviousN(n int) {
	it.index -= n
}

// MoveBy implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) MoveBy(n int) {
	it.index += n
}

// MoveTo implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) MoveTo(i int) bool {
	it.index = i

	return false
}

// Get implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) Get() (value T, found bool) {
	if !it.IsValid() {
		return
	}

	return it.values[it.index], true
}

// GetAt implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) GetAt(i int) (value T, found bool) {
	if i < 0 || i >= it.Size() {
		return
	}

	return it.values[i], true
}

// Set implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) Set(value T) bool {
	if !it.IsValid() {
		return false
	}

	it.values[it.index] = value

	return true
}

// SetAt implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) SetAt(i int, value T) bool {
	if i < 0 || i >= it.Size() {
		return false
	}

	delete(it.s.items, it.values[i])
	it.s.items[value] = itemExists

	it.values[i] = value

	return true
}
