// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package binaryheap implements a binary heap backed by array list.
//
// Comparator defines this heap as either min or max heap.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Binary_heap
package binaryheap

import (
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/lists/arraylist"
	"github.com/JonasMuehlmann/datastructures.go/trees"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Tree implementation
var _ trees.Tree[int, any] = (*Heap[any])(nil)

// Heap holds elements in an array-list
type Heap[T any] struct {
	list       *arraylist.List[T]
	Comparator utils.Comparator[T]
}

// NewWith instantiates a new empty heap tree with the custom comparator.
func New[T any](comparator utils.Comparator[T], values ...T) *Heap[T] {
	heap := &Heap[T]{list: arraylist.New[T](), Comparator: comparator}

	for _, element := range values {
		heap.Push(element)

	}

	return heap
}

// NewFromSlice instantiates a new stack containing the provided slice.
func NewFromSlice[T any](comparator utils.Comparator[T], slice []T) *Heap[T] {
	heap := &Heap[T]{list: arraylist.New[T](), Comparator: comparator}

	for _, element := range slice {
		heap.Push(element)

	}

	return heap
}

// NewFromIterator instantiates a new stack containing the elements provided by the passed iterator.
func NewFromIterator[T any](comparator utils.Comparator[T], begin ds.ReadCompForIterator[T]) *Heap[T] {
	heap := &Heap[T]{list: arraylist.New[T](), Comparator: comparator}

	for begin.Next() {
		newItem, _ := begin.Get()
		heap.Push(newItem)
	}

	return heap
}

// NewFromIterators instantiates a new stack containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[T any](comparator utils.Comparator[T], begin ds.ReadCompForIterator[T], end ds.ComparableIterator) *Heap[T] {
	heap := &Heap[T]{list: arraylist.New[T](), Comparator: comparator}

	for !begin.IsEqual(end) && begin.Next() {
		newItem, _ := begin.Get()
		heap.Push(newItem)
	}

	return heap
}

// Push adds a value onto the heap and bubbles it up accordingly.
func (heap *Heap[T]) Push(values ...T) {
	if len(values) == 1 {
		heap.list.PushBack(values[0])
		heap.bubbleUp()
	} else {
		// Reference: https://en.wikipedia.org/wiki/Binary_heap#Building_a_heap
		for _, value := range values {
			heap.list.PushBack(value)
		}
		size := heap.list.Size()/2 + 1
		for i := size; i >= 0; i-- {
			heap.bubbleDownIndex(i)
		}
	}
}

// Pop removes top element on heap and returns it, or nil if heap is empty.
// Second return parameter is true, unless the heap was empty and there was nothing to pop.
func (heap *Heap[T]) Pop() (value T, ok bool) {
	value, ok = heap.list.Get(0)
	if !ok {
		return
	}
	lastIndex := heap.list.Size() - 1
	heap.list.Swap(0, lastIndex)
	heap.list.Remove(lastIndex)
	heap.bubbleDown()
	return
}

// Peek returns top element on the heap without removing it, or nil if heap is empty.
// Second return parameter is true, unless the heap was empty and there was nothing to peek.
func (heap *Heap[T]) Peek() (value T, ok bool) {
	return heap.list.Get(0)
}

// Empty returns true if heap does not contain any elements.
func (heap *Heap[T]) IsEmpty() bool {
	return heap.list.IsEmpty()
}

// Size returns number of elements within the heap.
func (heap *Heap[T]) Size() int {
	return heap.list.Size()
}

// Clear removes all elements from the heap.
func (heap *Heap[T]) Clear() {
	heap.list.Clear()
}

// Values returns all elements in the heap.
func (heap *Heap[T]) GetValues() []T {
	values := make([]T, 0, heap.list.Size())

	it := heap.OrderedBegin()
	for it.Next() {
		value, _ := it.Get()

		values = append(values, value)
	}

	return values
}

// String returns a string representation of container
func (heap *Heap[T]) ToString() string {
	str := "BinaryHeap\n"
	values := []string{}

	it := heap.OrderedBegin()
	for it.Next() {
		value, _ := it.Get()
		values = append(values, fmt.Sprintf("%v", value))
	}

	str += strings.Join(values, ", ")
	return str
}

// Performs the "bubble down" operation. This is to place the element that is at the root
// of the heap in its correct place so that the heap maintains the min/max-heap order property.
func (heap *Heap[T]) bubbleDown() {
	heap.bubbleDownIndex(0)
}

// Performs the "bubble down" operation. This is to place the element that is at the index
// of the heap in its correct place so that the heap maintains the min/max-heap order property.
func (heap *Heap[T]) bubbleDownIndex(index int) {
	size := heap.list.Size()
	for leftIndex := index<<1 + 1; leftIndex < size; leftIndex = index<<1 + 1 {
		rightIndex := index<<1 + 2
		smallerIndex := leftIndex
		leftValue, _ := heap.list.Get(leftIndex)
		rightValue, _ := heap.list.Get(rightIndex)
		if rightIndex < size && heap.Comparator(leftValue, rightValue) > 0 {
			smallerIndex = rightIndex
		}
		indexValue, _ := heap.list.Get(index)
		smallerValue, _ := heap.list.Get(smallerIndex)
		if heap.Comparator(indexValue, smallerValue) > 0 {
			heap.list.Swap(index, smallerIndex)
		} else {
			break
		}
		index = smallerIndex
	}
}

// Performs the "bubble up" operation. This is to place a newly inserted
// element (i.e. last element in the list) in its correct place so that
// the heap maintains the min/max-heap order property.
func (heap *Heap[T]) bubbleUp() {
	index := heap.list.Size() - 1
	for parentIndex := (index - 1) >> 1; index > 0; parentIndex = (index - 1) >> 1 {
		indexValue, _ := heap.list.Get(index)
		parentValue, _ := heap.list.Get(parentIndex)
		if heap.Comparator(parentValue, indexValue) <= 0 {
			break
		}
		heap.list.Swap(index, parentIndex)
		index = parentIndex
	}
}

// Check that the index is within bounds of the list
func (heap *Heap[T]) withinRange(index int) bool {
	return index >= 0 && index < heap.list.Size()
}

//******************************************************************//
//                         OrderedIterator                         //
//******************************************************************//

// Begin returns an initialized iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (heap *Heap[T]) OrderedBegin() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return heap.NewOrderedIterator(-1, heap.Size())
}

// End returns an initialized iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (heap *Heap[T]) OrderedEnd() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return heap.NewOrderedIterator(heap.list.Size(), heap.Size())
}

// First returns an initialized iterator, which points to it's first element.
func (heap *Heap[T]) OrderedFirst() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return heap.NewOrderedIterator(0, heap.Size())
}

// Last returns an initialized iterator, which points to it's last element.
func (heap *Heap[T]) OrderedLast() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return heap.NewOrderedIterator(heap.list.Size()-1, heap.Size())
}
