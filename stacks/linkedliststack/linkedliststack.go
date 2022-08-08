// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linkedliststack implements a stack backed by a singly-linked list.
//
// Structure is not thread safe.
//
// Reference:https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29#Linked_list
package linkedliststack

import (
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/lists/singlylinkedlist"
	"github.com/JonasMuehlmann/datastructures.go/stacks"
)

// Assert Stack implementation
var _ stacks.Stack[any] = (*Stack[any])(nil)

// Stack holds elements in a singly-linked-list
type Stack[T any] struct {
	list *singlylinkedlist.List[T]
}

// New nnstantiates a new empty stack
func New[T any](values ...T) *Stack[T] {
	return &Stack[T]{list: singlylinkedlist.New[T](values...)}
}

// NewFromSlice instantiates a new stack containing the provided slice.
func NewFromSlice[T any](slice []T) *Stack[T] {
	list := &Stack[T]{list: singlylinkedlist.NewFromSlice(slice)}

	return list
}

// NewFromIterator instantiates a new stack containing the elements provided by the passed iterator.
func NewFromIterator[T any](begin ds.ReadCompForIterator[T]) *Stack[T] {
	list := &Stack[T]{list: singlylinkedlist.New[T]()}

	for begin.Next() {
		newItem, _ := begin.Get()
		list.Push(newItem)
	}

	return list
}

// NewFromIterators instantiates a new stack containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[T any](begin ds.ReadCompForIterator[T], end ds.ComparableIterator) *Stack[T] {
	list := &Stack[T]{list: singlylinkedlist.New[T]()}
	for !begin.IsEqual(end) && begin.Next() {
		newItem, _ := begin.Get()
		list.Push(newItem)
	}

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
	return stack.list.Get(0)
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
	str := "LinkedListStack\n"
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
func (stack *Stack[T]) Begin() ds.ReadWriteOrdCompForRandCollIterator[T] {
	return stack.NewIterator(-1, stack.Size())
}

// End returns an initialized iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (stack *Stack[T]) End() ds.ReadWriteOrdCompForRandCollIterator[T] {
	return stack.NewIterator(stack.Size(), stack.Size())
}

// First returns an initialized iterator, which points to it's first element.
func (stack *Stack[T]) First() ds.ReadWriteOrdCompForRandCollIterator[T] {
	return stack.NewIterator(0, stack.Size())
}

// Last returns an initialized iterator, which points to it's last element.
func (stack *Stack[T]) Last() ds.ReadWriteOrdCompForRandCollIterator[T] {
	return stack.NewIterator(stack.Size()-1, stack.Size())
}
