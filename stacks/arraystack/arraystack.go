// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package arraystack implements a stack backed by array list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29#Array
package arraystack

import (
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/lists/arraylist"
	"github.com/JonasMuehlmann/datastructures.go/stacks"
)

// Assert Stack implementation
var _ stacks.Stack[any] = (*Stack[any])(nil)

// Stack holds elements in an array-list
type Stack[T any] struct {
	list *arraylist.List[T]
}

// New instantiates a new empty stack
func New[T any](values ...T) *Stack[T] {
	return &Stack[T]{list: arraylist.New[T](values...)}
}

// NewFromSlice instantiates a new stack containing the provided slice.
func NewFromSlice[T any](slice []T) *Stack[T] {
	list := &Stack[T]{list: arraylist.NewFromSlice(slice)}
	return list
}

// NewFromIterator instantiates a new stack containing the elements provided by the passed iterator.
func NewFromIterator[T any](it ds.ReadCompForIterator[T]) *Stack[T] {
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

	list := &Stack[T]{list: arraylist.NewFromSlice(elements)}

	return list
}

// NewFromIterators instantiates a new stack containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[T any](first ds.ReadCompForIterator[T], end ds.ComparableIterator) *Stack[T] {
	length := 0
	sizedFirst, ok := first.(ds.OrderedIterator)
	sizedLast, ok2 := end.(ds.OrderedIterator)
	if ok && ok2 {
		length = -sizedFirst.DistanceTo(sizedLast)
		if length < 0 {
			length = 0
		}
	}

	elements := make([]T, 0, length)

	for ; !first.IsEqual(end); first.Next() {
		newItem, _ := first.Get()
		elements = append(elements, newItem)
	}

	list := &Stack[T]{list: arraylist.NewFromSlice(elements)}

	return list
}

// Push adds a value onto the top of the stack
func (stack *Stack[T]) Push(value T) {
	stack.list.PushBack(value)
}

// Pop removes top element on stack and returns it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to pop.
func (stack *Stack[T]) Pop() (value T, ok bool) {
	value, ok = stack.list.Get(stack.list.Size() - 1)
	stack.list.Remove(stack.list.Size() - 1)
	return
}

// Peek returns top element on the stack without removing it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to peek.
func (stack *Stack[T]) Peek() (value T, ok bool) {
	return stack.list.Get(stack.list.Size() - 1)
}

// Empty returns true if stack does not contain any elements.
func (stack *Stack[T]) IsEmpty() bool {
	return stack.list.IsEmpty()
}

// Size returns number of elements within the stack.
func (stack *Stack[T]) Size() int {
	return stack.list.Size()
}

// Clear removes all elements from the stack.
func (stack *Stack[T]) Clear() {
	stack.list.Clear()
}

// Values returns all elements in the stack (LIFO order).
func (stack *Stack[T]) GetValues() []T {
	size := stack.list.Size()
	elements := make([]T, size, size)
	for i := 1; i <= size; i++ {
		elements[size-i], _ = stack.list.Get(i - 1) // in reverse (LIFO)
	}
	return elements
}

// String returns a string representation of container
func (stack *Stack[T]) ToString() string {
	str := "ArrayStack\n"
	values := []string{}
	for _, value := range stack.list.GetValues() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (stack *Stack[T]) withinRange(index int) bool {
	return index >= 0 && index < stack.list.Size()
}
