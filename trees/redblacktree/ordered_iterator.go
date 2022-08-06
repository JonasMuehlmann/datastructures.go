// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redblacktree

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert  implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[string, any] = (*OrderedIterator[string, any])(nil)

// Ordered holding the iterator's state
type OrderedIterator[TKey comparable, TValue any] struct {
	tree  *Tree[TKey, TValue]
	node  *Node[TKey, TValue]
	index int
	// Redundant but has better locality
	key   TKey
	value TValue
	size  int
}

//  returns a stateful iterator whose elements are key/value pairs.
func (tree *Tree[TKey, TValue]) NewOrderedIterator(t *Tree[TKey, TValue], position int) *OrderedIterator[TKey, TValue] {
	it := &OrderedIterator[TKey, TValue]{
		tree:  t,
		index: 0,
		size:  t.Size(),
	}

	if t.size == 0 {
		return it
	}

	switch position {
	case -1:
		it.node = t.Left()
		it.index = -1
	case 0:
		it.node = t.Left()
		it.index = 0
		it.key = it.node.Key
		it.value = it.node.Value
	case t.size:
		it.node = t.Right()
		it.index = t.size
	case t.size - 1:
		it.node = t.Right()
		it.index = t.size - 1
		it.key = it.node.Key
		it.value = it.node.Value
	}

	if !it.IsValid() {
		return it
	}

	it.MoveBy(position - it.index)

	return it
}

// At returns a stateful iterator whose elements are key/value pairs that is initialised at a particular node.
func (tree *Tree[TKey, TValue]) NewOrderedteratorAt(t *Tree[TKey, TValue], key TKey) *OrderedIterator[TKey, TValue] {
	it := &OrderedIterator[TKey, TValue]{tree: t, index: -1, size: t.Size()}

	it.MoveTo(key)

	return it
}


func (it *OrderedIterator[TKey, TValue]) IsBegin() bool {
	return it.index <= -1
}


func (it *OrderedIterator[TKey, TValue]) IsEnd() bool {
	return it.size == 0 || it.index >= it.size
}


func (it *OrderedIterator[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}


func (it *OrderedIterator[TKey, TValue]) IsLast() bool {
	return it.index == it.size-1
}


func (it *OrderedIterator[TKey, TValue]) IsValid() bool {
	return it.size > 0 && !it.IsBegin() && !it.IsEnd()
}


func (it *OrderedIterator[TKey, TValue]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*OrderedIterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) == 0
}


func (it *OrderedIterator[TKey, TValue]) DistanceTo(other ds.OrderedIterator) int {
	otherThis, ok := other.(*OrderedIterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.index - otherThis.index
}


func (it *OrderedIterator[TKey, TValue]) IsAfter(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*OrderedIterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) > 0
}


func (it *OrderedIterator[TKey, TValue]) IsBefore(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*OrderedIterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) < 0
}


func (it *OrderedIterator[TKey, TValue]) Size() int {
	return it.size
}

// Next moves the  to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
// If Next() was called for the first time, then it will point the  to the first element if it exists.
// Modifies the state of the .
func (it *OrderedIterator[TKey, TValue]) Next() bool {
	it.index = utils.Min(it.index+1, it.size)

	if !it.IsValid() {
		return false
	}

	if it.IsFirst() {
		it.node = it.tree.Left()

		it.key = it.node.Key
		it.value = it.node.Value

		return true
	}

	if it.node.Right != nil {
		it.node = it.node.Right

		for it.node.Left != nil {
			it.node = it.node.Left
		}
	} else {
		for it.node.Parent != nil {
			node := it.node
			it.node = it.node.Parent

			if node == it.node.Left {
				break
			}
		}
	}

	it.key = it.node.Key
	it.value = it.node.Value

	return true
}

func (it *OrderedIterator[TKey, TValue]) NextN(n int) bool {
	var found bool

	for i := 0; i < n; i++ {
		found = it.Next()
	}

	return found
}

// Prev moves the  to the previous element and returns true if there was a previous element in the container.
// If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value().
// Modifies the state of the .
func (it *OrderedIterator[TKey, TValue]) Previous() bool {
	it.index = utils.Max(it.index-1, -1)

	if !it.IsValid() {
		return false
	}

	if it.IsLast() {
		it.node = it.tree.Right()

		it.key = it.node.Key
		it.value = it.node.Value

		return true
	}

	if it.node.Left != nil {
		it.node = it.node.Left

		for it.node.Right != nil {
			it.node = it.node.Right
		}
	} else {
		for it.node.Parent != nil {
			node := it.node
			it.node = it.node.Parent

			if node == it.node.Right {
				break
			}
		}
	}

	it.key = it.node.Key
	it.value = it.node.Value

	return true

}

func (it *OrderedIterator[TKey, TValue]) PreviousN(n int) bool {
	var found bool

	for i := 0; i < n; i++ {
		found = it.Previous()
	}

	return found
}

func (it *OrderedIterator[TKey, TValue]) MoveBy(n int) bool {
	if n > 0 {
		return it.NextN(n)
	} else if n < 0 {
		return it.PreviousN(-n)
	}

	return it.IsValid()
}

// https://www.geeksforgeeks.org/find-distance-between-two-nodes-of-a-binary-tree/
func (it *OrderedIterator[TKey, TValue]) MoveTo(key TKey) (found bool) {
	if it.IsValid() && it.tree.Comparator(key, it.node.Key) == 0 {

		return true
	}

	targetNode := it.tree.lookup(key)
	if targetNode == nil {
		return false
	}

	distance := distanceBetween(it.tree.Comparator, it.tree.Root, it.node, targetNode)
	if it.IsBegin() || it.IsEnd() {
		// Even if the index is Begin(), the starting node is First()
		distance += 1
	}

	if it.tree.Comparator(key, it.node.Key) < 0 || it.IsEnd() {
		distance = -distance
	}

	it.index += distance

	it.key = targetNode.Key
	it.value = targetNode.Value

	return true
}

// Value returns the current element's value.
// Does not modify the state of the .
func (it *OrderedIterator[TKey, TValue]) Get() (value TValue, found bool) {
	if !it.IsValid() {
		return
	}

	return it.node.Value, true
}

func (it *OrderedIterator[TKey, TValue]) GetAt(key TKey) (value TValue, found bool) {
	return it.tree.Get(key)
}

func (it *OrderedIterator[TKey, TValue]) SetAt(key TKey, value TValue) bool {
	it.tree.Put(key, value)

	return true
}

func (it *OrderedIterator[TKey, TValue]) Set(value TValue) bool {
	if !it.IsValid() {
		return false
	}

	it.value = value
	it.node.Value = value

	return true
}

// Key returns the current element's key.
// Does not modify the state of the .
func (it *OrderedIterator[TKey, TValue]) Index() (key TKey, found bool) {
	if !it.IsValid() {
		return
	}

	return it.key, true
}

// Node returns the current element's node.
// Does not modify the state of the .
func (it *OrderedIterator[TKey, TValue]) Node() (*Node[TKey, TValue], bool) {
	return it.node, it.IsValid()
}
