// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doublylinkedlist

import "github.com/JonasMuehlmann/datastructures.go/ds"

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[int, *element[any]] = (*Iterator[any])(nil)

// Iterator holding the iterator's state
type Iterator[T any] struct {
	list    *List[T]
	index   int
	element *element[T]
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
func (list *List[T]) NewIterator(l *List[T], position int) *Iterator[T] {
	it := &Iterator[T]{list: l}

	it.element = l.first
	it.NextN(position)

	return it
}

// IsValid implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsValid() bool {
	return it.list.withinRange(it.index)
}

// Get implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Get() (value *element[T], found bool) {
	if it.list.size == 0 || !it.IsValid() {
		return
	}

	value = it.element

	if value != nil {
		found = true
	}

	return
}

// Set implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Set(value *element[T]) bool {
	if it.list.size == 0 || !it.IsValid() {
		return false
	}

	it.element = value

	return true
}

// Get implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) GetAt(i int) (value *element[T], found bool) {
	if it.list.size == 0 || !it.list.withinRange(i) {
		return
	}

	oldIndex := it.index

	it.MoveTo(i)

	value = it.element

	if value != nil {
		found = true
	}

	it.MoveTo(oldIndex)

	return
}

// Set implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) SetAt(i int, value *element[T]) bool {
	if it.list.size == 0 || !it.list.withinRange(i) {
		return false
	}

	oldIndex := it.index

	it.MoveTo(i)

	it.element = value

	it.MoveTo(oldIndex)

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

	if it.index >= it.list.size || it.list.size == 0 {
		it.element = nil

		return
	}

	it.element = it.element.next
}

// NextN implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) NextN(n int) {
	if it.index+n > it.list.size || n < 0 {
		it.index += n
		it.element = nil

		return
	}

	if it.index+n == it.list.size-1 {
		it.element = it.list.last
		it.index += n

		return
	}

	for i := 0; i < n; i++ {
		it.element = it.element.next
	}

	it.index += n
}

// Next implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Previous() {
	it.index--

	if it.index < 0 || it.list.size == 0 {
		it.element = nil

		return
	}

	if it.index != 0 {
		it.element = it.element.prev
	}

}

// NextN implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) PreviousN(n int) {
	if it.index-n < 0 || it.list.size == 0 || n < 0 {
		it.index -= n
		it.element = nil

		return
	}

	if it.index-n == 0 {
		it.element = it.list.first
		it.index -= n

		return
	}

	for i := 0; i < n; i++ {
		it.element = it.element.prev
	}

	it.index -= n
}

// MoveTo implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) MoveTo(n int) bool {
	if it.list.size == 0 || n < 0 {
		it.index = n
		it.element = nil

		return true

	}

	it.MoveBy(n - it.index)

	return true
}

// NextN implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) MoveBy(n int) {
	if n < 0 {
		it.PreviousN(-n)
	} else {
		it.NextN(n)
	}
}

// Size implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Size() int {
	return it.list.size
}

// Index implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) Index() (int, bool) {
	return it.index, true
}

// IsBegin implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsBegin() bool {
	return it.index == -1
}

// IsEnd implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsEnd() bool {
	return it.list.size == 0 || it.index == it.list.size
}

// IsFirst implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsFirst() bool {
	return it.index == 0
}

// IsLast implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsLast() bool {
	return it.index == it.list.size-1
}
