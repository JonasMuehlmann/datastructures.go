// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doublylinkedlist

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[int, any] = (*Iterator[any])(nil)

// Iterator holding the iterator's state
type Iterator[T any] struct {
	list    *List[T]
	index   int
	element *element[T]
	// Redundant but stored for better locality
	size int
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
func (list *List[T]) NewIterator(position int) *Iterator[T] {
	it := &Iterator[T]{list: list, index: position, size: list.Size()}

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

func (it *Iterator[T]) GetAt(i int) (value T, found bool) {
	if it.list.size == 0 || !it.list.withinRange(i) {
		return
	}

	oldIndex := it.index

	it.MoveTo(i)

	value = it.element.value
	found = true

	it.MoveTo(oldIndex)

	return
}

func (it *Iterator[T]) SetAt(i int, value T) bool {
	if it.list.size == 0 || !it.list.withinRange(i) {
		return false
	}

	oldIndex := it.index

	it.MoveTo(i)

	it.element.value = value

	it.MoveTo(oldIndex)

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

	if it.IsFirst() {
		it.element = it.list.first
	} else {
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

	if it.index-n == -1 {
		n -= 1
	}

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

func (it *Iterator[T]) Previous() bool {
	it.index = utils.Max(it.index-1, -1)

	if !it.IsValid() {
		return false
	}

	if it.IsLast() {
		it.element = it.list.last
	} else {
		it.element = it.element.prev
	}

	return true
}

func (it *Iterator[T]) PreviousN(n int) bool {
	if n < 0 {
		return false
	}

	n = it.index - utils.Max(it.index-n, -1)
	it.index -= n

	if !it.IsValid() {
		return false
	}

	if it.IsFirst() {
		it.element = it.list.first

		return true
	}

	for i := 0; i < n; i++ {
		it.element = it.element.prev
	}

	return true
}

func (it *Iterator[T]) MoveTo(n int) bool {
	return it.MoveBy(n - it.index)

}

func (it *Iterator[T]) MoveBy(n int) bool {
	if n < 0 {
		return it.PreviousN(-n)
	} else {
		return it.NextN(n)
	}
}

func (it *Iterator[T]) Size() int {
	return it.list.size
}

func (it *Iterator[T]) Index() (int, bool) {
	return it.index, true
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
