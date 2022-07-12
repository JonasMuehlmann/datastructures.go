// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package arraylist implements the array list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
package arraylist

import (
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/lists"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert List implementation.
var _ lists.List[any] = (*List[any])(nil)

// TODO: Try and reimplement methods through iterator
// List holds the elements in a slice.
type List[T any] struct {
	elements []T
	size     int
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// New instantiates a new list and adds the passed values, if any, to the list.
func New[T any](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// NewFromSLice instantiates a new list containing the provided slice.
func NewFromSlice[T any](slice []T) *List[T] {
	list := &List[T]{elements: slice, size: len(slice)}
	return list
}

func (list *List[T]) PushBack(values ...T)  { panic("Not implemented") }
func (list *List[T]) PushFront(values ...T) { panic("Not implemented") }
func (list *List[T]) PopBack()              { panic("Not implemented") }
func (list *List[T]) PopFront()             { panic("Not implemented") }

// Add appends a value at the end of the list.
func (list *List[T]) Add(values ...T) {
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List[T]) Get(index int) (value T, wasFound bool) {
	if !list.withinRange(index) {
		return
	}

	return list.elements[index], true
}

// TODO: Implement RemoveStable which does a swap and shrink
// Remove removes the element at the given index from the list.
func (list *List[T]) Remove(index int) {
	if !list.withinRange(index) {
		return
	}

	copy(list.elements[index:], list.elements[index+1:list.size]) // shift to the left by one (slow operation, need ways to optimize this)
	list.size--

	list.shrink()
}

// PERF: Maybe we can provide separated implementations of the data structures (e.g. BasicList) through code generation, which are constrained with comparable
// PERF: Iterate over elements only once and keep counter of found values
// Contains checks if elements (one or more) are present in the set.
// All elements have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *List[T]) Contains(comparator utils.Comparator[T], values ...T) bool {
	for _, searchValue := range values {
		found := false
		for index := 0; index < list.size; index++ {
			if comparator(list.elements[index], searchValue) == 0 {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Values returns all elements in the list.
func (list *List[T]) GetValues() []T {
	newElements := make([]T, list.size, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

func (list *List[T]) IndexOf(comparator utils.Comparator[T], value T) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.elements {
		if comparator(element, value) == 0 {
			return index
		}
	}
	return -1
}

// GetSlice returns the underlying slice.
func (list *List[T]) GetSlice() []T {
	return list.elements
}

// Empty returns true if list does not contain any elements.
func (list *List[T]) IsEmpty() bool {
	return list.size == 0
}

// Size returns number of elements within the list.
func (list *List[T]) Size() int {
	return list.size
}

// Clear removes all elements from the list.
func (list *List[T]) Clear() {
	list.size = 0
	list.elements = []T{}
}

// Sort sorts values (in-place) using.
func (list *List[T]) Sort(comparator utils.Comparator[T]) {
	if len(list.elements) < 2 {
		return
	}
	utils.Sort(list.elements[:list.size], comparator)
}

// Swap swaps the two values at the specified positions.
func (list *List[T]) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[T]) Insert(index int, values ...T) {
	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	l := len(values)
	list.growBy(l)
	list.size += l
	copy(list.elements[index+l:], list.elements[index:list.size-l])
	copy(list.elements[index:], values)
}

// Set the value at specified index
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[T]) Set(index int, value T) {
	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(value)
		}
		return
	}

	list.elements[index] = value
}

// String returns a string representation of container.
func (list *List[T]) ToString() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

//******************************************************************//
//                              Helper                              //
//******************************************************************//

// Check that the index is within bounds of the list.
func (list *List[T]) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

func (list *List[T]) resize(cap int) {
	newElements := make([]T, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

// Expand the array if necessary, i.e. capacity will be reached if we add n elements.
func (list *List[T]) growBy(n int) {
	// When capacity is reached, grow by a factor of growthFactor and add number of elements
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

// Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity.
func (list *List[T]) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	// Shrink when size is at shrinkFactor * capacity
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}

//******************************************************************//
//                             Iterator                             //
//******************************************************************//

// Begin returns an initialized iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (list *List[T]) Begin() ds.RWOrdCompBidRandCollIterator[T, int] {
	return list.NewIterator(list, -1)
}

// End returns an initialized iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (list *List[T]) End() ds.RWOrdCompBidRandCollIterator[T, int] {
	return list.NewIterator(list, list.size)
}

// First returns an initialized iterator, which points to it's first element.
func (list *List[T]) First() ds.RWOrdCompBidRandCollIterator[T, int] {
	return list.NewIterator(list, 0)
}

// Last returns an initialized iterator, which points to it's last element.
func (list *List[T]) Last() ds.RWOrdCompBidRandCollIterator[T, int] {
	return list.NewIterator(list, list.size-1)
}

//******************************************************************//
//                         Reverse iterator                         //
//******************************************************************//

// ReverseBegin returns an initialized, reversed iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (list *List[T]) ReverseBegin() ds.RWOrdCompBidRevRandCollIterator[T, int] {
	return list.NewReverseIterator(list, list.size)
}

// ReverseEnd returns an initialized,reversed iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (list *List[T]) ReverseEnd() ds.RWOrdCompBidRevRandCollIterator[T, int] {
	return list.NewReverseIterator(list, -1)
}

// ReverseFirst returns an initialized, reversed iterator, which points to it's first element.
func (list *List[T]) ReverseFirst() ds.RWOrdCompBidRevRandCollIterator[T, int] {
	return list.NewReverseIterator(list, list.size-1)
}

// ReverseLast returns an initialized, reversed iterator, which points to it's last element.
func (list *List[T]) ReverseLast() ds.RWOrdCompBidRevRandCollIterator[T, int] {
	return list.NewReverseIterator(list, 0)
}
