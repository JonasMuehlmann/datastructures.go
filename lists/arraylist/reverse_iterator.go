// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arraylist

import "github.com/JonasMuehlmann/datastructures.go/ds"

// Assert Iterator implementation
var _ ds.RWOrdCompBidRevRandCollIterator[any, int] = (*ReverseIterator[any])(nil)

// ReverseIterator holding the iterator's state
type ReverseIterator[T any] struct {
	list  *List[T]
	index int
}

// NewReverseIterator returns a stateful iterator whose values can be fetched by an index.
func (list *List[T]) NewReverseIterator(list_ *List[T], index int) *ReverseIterator[T] {
	return &ReverseIterator[T]{list: list_, index: index}
}

// IsValid implements ds.RWOrdCompBidRandCollReverseIterator
func (it *ReverseIterator[T]) IsValid() bool {
	return it.list.withinRange(it.index)
}

// Get implements ds.RWOrdCompBidRandCollReverseIterator
func (it *ReverseIterator[T]) Get() (value T, found bool) {
	if it.list.size == 0 || !it.IsValid() {
		return
	}

	return it.list.elements[it.index], true
}

// Set implements ds.RWOrdCompBidRandCollReverseIterator
func (it *ReverseIterator[T]) Set(value T) bool {
	if it.list.size == 0 || !it.IsValid() {
		return false
	}
	it.list.elements[it.index] = value

	return true
}

// DistanceTo implements ds.RWOrdCompBidRandCollReverseIterator
// If other is of type CollectionIterator, CollectionIterator.Index() will be used, possibly executing in O(1)
func (it *ReverseIterator[T]) DistanceTo(other ds.OrderedIterator) int {
	forwardIterator := ds.RWOrdCompBidRandCollIterator[T, int](it)

	return -forwardIterator.DistanceTo(other)
}

// IsAfter implements ds.RWOrdCompBidRandCollReverseIterator
func (it *ReverseIterator[T]) IsAfter(other ds.OrderedIterator) bool {
	forwardIterator := ds.RWOrdCompBidRandCollIterator[T, int](it)

	return !forwardIterator.IsAfter(other)
}

// IsBefore implements ds.RWOrdCompBidRandCollReverseIterator
func (it *ReverseIterator[T]) IsBefore(other ds.OrderedIterator) bool {
	forwardIterator := ds.RWOrdCompBidRandCollIterator[T, int](it)

	return !forwardIterator.IsBefore(other)
}

// IsEqual implements ds.RWOrdCompBidRandCollReverseIterator
func (it *ReverseIterator[T]) IsEqual(other ds.ComparableIterator) bool {
	forwardIterator := ds.RWOrdCompBidRandCollIterator[T, int](it)

	return forwardIterator.IsEqual(other)
}

// Next implements ds.RWOrdCompBidRandCollReverseIterator
func (it *ReverseIterator[T]) Next() {
	it.index--
}

// NextN implements ds.RWOrdCompBidRandCollReverseIterator
func (it *ReverseIterator[T]) NextN(i int) {
	it.index -= i
}

// Previous implements ds.RWOrdCompBidRandCollReverseIterator
func (it *ReverseIterator[T]) Previous() {
	it.index++
}

// PreviousN implements ds.RWOrdCompBidRandCollReverseIterator
func (it *ReverseIterator[T]) PreviousN(n int) {
	it.index += n
}

// MoveBy implements ds.RWOrdCompBidRandCollReverseIterator
func (it *ReverseIterator[T]) MoveBy(n int) {
	it.index -= n
}

// Size implements ds.RWOrdCompBidRandCollReverseIterator
func (it *ReverseIterator[T]) Size() int {
	return it.list.size
}

// Index implements ds.RWOrdCompBidRandCollReverseIterator
func (it *ReverseIterator[T]) Index() int {
	return it.index
}

// MoveTo implements ds.RWOrdCompBidRandCollReverseIterator
func (it *ReverseIterator[T]) MoveTo(i int) {
	it.index = i
}

// IsBegin implements ds.ReverseIterator
func (it *ReverseIterator[T]) IsBegin() bool {
	return it.index == it.list.size
}

// IsEnd implements ds.ReverseIterator
func (it *ReverseIterator[T]) IsEnd() bool {
	return it.index == -1
}

// IsFirst implements ds.ReverseIterator
func (it *ReverseIterator[T]) IsFirst() bool {
	return it.index == it.list.size-1
}

// IsLast implements ds.ReverseIterator
func (it *ReverseIterator[T]) IsLast() bool {
	return it.index == 0
}

// PERF: Maybe an unsafe version is useful here
// GetAt implements ds.RWOrdCompBidRandCollIterator
func (it *ReverseIterator[T]) GetAt(i int) (value T, found bool) {
	if it.list.size == 0 || !it.list.withinRange(i) {
		return
	}

	return it.list.elements[i], true
}

// PERF: Maybe an unsafe version is useful here
// SetAt implements ds.RWOrdCompBidRandCollIterator
func (it *ReverseIterator[T]) SetAt(i int, value T) bool {
	if it.list.size == 0 || !it.list.withinRange(i) {
		return false
	}
	it.list.elements[i] = value

	return true
}
