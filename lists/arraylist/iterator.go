// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arraylist

import "github.com/JonasMuehlmann/datastructures.go/ds"

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[any, int] = (*Iterator[any])(nil)

// Iterator holding the iterator's state
type Iterator[T any] struct {
	list  *List[T]
	index int
}

// NewIterator returns a stateful iterator whose values can be fetched by an index.
func (list *List[T]) NewIterator(list_ *List[T], index int) *Iterator[T] {
	return &Iterator[T]{list: list_, index: index}
}

// IsValid implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsValid() bool {
	return it.list.withinRange(it.index)
}

// PERF: Maybe an unsafe version is useful here
// Get implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Get() (value T, found bool) {
	if it.list.size == 0 || !it.IsValid() {
		return
	}

	return it.list.elements[it.index], true
}

// PERF: Maybe an unsafe version is useful here
// Set implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Set(value T) bool {
	if it.list.size == 0 || !it.IsValid() {
		return false
	}
	it.list.elements[it.index] = value

	return true
}

// DistanceTo implements ds.ReadWriteOrdCompBidRandCollIterator
// If other is of type CollectionIterator, CollectionIterator.Index() will be used, possibly executing in O(1)
func (it *Iterator[T]) DistanceTo(other ds.OrderedIterator) int {
	otherCollection, ok := other.(ds.CollectionIterator[int])

	if ok {
		return it.Index() - otherCollection.Index()
	} else {
		panic("Not implemented for non collection types")
	}
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
	return it.list.size
}

// Index implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Index() int {
	return it.index
}

// MoveTo implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) MoveTo(i int) {
	it.index = i
}

// IsBegin implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsBegin() bool {
	return it.index == -1
}

// IsEnd implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsEnd() bool {
	return it.index == -1 || it.index == it.list.size
}

// IsFirst implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsFirst() bool {
	return it.index == 0
}

// IsLast implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsLast() bool {
	return it.index == it.list.size-1
}

// PERF: Maybe an unsafe version is useful here
// GetAt implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) GetAt(i int) (value T, found bool) {
	if it.list.size == 0 || !it.list.withinRange(i) {
		return
	}

	return it.list.elements[i], true
}

// PERF: Maybe an unsafe version is useful here
// SetAt implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) SetAt(i int, value T) bool {
	if it.list.size == 0 || !it.list.withinRange(i) {
		return false
	}
	it.list.elements[i] = value

	return true
}
