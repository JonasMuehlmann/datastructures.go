// Copyright (c) 2021, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package circularbuffer implements the circular buffer.
//
// In computer science, a circular buffer, circular queue, cyclic buffer or ring buffer is a data structure that uses a single, fixed-size buffer as if it were connected end-to-end. This structure lends itself easily to buffering data streams.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Circular_buffer
package circularbuffer

import (
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/queues"
)

// Assert Queue implementation
var _ queues.Queue[any] = (*Queue[any])(nil)

// Queue holds values in a slice.
type Queue[T any] struct {
	values  []T
	start   int
	end     int
	full    bool
	maxSize int
	size    int
}

// New instantiates a new empty queue with the specified size of maximum number of elements that it can hold.
// This max size of the buffer cannot be changed.
func New[T any](maxSize int) *Queue[T] {
	if maxSize < 1 {
		panic("Invalid maxSize, should be at least 1")
	}

	queue := &Queue[T]{maxSize: maxSize, values: make([]T, 0, maxSize)}

	return queue
}

func NewFromSlice[T any](maxSize int, slice []T) *Queue[T] {
	if maxSize < 1 {
		panic("Invalid maxSize, should be at least 1")
	}

	list := &Queue[T]{
		maxSize: maxSize,
		values:  append(make([]T, 0, maxSize), slice...),
		size:    len(slice),
		end:     len(slice),
	}

	return list
}

// NewFromIterator instantiates a new queue containing the elements provided by the passed iterator.
func NewFromIterator[T any](maxSize int, it ds.ReadCompForIterator[T]) *Queue[T] {
	if maxSize < 1 {
		panic("Invalid maxSize, should be at least 1")
	}

	length := 0
	sizedIterator, ok := it.(ds.SizedIterator)
	if ok {
		length = sizedIterator.Size()
	}

	elements := make([]T, 0, length)

	for ; !it.IsEnd(); it.Next() {
		newItem, _ := it.Get()
		elements = append(elements, newItem)
	}

	queue := &Queue[T]{
		values:  elements,
		end:     len(elements),
		size:    len(elements),
		maxSize: maxSize,
	}

	return queue
}

// NewFromIterators instantiates a new queue containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[T any](maxSize int, first ds.ReadCompForIterator[T], end ds.ComparableIterator) *Queue[T] {
	if maxSize < 1 {
		panic("Invalid maxSize, should be at least 1")
	}

	length := 0
	sizedFirst, ok := first.(ds.OrderedIterator)
	sizedLast, ok2 := end.(ds.OrderedIterator)
	if ok && ok2 {
		length = -sizedFirst.DistanceTo(sizedLast)
		if length < 0 {
			length = 0
		}
	}

	elements := make([]T, 0, maxSize)

	for ; !first.IsEqual(end); first.Next() {
		newItem, _ := first.Get()
		elements = append(elements, newItem)
	}

	queue := &Queue[T]{
		values:  elements,
		end:     len(elements),
		size:    len(elements),
		maxSize: maxSize,
	}

	return queue
}

// Enqueue adds a value to the end of the queue
func (queue *Queue[T]) Enqueue(value T) {
	if queue.Full() {
		queue.Dequeue()
	}
	queue.values[queue.end] = value
	queue.end = queue.end + 1
	if queue.end >= queue.maxSize {
		queue.end = 0
	}
	if queue.end == queue.start {
		queue.full = true
	}

	queue.size = queue.calculateSize()
}

// Dequeue removes first element of the queue and returns it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to dequeue.
func (queue *Queue[T]) Dequeue() (value T, ok bool) {
	if queue.IsEmpty() {
		return
	}

	value, ok = queue.values[queue.start], true

	queue.start = queue.start + 1
	if queue.start >= queue.maxSize {
		queue.start = 0
	}
	queue.full = false

	queue.size = queue.size - 1

	return
}

// Peek returns first element of the queue without removing it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to peek.
func (queue *Queue[T]) Peek() (value T, ok bool) {
	if queue.IsEmpty() {
		return
	}
	return queue.values[queue.start], true
}

// Empty returns true if queue does not contain any elements.
func (queue *Queue[T]) IsEmpty() bool {
	return queue.Size() == 0
}

// Full returns true if the queue is full, i.e. has reached the maximum number of elements that it can hold.
func (queue *Queue[T]) Full() bool {
	return queue.Size() == queue.maxSize
}

// Size returns number of elements within the queue.
func (queue *Queue[T]) Size() int {
	return queue.size
}

// Clear removes all elements from the queue.
func (queue *Queue[T]) Clear() {
	queue.values = make([]T, queue.maxSize, queue.maxSize)
	queue.start = 0
	queue.end = 0
	queue.full = false
	queue.size = 0
}

// Values returns all elements in the queue (FIFO order).
func (queue *Queue[T]) GetValues() []T {
	values := make([]T, queue.Size(), queue.Size())
	for i := 0; i < queue.Size(); i++ {
		values[i] = queue.values[(queue.start+i)%queue.maxSize]
	}
	return values
}

// String returns a string representation of container
func (queue *Queue[T]) ToString() string {
	str := "CircularBuffer\n"
	var values []string
	for _, value := range queue.GetValues() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (queue *Queue[T]) withinRange(index int) bool {
	return index >= 0 && index < queue.size
}

func (queue *Queue[T]) calculateSize() int {
	if queue.end < queue.start {
		return queue.maxSize - queue.start + queue.end
	} else if queue.end == queue.start {
		if queue.full {
			return queue.maxSize
		}
		return 0
	}

	return queue.end - queue.start
}

//******************************************************************//
//                             Iterator                             //
//******************************************************************//

// Begin returns an initialized iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (queue *Queue[T]) Begin() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return queue.NewIterator(queue, queue.start-1)
}

// End returns an initialized iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (queue *Queue[T]) End() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return queue.NewIterator(queue, queue.end)
}

// First returns an initialized iterator, which points to it's first element.
func (queue *Queue[T]) First() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return queue.NewIterator(queue, queue.start)
}

// Last returns an initialized iterator, which points to it's last element.
func (queue *Queue[T]) Last() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return queue.NewIterator(queue, queue.end-1)
}
