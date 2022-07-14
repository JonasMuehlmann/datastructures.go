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
}

const (
	ShrinkThresholdPercent  = float32(0.25) // shrink when cap * ShrinkThresholdPercent > len (0 means never shrink)
	ShrinkThresholdAbsolute = 100           // shrink when cap - len >= ShrinkThresholdAbsolute
	ShrinkFactor            = float32(0.5)  // shrink by ShrinkFactor * cap - len
)

// New instantiates a new list and adds the passed values, if any, to the list.
func New[T any](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.PushBack(values...)
	}
	return list
}

// NewFromSLice instantiates a new list containing the provided slice.
func NewFromSlice[T any](slice []T) *List[T] {
	list := &List[T]{elements: slice}
	return list
}

// Add appends a value at the end of the list.
func (list *List[T]) PushBack(values ...T) {
	list.elements = append(list.elements, values...)

}

func (list *List[T]) PushFront(values ...T) {
	list.elements = append(values, list.elements...)
}

func (list *List[T]) PopBack(n int) {
	if len(list.elements) > 0 && len(list.elements) >= n {
		list.elements = list.elements[:len(list.elements)-n]
	}
}
func (list *List[T]) PopFront(n int) {
	if len(list.elements) > 0 && len(list.elements) >= n {
		list.elements = list.elements[n:]
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

// Remove removes the element at the given index from the list.
// The order of the elements is allowed to be changed for performance reasons.
func (list *List[T]) Remove(index int) {
	if !list.withinRange(index) {
		return
	}

	list.elements[index] = list.elements[len(list.elements)-1]
	list.elements = list.elements[:len(list.elements)-1]
}

// RemoveStable removes the element at the given index from the list.
// THe order of the elements is NOT allowed to be changed.
func (list *List[T]) RemoveStable(index int) {
	if !list.withinRange(index) {
		return
	}

	list.elements = append(list.elements[:index], list.elements[index+1:]...) // shift to the left by one (slow operation, need ways to optimize this)
}

// Contains checks if elements (one or more) are present in the set.
// All elements have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *List[T]) Contains(comparator utils.Comparator[T], values ...T) bool {
	for _, searchValue := range values {
		found := false
		for index := 0; index < len(list.elements); index++ {
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
	newElements := make([]T, len(list.elements), len(list.elements))
	copy(newElements, list.elements[:len(list.elements)])
	return newElements
}

func (list *List[T]) IndexOf(comparator utils.Comparator[T], value T) int {
	if len(list.elements) == 0 {
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
	return len(list.elements) == 0
}

// Size returns number of elements within the list.
func (list *List[T]) Size() int {
	return len(list.elements)
}

// Clear removes all elements from the list.
func (list *List[T]) Clear() {
	list.elements = []T{}
}

// Sort sorts values (in-place) using.
func (list *List[T]) Sort(comparator utils.Comparator[T]) {
	if len(list.elements) < 2 {
		return
	}
	utils.Sort(list.elements[:len(list.elements)], comparator)
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
		if index == len(list.elements) {
			list.PushBack(values...)
		}
		return
	}

	newList := make([]T, len(list.elements)+len(values))

	copy(newList[:index], list.elements[:index])
	copy(newList[index:index+len(values)], values)
	copy(newList[index+len(values):], list.elements[index+len(values)-1:])

	list.elements = newList
}

// Set the value at specified index
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[T]) Set(index int, value T) {
	if !list.withinRange(index) {
		// Append
		if index == len(list.elements) {
			list.PushBack(value)
		}
		return
	}

	list.elements[index] = value
}

// String returns a string representation of container.
func (list *List[T]) ToString() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list.elements[:len(list.elements)] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// ShrinkToFit shrinks the array so that len == cap.
func (list *List[T]) ShrinkToFit() {
	list.elements = append([]T{}, list.elements...)
}

// TryShrink the array if possible/worthwhile.
// Shrinking is worthwile if:
// cap - len > ShrinkThresholdAbsolute and
// cap * ShrinkThresholdPercent > len
//
// To reduce the number of reslices upon appending, the new length will be len + ((cap - len) * ShrinkFactor).
func (list *List[T]) TryShrink() {
	if ShrinkThresholdPercent == 0.0 {
		return
	}
	currentCapacity := cap(list.elements)
	currentLength := len(list.elements)
	diff := currentCapacity - currentLength

	if diff > ShrinkThresholdAbsolute && float32(currentCapacity)*ShrinkThresholdPercent > float32(currentLength) {
		newElements := make([]T, currentLength+int(float32(diff)*ShrinkFactor))
		copy(newElements, list.elements)
		list.elements = newElements
	}
}

//******************************************************************//
//                              Helper                              //
//******************************************************************//

// Check that the index is within bounds of the list.
func (list *List[T]) withinRange(index int) bool {
	return index >= 0 && index < len(list.elements)
}

//******************************************************************//
//                             Iterator                             //
//******************************************************************//

// Begin returns an initialized iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (list *List[T]) Begin() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return list.NewIterator(list, -1)
}

// End returns an initialized iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (list *List[T]) End() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return list.NewIterator(list, len(list.elements))
}

// First returns an initialized iterator, which points to it's first element.
func (list *List[T]) First() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return list.NewIterator(list, 0)
}

// Last returns an initialized iterator, which points to it's last element.
func (list *List[T]) Last() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return list.NewIterator(list, len(list.elements)-1)
}

//******************************************************************//
//                         Reverse iterator                         //
//******************************************************************//

// ReverseBegin returns an initialized, reversed iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (list *List[T]) ReverseBegin() ds.ReadWriteOrdCompBidRevRandCollIterator[int, T] {
	return list.NewReverseIterator(list, len(list.elements))
}

// ReverseEnd returns an initialized,reversed iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (list *List[T]) ReverseEnd() ds.ReadWriteOrdCompBidRevRandCollIterator[int, T] {
	return list.NewReverseIterator(list, -1)
}

// ReverseFirst returns an initialized, reversed iterator, which points to it's first element.
func (list *List[T]) ReverseFirst() ds.ReadWriteOrdCompBidRevRandCollIterator[int, T] {
	return list.NewReverseIterator(list, len(list.elements)-1)
}

// ReverseLast returns an initialized, reversed iterator, which points to it's last element.
func (list *List[T]) ReverseLast() ds.ReadWriteOrdCompBidRevRandCollIterator[int, T] {
	return list.NewReverseIterator(list, 0)
}
