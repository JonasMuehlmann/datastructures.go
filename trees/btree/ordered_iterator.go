// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package btree

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[string, any] = (*OrderedIterator[string, any])(nil)

// Ordered holding the iterator's state
type OrderedIterator[TKey comparable, TValue any] struct {
	tree          *Tree[TKey, TValue]
	node          *Node[TKey, TValue]
	index         int
	iCurrentEntry int
	// Redundant but has better locality
	size  int
	key   TKey
	value TValue
}

//  returns a stateful iterator whose elements are key/value pairs.
func (tree *Tree[TKey, TValue]) NewOrderedIterator(t *Tree[TKey, TValue], position int) *OrderedIterator[TKey, TValue] {
	it := &OrderedIterator[TKey, TValue]{
		tree: t,
		// index: 0,
		index: -1,
		size:  t.Size(),
	}

	if t.size == 0 {
		return it
	}

	// switch position {
	// case -1:
	// 	it.node = t.Left()
	// 	it.index = -1
	// case 0:
	// 	it.node = t.Left()
	// 	it.index = 0
	// 	it.iCurrentEntry = len(it.node.Entries)
	// 	it.key = it.node.Entries[it.iCurrentEntry].Key
	// 	it.value = it.node.Entries[it.iCurrentEntry].Value
	// case t.size:
	// 	it.node = t.Right()
	// 	it.index = t.size
	// case t.size - 1:
	// 	it.node = t.Right()
	// 	it.index = t.size - 1
	// 	it.iCurrentEntry = len(it.node.Entries)
	// 	it.key = it.node.Entries[it.iCurrentEntry].Key
	// 	it.value = it.node.Entries[it.iCurrentEntry].Value
	// }

	// if !it.IsValid() {
	// 	return it
	// }

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
	// If already at end, go to end
	if it.index == it.size {
		goto end
	}
	// If at beginning, get the left-most entry in the tree
	if it.index == -1 {
		left := it.tree.Left()
		if left == nil {
			goto end
		}

		it.node = left
		it.iCurrentEntry = 0

		goto between
	}
	{
		// Find current entry position in current node
		e, _ := it.tree.search(it.node, it.key)
		// Try to go down to the child right of the current entry
		if e+1 < len(it.node.Children) {
			it.node = it.node.Children[e+1]

			// Try to go down to the child left of the current node
			for len(it.node.Children) > 0 {
				it.node = it.node.Children[0]
			}
			// Return the left-most entry
			it.iCurrentEntry = 0

			goto between
		}

		// Above assures that we have reached a leaf node, so return the next entry in current node (if any)
		if e+1 < len(it.node.Entries) {
			it.iCurrentEntry = e + 1

			goto between
		}
	}
	// Reached leaf node and there are no entries to the right of the current entry, so go up to the parent
	for it.node.Parent != nil {
		it.node = it.node.Parent

		// Find next entry position in current node (note: search returns the first equal or bigger than entry)
		e, _ := it.tree.search(it.node, it.key)
		// Check that there is a next entry position in current node
		if e < len(it.node.Entries) {
			it.iCurrentEntry = e
			goto between
		}
	}

end:
	it.index = it.size
	return false

between:
	it.index = utils.Min(it.index+1, it.size)
	it.key = it.node.Entries[it.iCurrentEntry].Key
	it.value = it.node.Entries[it.iCurrentEntry].Value

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
	// If already at beginning, go to begin
	if it.index == -1 {
		goto begin
	}
	// If at end, get the right-most entry in the tree
	if it.index == it.size {
		right := it.tree.Right()
		if right == nil {
			goto begin
		}

		it.node = right
		it.iCurrentEntry = len(right.Entries) - 1

		goto between
	}
	{
		// Find current entry position in current node
		e, _ := it.tree.search(it.node, it.key)
		// Try to go down to the child left of the current entry
		if e < len(it.node.Children) {
			it.node = it.node.Children[e]
			// Try to go down to the child right of the current node
			for len(it.node.Children) > 0 {
				it.node = it.node.Children[len(it.node.Children)-1]
			}

			// Return the right-most entry
			it.iCurrentEntry = len(it.node.Entries) - 1

			goto between
		}
		// Above assures that we have reached a leaf node, so return the previous entry in current node (if any)
		if e-1 >= 0 {
			it.iCurrentEntry = e - 1

			goto between
		}
	}
	// Reached leaf node and there are no entries to the left of the current entry, so go up to the parent
	for it.node.Parent != nil {
		it.node = it.node.Parent

		// Find previous entry position in current node (note: search returns the first equal or bigger than entry)
		e, _ := it.tree.search(it.node, it.key)
		// Check that there is a previous entry position in current node
		if e-1 >= 0 {
			it.iCurrentEntry = e - 1

			goto between
		}
	}

begin:
	it.index = -1
	return false

between:
	it.index = utils.Max(it.index-1, -1)
	it.key = it.node.Entries[it.iCurrentEntry].Key
	it.value = it.node.Entries[it.iCurrentEntry].Value

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
	} else {
		return it.PreviousN(-n)
	}
}

// https://www.geeksforgeeks.org/find-distance-between-two-nodes-of-a-binary-tree/
func (it *OrderedIterator[TKey, TValue]) MoveTo(key TKey) (found bool) {
	// if it.IsValid() && it.tree.Comparator(key, it.node.Entries[it.iCurrentEntry].Key) == 0 {

	// 	return true
	// }

	// targetNode, targetNodeIndex, found := it.tree.searchRecursively(it.tree.Root, key)
	// if !found {
	// 	return false
	// }

	// distance := distanceBetween(it.tree.Comparator, it.tree.Root, it.node, it.iCurrentEntry, targetNode, targetNodeIndex)
	// if it.IsBegin() || it.IsEnd() {
	// 	// Even if the index is Begin(), the starting node is First()
	// 	distance += 1
	// }

	// if it.tree.Comparator(key, it.node.Entries[it.iCurrentEntry].Key) < 0 || it.IsEnd() {
	// 	distance = -distance
	// }

	// it.index += distance
	// it.iCurrentEntry = targetNodeIndex

	// it.key = targetNode.Entries[it.iCurrentEntry].Key
	// it.value = targetNode.Entries[it.iCurrentEntry].Value

	// return true

	cmp := it.tree.Comparator(it.key, key)
	if cmp == 0 {
		if it.IsEnd() {
			it.Previous()
		}

		return true
	} else if cmp < 0 {
		for it.Next() {
			if it.tree.Comparator(it.key, key) == 0 {
				return true
			}
		}
	} else {

		for it.Previous() {
			if it.tree.Comparator(it.key, key) == 0 {
				return true
			}

		}
	}

	return false
}

// Value returns the current element's value.
// Does not modify the state of the .
func (it *OrderedIterator[TKey, TValue]) Get() (value TValue, found bool) {
	if !it.IsValid() {
		return
	}

	return it.node.Entries[it.iCurrentEntry].Value, true
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
	it.node.Entries[it.iCurrentEntry].Value = value

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
