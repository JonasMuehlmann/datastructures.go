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
	// Redundant but has better locality
	value T
	index int
	size  int
}

func (s *Set[T]) NewIterator(position int, size int, comparator utils.Comparator[T]) *Iterator[T] {
	it := &Iterator[T]{
		s:             s,
		orderIterator: s.ordering.NewIterator(position, size),
		comparator:    comparator,
		index:         position,
		size:          size,
	}
	it.size = utils.Min(s.Size(), size)
	it.size = utils.Max(s.Size(), -1)

	it.value, _ = it.orderIterator.Get()

	return it
}

func (it *Iterator[T]) IsBegin() bool {
	return it.index == -1
}

func (it *Iterator[T]) IsEnd() bool {
	return it.index == it.size
}

func (it *Iterator[T]) IsFirst() bool {
	return it.index == 0
}

func (it *Iterator[T]) IsLast() bool {
	return it.index == it.size-1
}

func (it *Iterator[T]) IsValid() bool {
	return it.size > 0 && !it.IsBegin() && !it.IsEnd()
}

func (it *Iterator[T]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.index == otherThis.index

}

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

func (it *Iterator[T]) Size() int {
	return it.size
}

func (it *Iterator[T]) Index() (index int, found bool) {
	return it.index, it.IsValid()
}
func (it *Iterator[T]) GetKey() (index int, found bool) {
	return it.Index()
}

func (it *Iterator[T]) Next() bool {
	it.orderIterator.Next()
	it.index = utils.Min(it.index+1, it.size)

	var found bool

	it.value, found = it.orderIterator.Get()

	return found
}

func (it *Iterator[T]) NextN(i int) bool {
	it.orderIterator.NextN(i)
	it.index = utils.Min(it.index+i, it.size)

	var found bool

	it.value, found = it.orderIterator.Get()

	return found
}

func (it *Iterator[T]) Previous() bool {
	it.orderIterator.Previous()
	it.index = utils.Max(it.index-1, -1)

	var found bool

	it.value, found = it.orderIterator.Get()

	return found
}

func (it *Iterator[T]) PreviousN(n int) bool {
	it.orderIterator.PreviousN(n)
	it.index = utils.Max(it.index-n, -1)

	var found bool

	it.value, found = it.orderIterator.Get()

	return found
}

func (it *Iterator[T]) MoveBy(n int) bool {
	if n > 0 {
		return it.NextN(n)
	}

	return it.PreviousN(-n)
}

func (it *Iterator[T]) MoveTo(i int) bool {
	return it.MoveBy(i - it.index)
}

func (it *Iterator[T]) MoveToKey(i int) bool {
	return it.MoveTo(i)
}

func (it *Iterator[T]) Get() (value T, found bool) {
	return it.value, it.IsValid()
}

func (it *Iterator[T]) GetAt(i int) (value T, found bool) {
	return it.orderIterator.GetAt(i)
}

func (it *Iterator[T]) GetAtKey(i int) (value T, found bool) {
	return it.GetAt(i)
}

func (it *Iterator[T]) Set(value T) bool {
	if !it.IsValid() {
		return false
	}

	it.orderIterator.Set(value)

	valueToRemove, _ := it.Get()
	it.s.Remove(it.comparator, valueToRemove)
	it.s.table[value] = itemExists
	it.value = value

	return true
}

func (it *Iterator[T]) SetAt(i int, value T) bool {
	if i < 0 || i >= it.Size() {
		return false
	}

	valueToRemove, _ := it.Get()
	delete(it.s.table, valueToRemove)
	it.s.table[value] = itemExists

	return true
}

func (it *Iterator[T]) SetAtKey(i int, value T) bool {
	return it.SetAt(i, value)
}
