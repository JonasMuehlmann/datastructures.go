// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedliststack

import "github.com/JonasMuehlmann/datastructures.go/ds"

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[int, any] = (*Iterator[any])(nil)

// Iterator holding the iterator's state
type Iterator[T any] struct {
	stack *Stack[T]
	index int
}

// NewIterator returns a stateful iterator whose values can be fetched by an index.
func (list *Stack[T]) NewIterator(list_ *Stack[T], index int) *Iterator[T] {
	return &Iterator[T]{stack: list_, index: index}
}

// IsValid implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsValid() bool {
	return it.stack.withinRange(it.index)
}

// Get implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Get() (value T, found bool) {
	if it.stack.Size() == 0 || !it.IsValid() {
		return
	}

	return it.stack.list.Get(it.index)
}

// Set implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Set(value T) bool {
	if it.stack.Size() == 0 || !it.IsValid() {
		return false
	}

	it.stack.list.Set(it.index, value)

	return true
}

// DistanceTo implements ds.ReadWriteOrdCompBidRandCollIterator
// If other is of type IndexedIterator, IndexedIterator.Index() will be used, possibly executing in O(1)
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

// IsEqual implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) == 0
}

// Next implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Next() {
	it.index++
}

// NextN implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) NextN(i int) {
	it.index += i
}

// Previous implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Previous() {
	it.index--
}

// PreviousN implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) PreviousN(n int) {
	it.index -= n
}

// MoveBy implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) MoveBy(n int) {
	it.index += n
}

// Size implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Size() int {
	return it.stack.Size()
}

// Index implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Index() (int, bool) {
	return it.index, true
}

// MoveTo implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) MoveTo(i int) bool {
	it.index = i

	return true
}

// IsBegin implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsBegin() bool {
	return it.index == -1
}

// IsEnd implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsEnd() bool {
	return it.stack.Size() == 0 || it.index == it.stack.Size()
}

// IsFirst implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsFirst() bool {
	return it.index == 0
}

// IsLast implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsLast() bool {
	return it.index == it.stack.Size()-1
}

// GetAt implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) GetAt(i int) (value T, found bool) {
	if it.stack.Size() == 0 || !it.stack.withinRange(i) {
		return
	}

	return it.stack.list.Get(i)
}

// SetAt implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) SetAt(i int, value T) bool {
	if it.stack.Size() == 0 || !it.stack.withinRange(i) {
		return false
	}
	it.stack.list.Set(i, value)

	return true
}
