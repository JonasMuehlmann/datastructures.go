// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arraylist

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[int, any] = (*ReverseIterator[any])(nil)

// Iterator holding the iterator's state
type ReverseIterator[T any] struct {
	list  *List[T]
	value T
	index int
	// Redundant but has better locality
	size int
}

// NewIterator returns a stateful iterator whose values can be fetched by an index.
func (list *List[T]) NewReverseIterator(index int) *ReverseIterator[T] {
	it := &ReverseIterator[T]{list: list, index: index, size: list.Size()}

	if it.IsValid() {
		it.value = list.elements[it.index]
	}

	return it
}

func (it *ReverseIterator[T]) IsValid() bool {
	return it.size > 0 && !it.IsBegin() && !it.IsEnd()
}

func (it *ReverseIterator[T]) Get() (value T, found bool) {
	if !it.IsValid() {
		return
	}

	return it.list.elements[it.index], true
}

func (it *ReverseIterator[T]) Set(value T) bool {
	if !it.IsValid() {
		return false
	}
	it.list.elements[it.index] = value

	return true
}

// If other is of type IndexedIterator, IndexedIterator.Index() will be used, possibly executing in O(1)
func (it *ReverseIterator[T]) DistanceTo(other ds.OrderedIterator) int {
	otherThis, ok := other.(*ReverseIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return otherThis.index - it.index
}

func (it *ReverseIterator[T]) IsAfter(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*ReverseIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) > 0
}

func (it *ReverseIterator[T]) IsBefore(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*ReverseIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) < 0
}

func (it *ReverseIterator[T]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*ReverseIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) == 0
}

func (it *ReverseIterator[T]) Previous() bool {
	it.index = utils.Min(it.index+1, it.size)

	if !it.IsValid() {
		return false
	}

	it.value = it.list.elements[it.index]

	return true
}

func (it *ReverseIterator[T]) PreviousN(i int) bool {
	it.index = utils.Min(it.index+i, it.size)

	if !it.IsValid() {
		return false
	}

	it.value = it.list.elements[it.index]

	return true
}

func (it *ReverseIterator[T]) Next() bool {
	it.index = utils.Max(it.index-1, -1)

	if !it.IsValid() {
		return false
	}

	it.value = it.list.elements[it.index]

	return true
}

func (it *ReverseIterator[T]) NextN(n int) bool {
	it.index = utils.Max(it.index-n, -1)

	if !it.IsValid() {
		return false
	}

	it.value = it.list.elements[it.index]

	return true
}

func (it *ReverseIterator[T]) MoveBy(n int) bool {
	if n > 0 {
		return it.NextN(n)
	}

	return it.PreviousN(-n)
}

func (it *ReverseIterator[T]) Size() int {
	return it.size
}

func (it *ReverseIterator[T]) Index() (int, bool) {
	return it.index, it.IsValid()
}

func (it *ReverseIterator[T]) MoveTo(i int) bool {
	return it.MoveBy(it.index - i)
}

func (it *ReverseIterator[T]) IsBegin() bool {
	return it.size == 0 || it.index == it.size
}

func (it *ReverseIterator[T]) IsEnd() bool {
	return it.size == 0 || it.index == -1
}

func (it *ReverseIterator[T]) IsFirst() bool {
	return it.index == it.size-1
}

func (it *ReverseIterator[T]) IsLast() bool {
	return it.index == 0
}

func (it *ReverseIterator[T]) GetAt(i int) (value T, found bool) {
	if it.size == 0 || !it.list.withinRange(i) {
		return
	}

	return it.list.elements[i], true
}

func (it *ReverseIterator[T]) SetAt(i int, value T) bool {
	if it.size == 0 || !it.list.withinRange(i) {
		return false
	}
	it.list.elements[i] = value

	return true
}
