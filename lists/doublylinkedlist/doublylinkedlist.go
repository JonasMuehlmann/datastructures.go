// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package doublylinkedlist implements the doubly-linked list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/List[T]_%28abstract_data_type%29
package doublylinkedlist

import (
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/lists"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert List[T] implementation
var _ lists.List[any] = (*List[any])(nil)

// List[T] holds the elements, where each element points to the next and previous element
type List[T any] struct {
	first *element[T]
	last  *element[T]
	size  int
}

type element[T any] struct {
	value T
	prev  *element[T]
	next  *element[T]
}

// New instantiates a new list and adds the passed values, if any, to the list
func New[T any](values ...T) *List[T] {
	list := &List[T]{}
	if len(values) > 0 {
		list.PushBack(values...)
	}
	return list
}

// NewFromSlice instantiates a new list containing the provided slice.
func NewFromSlice[T any](slice []T) *List[T] {
	list := &List[T]{}

	for _, element := range slice {
		list.PushBack(element)
	}

	return list
}

// NewFromIterator instantiates a new list containing the elements provided by the passed iterator.
func NewFromIterator[T any](begin ds.ReadCompForIterator[T]) *List[T] {
	list := &List[T]{}

	for begin.Next() {
		newItem, _ := begin.Get()
		list.PushBack(newItem)
	}

	return list
}

// NewFromIterators instantiates a new list containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[T any](begin ds.ReadCompForIterator[T], end ds.ComparableIterator) *List[T] {
	list := &List[T]{}
	for !begin.IsEqual(end) && begin.Next() {
		newItem, _ := begin.Get()
		list.PushBack(newItem)
	}

	return list
}

func (list *List[T]) PopBack(n int) (popped []T) {
	if list.size < n || n == 0 || list.size == 0 {
		return
	}

	popped = make([]T, 0, n)

	for i := 0; i < n; i++ {
		popped = append(popped, list.last.value)

		list.last = list.last.prev
		list.last.next = nil
	}

	list.size -= n

	return
}

func (list *List[T]) PopFront(n int) (popped []T) {
	if list.size < n || n == 0 || list.size == 0 {
		return
	}

	popped = make([]T, 0, n)

	for i := 0; i < n; i++ {
		popped = append(popped, list.first.value)

		list.first = list.first.next
		list.first.prev = nil
	}

	list.size -= n

	return
}

