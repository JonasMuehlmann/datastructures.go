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
var _ ds.ReadWriteOrdCompBidRandCollIterator[int, any] = (*Iterator[any])(nil)

// Iterator holding the iterator's state
type Iterator[T any] struct {
	list  *List[T]
	index int
	// Redundant but has better locality
	value T
	size  int
}

// NewIterator returns a stateful iterator whose values can be fetched by an index.
func (list *List[T]) NewIterator(index int, size int) *Iterator[T] {
	it := &Iterator[T]{list: list, index: index, size: size}
	it.size = utils.Min(list.Size(), size)
	it.size = utils.Max(list.Size(), -1)

	if it.IsValid() {
		it.value = list.elements[it.index]
	}

	return it
}

func (it *Iterator[T]) IsValid() bool {
	return it.size > 0 && !it.IsBegin() && !it.IsEnd()
}

func (it *Iterator[T]) Get() (value T, found bool) {
	if !it.IsValid() {
		return
	}

	return it.value, true
}

func (it *Iterator[T]) Set(value T) bool {
	if !it.IsValid() {
		return false
	}
	it.list.elements[it.index] = value
	it.value = value

	return true
}

// If other is of type IndexedIterator, IndexedIterator.Index() will be used, possibly executing in O(1)
func (it *Iterator[T]) DistanceTo(other ds.OrderedIterator) int {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.index - otherThis.index
}

func (it *Iterator[T]) IsAfter(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) > 0
}

func (it *Iterator[T]) IsBefore(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) < 0
}

func (it *Iterator[T]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) == 0
}

func (it *Iterator[T]) Next() bool {
	it.index = utils.Min(it.index+1, it.size)

	if !it.IsValid() {
		return false
	}

	it.value = it.list.elements[it.index]

	return true
}

func (it *Iterator[T]) NextN(i int) bool {
	it.index = utils.Min(it.index+i, it.size)

	if !it.IsValid() {
		return false
	}

	it.value = it.list.elements[it.index]

	return true
}

func (it *Iterator[T]) Previous() bool {
	it.index = utils.Max(it.index-1, -1)

	if !it.IsValid() {
		return false
	}

	it.value = it.list.elements[it.index]

	return true
}

func (it *Iterator[T]) PreviousN(n int) bool {
	it.index = utils.Max(it.index-n, -1)

	if !it.IsValid() {
		return false
	}

	it.value = it.list.elements[it.index]

	return true
}

func (it *Iterator[T]) MoveBy(n int) bool {
	if n > 0 {
		return it.NextN(n)
	} else if n < 0 {
		return it.PreviousN(-n)
	}

	return it.IsValid()
}

func (it *Iterator[T]) Size() int {
	return it.size
}

func (it *Iterator[T]) Index() (int, bool) {
	return it.index, it.IsValid()
}

func (it *Iterator[T]) GetKey() (int, bool) {
	return it.Index()
}

func (it *Iterator[T]) MoveTo(i int) bool {
	return it.MoveBy(i - it.index)
}

func (it *Iterator[T]) MoveToKey(i int) bool {
	return it.MoveTo(i)
}

func (it *Iterator[T]) IsBegin() bool {
	return it.index == -1
}

func (it *Iterator[T]) IsEnd() bool {
	return it.size == 0 || it.index == it.size
}

func (it *Iterator[T]) IsFirst() bool {
	return it.index == 0
}

func (it *Iterator[T]) IsLast() bool {
	return it.index == it.size-1
}

func (it *Iterator[T]) GetAt(i int) (value T, found bool) {
	if it.size == 0 || !it.list.withinRange(i) {
		return
	}

	return it.list.elements[i], true
}

func (it *Iterator[T]) SetAt(i int, value T) bool {
	if it.size == 0 || !it.list.withinRange(i) {
		return false
	}
	it.list.elements[i] = value

	return true
}

func (it *Iterator[T]) GetAtKey(i int) (value T, found bool) {
	return it.GetAt(i)
}

func (it *Iterator[T]) SetAtKey(i int, value T) bool {
	return it.SetAt(i, value)
}
