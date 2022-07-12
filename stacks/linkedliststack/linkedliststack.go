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
func New[T any]() *Stack[T] {
	return &Stack[T]{list: &singlylinkedlist.List[T]{}}
}

// Push adds a value onto the top of the stack
func (stack *Stack[T]) Push(value T) {
	stack.list.Prepend(value)
}

// Pop removes top element on stack and returns it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to pop.
func (stack *Stack[T]) Pop() (value T, ok bool) {
	value, ok = stack.list.Get(0)
	stack.list.Remove(0)
	return
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
