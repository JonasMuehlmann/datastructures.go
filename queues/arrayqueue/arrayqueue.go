// Copyright (c) 2021, Aryan Ahadinia. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package arrayqueue implements a queue backed by array list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Queue_(abstract_data_type)
package arrayqueue

import (
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/lists/arraylist"
	"github.com/JonasMuehlmann/datastructures.go/queues"
)

// Assert Queue implementation
var _ queues.Queue[any] = (*Queue[any])(nil)

// Queue holds elements in an array-list
type Queue[T any] struct {
	list *arraylist.List[T]
}

// New instantiates a new empty queue
func New[T any](items ...T) *Queue[T] {
	return &Queue[T]{list: arraylist.NewFromSlice[T](items)}
}

func NewFromSlice[T any](slice []T) *Queue[T] {
	list := &Queue[T]{list: arraylist.NewFromSlice(slice)}
	return list
}

// NewFromIterator instantiates a new queue containing the elements provided by the passed iterator.
func NewFromIterator[T any](begin ds.ReadForIterator[T]) *Queue[T] {
	length := 0
	sizedIterator, ok := begin.(ds.SizedIterator)
	if ok {
		length = sizedIterator.Size()
	}

	elements := make([]T, 0, length)

	for begin.Next() {
		newItem, _ := begin.Get()
		elements = append(elements, newItem)
	}

	queue := &Queue[T]{list: arraylist.NewFromSlice(elements)}

	return queue
}

// NewFromIterators instantiates a new queue containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[T any](begin ds.ReadCompForIterator[T], end ds.ComparableIterator) *Queue[T] {
	length := 0
	sizedFirst, ok := begin.(ds.OrderedIterator)
	sizedLast, ok2 := end.(ds.OrderedIterator)
	if ok && ok2 {
		length = -sizedFirst.DistanceTo(sizedLast)
		if length < 0 {
			length = 0
		}
	}

	elements := make([]T, 0, length)

	for !begin.IsEqual(end) && begin.Next() {
		newItem, _ := begin.Get()
		elements = append(elements, newItem)
	}

	queue := &Queue[T]{list: arraylist.NewFromSlice(elements)}

	return queue
}

// Enqueue adds a value to the end of the queue
func (queue *Queue[T]) Enqueue(value T) {
	queue.list.PushBack(value)
}

// Dequeue removes first element of the queue and returns it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to dequeue.
func (queue *Queue[T]) Dequeue() (value T, ok bool) {
	value, ok = queue.list.Get(0)
	if ok {
		queue.list.PopFront(1)
	}
	return
}

// Peek returns first element of the queue without removing it, or nil if queue is empty.
// Second return parameter is true, unless the queue was empty and there was nothing to peek.
func (queue *Queue[T]) Peek() (value T, ok bool) {
	return queue.list.Get(0)
}

// Empty returns true if queue does not contain any elements.
func (queue *Queue[T]) IsEmpty() bool {
	return queue.list.IsEmpty()
}

// Size returns number of elements within the queue.
func (queue *Queue[T]) Size() int {
	return queue.list.Size()
}

// Clear removes all elements from the queue.
func (queue *Queue[T]) Clear() {
	queue.list.Clear()
}

// Values returns all elements in the queue (FIFO order).
func (queue *Queue[T]) GetValues() []T {
	return queue.list.GetValues()
}

// String returns a string representation of container
func (queue *Queue[T]) ToString() string {
	str := "ArrayQueue\n"
	values := []string{}
	for _, value := range queue.list.GetValues() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (queue *Queue[T]) withinRange(index int) bool {
	return index >= 0 && index < queue.list.Size()
}

//******************************************************************//
//                             Iterator                             //
//******************************************************************//

// Begin returns an initialized iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (queue *Queue[T]) Begin() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return queue.NewIterator(-1, queue.Size())
}

// End returns an initialized iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (queue *Queue[T]) End() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return queue.NewIterator(queue.list.Size(), queue.Size())
}

// First returns an initialized iterator, which points to it's first element.
func (queue *Queue[T]) First() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return queue.NewIterator(0, queue.Size())
}

// Last returns an initialized iterator, which points to it's last element.
func (queue *Queue[T]) Last() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return queue.NewIterator(queue.list.Size()-1, queue.Size())
}
