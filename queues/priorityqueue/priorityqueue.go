// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package priorityqueue implements a priority queue backed by binary queue.
//
// An unbounded priority queue based on a priority queue.
// The elements of the priority queue are ordered by a comparator provided at queue construction time.
//
// The heap of this queue is the least/smallest element with respect to the specified ordering.
// If multiple elements are tied for least value, the heap is one of those elements arbitrarily.
//
// Structure is not thread safe.
//
// References: https://en.wikipedia.org/wiki/Priority_queue
package priorityqueue

import (
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/queues"
	"github.com/JonasMuehlmann/datastructures.go/trees/binaryheap"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Queue implementation
var _ queues.Queue[any] = (*Queue[any])(nil)

// Queue holds elements in an array-queue
type Queue[T any] struct {
	heap       *binaryheap.Heap[T]
	Comparator utils.Comparator[T]
}

// New instantiates a new empty queue with the custom comparator.
func New[T any](comparator utils.Comparator[T], values ...T) *Queue[T] {
	queue := &Queue[T]{heap: binaryheap.New(comparator), Comparator: comparator}

	for _, value := range values {
		queue.Enqueue(value)
	}

	return queue
}

// NewFromSlice instantiates a new queue containing the provided slice.
func NewFromSlice[T any](comparator utils.Comparator[T], slice []T) *Queue[T] {
	queue := &Queue[T]{heap: binaryheap.New(comparator), Comparator: comparator}

	for _, value := range slice {
		queue.Enqueue(value)
	}

	return queue
}

// NewFromIterator instantiates a new queue containing the elements provided by the passed iterator.
func NewFromIterator[T any](comparator utils.Comparator[T], begin ds.ReadForIterator[T]) *Queue[T] {
	queue := &Queue[T]{heap: binaryheap.New(comparator), Comparator: comparator}

	for begin.Next() {
		newItem, _ := begin.Get()
		queue.Enqueue(newItem)
	}

	return queue
}

// NewFromIterators instantiates a new queue containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[T any](comparator utils.Comparator[T], begin ds.ReadCompForIterator[T], end ds.ComparableIterator) *Queue[T] {
	queue := &Queue[T]{heap: binaryheap.New(comparator), Comparator: comparator}

	for !begin.IsEqual(end) && begin.Next() {
		newItem, _ := begin.Get()
		queue.Enqueue(newItem)
	}

	return queue
}

// Enqueue adds a value to the end of the queue
func (queue *Queue[T]) Enqueue(value T) {
	queue.heap.Push(value)
}

// Dequeue removes first element of the queue and returns it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to dequeue.
func (queue *Queue[T]) Dequeue() (value T, ok bool) {
	return queue.heap.Pop()
}

// Peek returns top element on the queue without removing it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to peek.
func (queue *Queue[T]) Peek() (value T, ok bool) {
	return queue.heap.Peek()
}

// Empty returns true if queue does not contain any elements.
func (queue *Queue[T]) IsEmpty() bool {
	return queue.heap.IsEmpty()
}

// Size returns number of elements within the queue.
func (queue *Queue[T]) Size() int {
	return queue.heap.Size()
}

// Clear removes all elements from the queue.
func (queue *Queue[T]) Clear() {
	queue.heap.Clear()
}

// Values returns all elements in the queue.
func (queue *Queue[T]) GetValues() []T {
	return queue.heap.GetValues()
}

// String returns a string representation of container
func (queue *Queue[T]) ToString() string {
	str := "PriorityQueue\n"
	values := make([]string, queue.heap.Size(), queue.heap.Size())
	for index, value := range queue.heap.GetValues() {
		values[index] = fmt.Sprintf("%v", value)
	}
	str += strings.Join(values, ", ")
	return str
}

//******************************************************************//
//                             Iterator                             //
//******************************************************************//

// Begin returns an initialized iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (stack *Queue[T]) Begin() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return stack.NewOrderedIterator(-1, stack.Size())
}

// End returns an initialized iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (stack *Queue[T]) End() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return stack.NewOrderedIterator(stack.Size(), stack.Size())
}

// First returns an initialized iterator, which points to it's first element.
func (stack *Queue[T]) First() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return stack.NewOrderedIterator(0, stack.Size())
}

// Last returns an initialized iterator, which points to it's last element.
func (stack *Queue[T]) Last() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return stack.NewOrderedIterator(stack.Size()-1, stack.Size())
}
