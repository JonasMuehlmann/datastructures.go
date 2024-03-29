// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package singlylinkedlist

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompForRandCollIterator[int, any] = (*Iterator[any])(nil)

// Iterator holding the iterator's state
type Iterator[T any] struct {
	list    *List[T]
	index   int
	element *element[T]
	// Redundant but stored for better locality
	size int
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
func (list *List[T]) NewIterator(position int, size int) *Iterator[T] {
	it := &Iterator[T]{list: list, index: position, size: size}
	it.size = utils.Min(list.Size(), size)
	it.size = utils.Max(list.Size(), -1)

	it.element = list.first
	it.MoveTo(position)

	return it
}

func (it *Iterator[T]) IsValid() bool {
	return it.list.size > 0 && !it.IsBegin() && !it.IsEnd()
}

func (it *Iterator[T]) Get() (value T, found bool) {
	if !it.IsValid() {
		return
	}

	return it.element.value, true
}

func (it *Iterator[T]) Set(value T) bool {
	if !it.IsValid() {
		return false
	}

	it.element.value = value

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

	if !it.IsFirst() {
		it.element = it.element.next
	}

	return true
}

func (it *Iterator[T]) NextN(n int) bool {
	if n < 0 {
		return false
	}

	n = utils.Min(it.index+n, it.size) - it.index
	it.index += n

	if !it.IsValid() {
		return false
	}

	if it.IsLast() {
		it.element = it.list.last

		return true
	}

	for i := 0; i < n; i++ {
		it.element = it.element.next
	}

	return true
}

func (it *Iterator[T]) MoveTo(n int) bool {
	if n < 0 {
		return false
	}

	return it.NextN(n - it.index)
}

func (it *Iterator[T]) MoveToKey(n int) bool {
	return it.MoveTo(n)
}

func (it *Iterator[T]) Size() int {
	return it.list.size
}

func (it *Iterator[T]) Index() (int, bool) {
	return it.index, true
}
func (it *Iterator[T]) GetKey() (int, bool) {
	return it.Index()
}

func (it *Iterator[T]) IsBegin() bool {
	return it.index == -1
}

func (it *Iterator[T]) IsEnd() bool {
	return it.list.size == 0 || it.index == it.list.size
}

func (it *Iterator[T]) IsFirst() bool {
	return it.index == 0
}

func (it *Iterator[T]) IsLast() bool {
	return it.index == it.list.size-1
}
