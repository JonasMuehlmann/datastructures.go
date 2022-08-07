// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binaryheap

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert OrderedIterator implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[int, any] = (*OrderedIterator[any])(nil)

type visitedChildren int

// OrderedIterator holding the iterator's state
type OrderedIterator[T any] struct {
	heap  *Heap[T]
	index int
	// heapIndex        int
	valueDirty bool
	// visitedChildren visitedChildren
	// Redundant but has better locality
	value T
	size  int
}

// NewOrderedIterator returns a stateful iterator whose values can be fetched by an index.
func (list *Heap[T]) NewOrderedIterator(index int, size int) *OrderedIterator[T] {
	it := &OrderedIterator[T]{heap: list, index: index, size: size, valueDirty: true}
	it.size = utils.Min(list.Size(), size)
	it.size = utils.Max(list.Size(), -1)

	return it
}

func (it *OrderedIterator[T]) IsValid() bool {
	return it.size > 0 && !it.IsBegin() && !it.IsEnd()
}

func (it *OrderedIterator[T]) Get() (value T, found bool) {
	if !it.IsValid() {
		return
	}

	if it.valueDirty {
		start, end := evaluateRange(it.index)

		if end > it.heap.Size() {
			end = it.heap.Size()
		}

		tmpHeap := New(it.heap.Comparator)

		for n := start; n < end; n++ {
			value, _ := it.heap.list.Get(n)
			tmpHeap.Push(value)
		}

		for n := 0; n < it.index-start; n++ {
			tmpHeap.Pop()
		}

		value, _ := tmpHeap.Pop()

		it.value = value
		it.valueDirty = false
	}

	return it.value, true
}

func (it *OrderedIterator[T]) Set(value T) bool {
	if !it.IsValid() {
		return false
	}

	// PERF: This is absolute madness
	values := it.heap.list.GetValues()
	atNthLargest := 0
	nthLargest, _ := it.heap.Peek()

	for _, value := range values {
		if atNthLargest == it.index {
			break
		}
		if it.heap.Comparator(nthLargest, value) < 0 {
			nthLargest = value
			atNthLargest++
		}
	}

	for i, value := range it.heap.list.GetSlice() {
		if it.heap.Comparator(nthLargest, value) == 0 {
			it.heap.list.Set(i, value)
			it.heap.bubbleDownIndex(i)
			break
		}
	}

	return true
}

// If other is of type IndexedOrderedIterator, IndexedOrderedIterator.Index() will be used, possibly executing in O(1)
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

func (it *OrderedIterator[T]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) == 0
}

func (it *OrderedIterator[T]) Next() bool {
	it.index = utils.Min(it.index+1, it.size)
	it.valueDirty = true

	return it.IsValid()
}

func (it *OrderedIterator[T]) NextN(n int) bool {
	it.index = utils.Min(it.index+n, it.size)
	it.valueDirty = true

	return it.IsValid()
}

func (it *OrderedIterator[T]) Previous() bool {
	it.index = utils.Max(it.index-1, -1)
	it.valueDirty = true

	return it.IsValid()
}

func (it *OrderedIterator[T]) PreviousN(n int) bool {
	it.index = utils.Max(it.index-n, -1)
	it.valueDirty = true

	return it.IsValid()
}

func (it *OrderedIterator[T]) MoveBy(n int) bool {
	if n > 0 {
		return it.NextN(n)
	}

	return it.PreviousN(-n)
}

func (it *OrderedIterator[T]) Size() int {
	return it.size
}

func (it *OrderedIterator[T]) Index() (int, bool) {
	return it.index, it.IsValid()
}

func (it *OrderedIterator[T]) MoveTo(i int) bool {
	return it.MoveBy(i - it.index)
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

func (it *OrderedIterator[T]) GetAt(i int) (value T, found bool) {
	if it.size == 0 || !it.heap.withinRange(i) {
		return
	}

	if it.index == i {
		return it.Get()
	}

	tmp := *it
	tmp.MoveTo(i)

	return tmp.Get()
}

func (it *OrderedIterator[T]) SetAt(i int, value T) bool {
	if it.size == 0 || !it.heap.withinRange(i) || i >= it.size {
		return false
	}

	if it.index == i {
		return it.Set(value)
	}

	tmp := *it
	tmp.MoveTo(i)
	it.heap.list.Set(tmp.index, value)

	it.heap.bubbleDownIndex(tmp.index)

	return true
}

// numOfBits counts the number of bits of an int
func numOfBits(n int) uint {
	var count uint
	for n != 0 {
		count++
		n >>= 1
	}
	return count
}

// evaluateRange evaluates the index range [start,end) of same level nodes in the heap as the index
func evaluateRange(index int) (start int, end int) {
	bits := numOfBits(index+1) - 1
	start = 1<<bits - 1
	end = start + 1<<bits
	return
}
