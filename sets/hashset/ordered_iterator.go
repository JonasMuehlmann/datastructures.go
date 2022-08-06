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
	// Redundant but has better locality
	value T
	size  int
}

func (m *Set[T]) NewOrderedIterator(s *Set[T], position int, comparator utils.Comparator[T]) *OrderedIterator[T] {
	keys := s.GetValues()
	utils.Sort(keys, comparator)

	it := &OrderedIterator[T]{
		s:          s,
		values:     keys,
		index:      position,
		comparator: comparator,
		size:       s.Size(),
	}

	if it.IsValid() {
		it.value = it.values[it.index]
	}

	return it
}


func (it *OrderedIterator[T]) IsBegin() bool {
	return it.index == -1
}


func (it *OrderedIterator[T]) IsEnd() bool {
	return it.size == 0 || it.index == it.size
}


func (it *OrderedIterator[T]) IsFirst() bool {
	return it.index == 0
}


func (it *OrderedIterator[T]) IsLast() bool {
	return it.index == it.size-1
}


func (it *OrderedIterator[T]) IsValid() bool {
	return it.size > 0 && !it.IsBegin() && !it.IsEnd()
}


func (it *OrderedIterator[T]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) == 0
}


func (it *OrderedIterator[T]) DistanceTo(other ds.OrderedIterator) int {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.index - otherThis.index
}


func (it *OrderedIterator[T]) IsAfter(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) > 0
}


func (it *OrderedIterator[T]) IsBefore(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) < 0
}


func (it *OrderedIterator[T]) Size() int {
	return it.size
}


func (it *OrderedIterator[T]) Index() (index int, found bool) {
	if !it.IsValid() {
		found = false

		return
	}

	index = it.index
	found = true

	return
}


func (it *OrderedIterator[T]) Next() bool {
	it.index = utils.Min(it.index+1, it.size)

	if !it.IsValid() {
		return false
	}

	it.value = it.values[it.index]

	return true

}


func (it *OrderedIterator[T]) NextN(i int) bool {
	it.index = utils.Min(it.index+i, it.size)

	if !it.IsValid() {
		return false
	}

	it.value = it.values[it.index]

	return true

}


func (it *OrderedIterator[T]) Previous() bool {
	it.index = utils.Max(it.index-1, -1)

	if !it.IsValid() {
		return false
	}

	it.value = it.values[it.index]

	return true
}


func (it *OrderedIterator[T]) PreviousN(n int) bool {
	it.index = utils.Max(it.index-n, -1)

	if !it.IsValid() {
		return false
	}

	it.value = it.values[it.index]

	return true
}


func (it *OrderedIterator[T]) MoveBy(n int) bool {
	if n > 0 {
		return it.NextN(n)
	}

	return it.PreviousN(-n)
}


func (it *OrderedIterator[T]) MoveTo(i int) bool {
	return it.MoveBy(i - it.index)
}


func (it *OrderedIterator[T]) Get() (value T, found bool) {
	if !it.IsValid() {
		return
	}

	return it.value, true
}


func (it *OrderedIterator[T]) GetAt(i int) (value T, found bool) {
	if i < 0 || i >= it.size {
		return
	}

	return it.values[i], true
}


func (it *OrderedIterator[T]) Set(value T) bool {
	if !it.IsValid() {
		return false
	}

	it.values[it.index] = value
	it.value = value

	return true
}


func (it *OrderedIterator[T]) SetAt(i int, value T) bool {
	if i < 0 || i >= it.size {
		return false
	}

	delete(it.s.items, it.values[i])
	it.s.items[value] = itemExists

	it.values[i] = value

	return true
}