// Add appends a value (one or more) at the end of the list (same as Append())
func (list *List[T]) PushBack(values ...T) {
	for _, value := range values {
		newElement := &element[T]{value: value, prev: list.last}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

// Prepend prepends a values (or more)
func (list *List[T]) PushFront(values ...T) {
	// in reverse to keep passed order i.e. ["c","d"] -> PushFront(["a","b"]) -> ["a","b","c",d"]
	for v := len(values) - 1; v >= 0; v-- {
		newElement := &element[T]{value: values[v], next: list.first}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.first.prev = newElement
			list.first = newElement
		}
		list.size++
	}
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List[T]) Get(index int) (value T, found bool) {

	if !list.withinRange(index) {
		return
	}

	// determine traveral direction, last to first or first to last
	if list.size-index < index {
		element := list.last
		for e := list.size - 1; e != index; e, element = e-1, element.prev {
		}
		return element.value, true
	}
	element := list.first
	for e := 0; e != index; e, element = e+1, element.next {
	}
	return element.value, true
}

// Remove removes the element at the given index from the list.
func (list *List[T]) Remove(index int) {

	if !list.withinRange(index) {
		return
	}

	if list.size == 1 {
		list.Clear()
		return
	}

	var element *element[T]
	// determine traversal direction, last to first or first to last
	if list.size-index < index {
		element = list.last
		for e := list.size - 1; e != index; e, element = e-1, element.prev {
		}
	} else {
		element = list.first
		for e := 0; e != index; e, element = e+1, element.next {
		}
	}

	if element == list.first {
		list.first = element.next
	}
	if element == list.last {
		list.last = element.prev
	}
	if element.prev != nil {
		element.prev.next = element.next
	}
	if element.next != nil {
		element.next.prev = element.prev
	}

	element = nil

	list.size--
}

// Contains check if values (one or more) are present in the set.
// All values have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *List[T]) Contains(comparator utils.Comparator[T], values ...T) bool {

	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}
	for _, value := range values {
		found := false
		for element := list.first; element != nil; element = element.next {
			if comparator(element.value, value) == 0 {
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
	values := make([]T, list.size, list.size)
	for e, element := 0, list.first; element != nil; e, element = e+1, element.next {
		values[e] = element.value
	}
	return values
}

//IndexOf returns index of provided element
func (list *List[T]) IndexOf(comparator utils.Comparator[T], value T) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.GetValues() {
		if comparator(element, value) == 0 {
			return index
		}
	}
	return -1
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
	list.first = nil
	list.last = nil
}

// Sort sorts values (in-place) using.
func (list *List[T]) Sort(comparator utils.Comparator[T]) {

	if list.size < 2 {
		return
	}

	values := list.GetValues()
	utils.Sort(values, comparator)

	list.Clear()

	list.PushBack(values...)

}

// Swap swaps values of two elements at the given indices.
func (list *List[T]) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) && i != j {
		var element1, element2 *element[T]
		for e, currentElement := 0, list.first; element1 == nil || element2 == nil; e, currentElement = e+1, currentElement.next {
			switch e {
			case i:
				element1 = currentElement
			case j:
				element2 = currentElement
			}
		}
		element1.value, element2.value = element2.value, element1.value
	}
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[T]) Insert(index int, values ...T) {

	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.PushBack(values...)
		}
		return
	}

	list.size += len(values)

	var beforeElement *element[T]
	var foundElement *element[T]
	// determine traversal direction, last to first or first to last
	if list.size-index < index {
		foundElement = list.last
		for e := list.size - 1; e != index; e, foundElement = e-1, foundElement.prev {
			beforeElement = foundElement.prev
		}
	} else {
		foundElement = list.first
		for e := 0; e != index; e, foundElement = e+1, foundElement.next {
			beforeElement = foundElement
		}
	}

	if foundElement == list.first {
		oldNextElement := list.first
		for i, value := range values {
			newElement := &element[T]{value: value}
			if i == 0 {
				list.first = newElement
			} else {
				newElement.prev = beforeElement
				beforeElement.next = newElement
			}
			beforeElement = newElement
		}
		oldNextElement.prev = beforeElement
		beforeElement.next = oldNextElement
	} else {
		oldNextElement := beforeElement.next
		for _, value := range values {
			newElement := &element[T]{value: value}
			newElement.prev = beforeElement
			beforeElement.next = newElement
			beforeElement = newElement
		}
		oldNextElement.prev = beforeElement
		beforeElement.next = oldNextElement
	}
}

// Set value at specified index position
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List[T]) Set(index int, value T) {

	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.PushBack(value)
		}
		return
	}

	var foundElement *element[T]
	// determine traversal direction, last to first or first to last
	if list.size-index < index {
		foundElement = list.last
		for e := list.size - 1; e != index; {
			fmt.Println("Set last", index, value, foundElement, foundElement.prev)
			e, foundElement = e-1, foundElement.prev
		}
	} else {
		foundElement = list.first
		for e := 0; e != index; {
			e, foundElement = e+1, foundElement.next
		}
	}

	foundElement.value = value
}

// String returns a string representation of container
func (list *List[T]) ToString() string {
	str := "DoublyLinkedList[T]\n"
	values := []string{}
	for element := list.first; element != nil; element = element.next {
		values = append(values, fmt.Sprintf("%v", element.value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (list *List[T]) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

//******************************************************************//
//                             Iterator                             //
//******************************************************************//

// Begin returns an initialized iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (list *List[T]) Begin() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return list.NewIterator(-1)
}

// End returns an initialized iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (list *List[T]) End() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return list.NewIterator(list.size)
}

// First returns an initialized iterator, which points to it's first element.
func (list *List[T]) First() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return list.NewIterator(0)
}

// Last returns an initialized iterator, which points to it's last element.
func (list *List[T]) Last() ds.ReadWriteOrdCompBidRandCollIterator[int, T] {
	return list.NewIterator(list.size - 1)
}
