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
func NewFromIterator[T any](begin ds.ReadForIterator[T]) *Stack[T] {
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

	list := &Stack[T]{list: arraylist.NewFromSlice(elements)}

	return list
}

// NewFromIterators instantiates a new stack containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[T any](begin ds.ReadCompForIterator[T], end ds.ComparableIterator) *Stack[T] {
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
	if stack.Size() == 0 {
		return
	}

	return stack.list.PopBack(1)[0], true
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
	return stack.list.GetValues()
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

//******************************************************************//
//                             Iterator                             //
//******************************************************************//

// Begin returns an initialized iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (set *Stack[T]) Begin() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return set.NewIterator(-1, set.Size())
}

// End returns an initialized iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (set *Stack[T]) End() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return set.NewIterator(set.list.Size(), set.Size())
}

// First returns an initialized iterator, which points to it's first element.
func (set *Stack[T]) First() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return set.NewIterator(0, set.Size())
}

// Last returns an initialized iterator, which points to it's last element.
func (set *Stack[T]) Last() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return set.NewIterator(set.list.Size()-1, set.Size())
}
