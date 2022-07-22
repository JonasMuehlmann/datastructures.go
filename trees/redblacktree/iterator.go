// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redblacktree

import "github.com/JonasMuehlmann/datastructures.go/ds"

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[any, any] = (*OrderedIterator[any, any])(nil)

// OrderedIterator holding the iterator's state
type OrderedIterator[TKey any, TValue any] struct {
	tree  *Tree[TKey, TValue]
	node  *Node[TKey, TValue]
	index int
}

// Iterator returns a stateful iterator whose elements are key/value pairs.
func (tree *Tree[TKey, TValue]) NewIterator(t *Tree[TKey, TValue], position int) *OrderedIterator[TKey, TValue] {
	it := &OrderedIterator[TKey, TValue]{tree: t, node: nil, index: position}

	if position > 0 && position < it.tree.Size() {
		for it.index != position {
			it.Next()
		}
	}

	return it
}

// IteratorAt returns a stateful iterator whose elements are key/value pairs that is initialised at a particular node.
func (tree *Tree[TKey, TValue]) NewIteratorAt(t *Tree[TKey, TValue], node *Node[TKey, TValue]) *OrderedIterator[TKey, TValue] {
	it := &OrderedIterator[TKey, TValue]{tree: t, node: nil}

	for it.node != node && !it.IsEnd() {
		it.Next()
	}

	return it

}

// IsBegin implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsBegin() bool {
	return it.index == -1
}

// IsEnd implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsEnd() bool {
	return it.tree.Size() == 0 || it.index == it.tree.Size()
}

// IsFirst implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}

// IsLast implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsLast() bool {
	return it.index == it.tree.Size()-1
}

// IsValid implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsValid() bool {
	return it.tree.Size() > 0 && it.index >= 0 && it.index < it.tree.Size()
}

// IsEqual implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*OrderedIterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) == 0
}

// DistanceTo implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) DistanceTo(other ds.OrderedIterator) int {
	otherThis, ok := other.(*OrderedIterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.index - otherThis.index
}

// IsAfter implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsAfter(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*OrderedIterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) > 0
}

// IsBefore implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) IsBefore(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*OrderedIterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) < 0
}

// Size implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[TKey, TValue]) Size() int {
	return it.tree.Size()
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's key and value can be retrieved by Key() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *OrderedIterator[TKey, TValue]) Next() {
	if iterator.IsEnd() {
		goto end
	}
	if iterator.IsBegin() {
		left := iterator.tree.Left()
		if left == nil {
			goto end
		}
		iterator.node = left
		goto between
	}
	if iterator.node.Right != nil {
		iterator.node = iterator.node.Right
		for iterator.node.Left != nil {
			iterator.node = iterator.node.Left
		}
		goto between
	}
	for iterator.node.Parent != nil {
		node := iterator.node
		iterator.node = iterator.node.Parent
		if node == iterator.node.Left {
			goto between
		}
	}

end:
	iterator.node = nil
	iterator.index = iterator.tree.Size()

between:
	iterator.index++
}

func (iterator *OrderedIterator[TKey, TValue]) NextN(n int) {
	for i := 0; i < n; i++ {
		iterator.Next()
	}
}

// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
// If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value().
// Modifies the state of the iterator.
func (iterator *OrderedIterator[TKey, TValue]) Prev() {
	if iterator.IsBegin() {
		goto begin
	}
	if iterator.IsEnd() {
		right := iterator.tree.Right()
		if right == nil {
			goto begin
		}
		iterator.node = right
		goto between
	}
	if iterator.node.Left != nil {
		iterator.node = iterator.node.Left
		for iterator.node.Right != nil {
			iterator.node = iterator.node.Right
		}
		goto between
	}
	for iterator.node.Parent != nil {
		node := iterator.node
		iterator.node = iterator.node.Parent
		if node == iterator.node.Right {
			goto between
		}
	}

begin:
	iterator.node = nil
	iterator.index = -1

between:
	iterator.index--
}

func (iterator *OrderedIterator[TKey, TValue]) PrevN(n int) {
	for i := 0; i < n; i++ {
		iterator.Prev()
	}
}

func (iterator *OrderedIterator[TKey, TValue]) MoveBy(n int) {
	if n > 0 {
		iterator.PrevN(-n)

	} else {
		iterator.NextN(n)
	}
}

// https://www.geeksforgeeks.org/find-distance-between-two-nodes-of-a-binary-tree/
func (iterator *OrderedIterator[TKey, TValue]) MoveTo(key TKey) (found bool) {
	if iterator.IsValid() && iterator.tree.Comparator(key, iterator.node.Key) == 0 {
		return true
	}

	targetNode, distance := iterator.tree.lookup(key)
	if targetNode == nil {
		return false
	}

	if !iterator.IsValid() {

	}

	cmp := iterator.tree.Comparator(key, iterator.node.Key)

	_, distance = iterator.tree.lookupFrom(iterator.node, targetNode.Key)
	if distance > 0 {
		if cmp > 0 {
			iterator.index += distance
		} else {
			iterator.index -= distance
		}
	}

	_, distance = iterator.tree.lookupFrom(targetNode, iterator.node.Key)
	if cmp > 0 {
		iterator.index += distance
	} else {
		iterator.index -= distance
	}

	return true
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *OrderedIterator[TKey, TValue]) Get() (value TValue, found bool) {
	if iterator.node == nil {
		return
	}

	return iterator.node.Value, true
}

// Key returns the current element's key.
// Does not modify the state of the iterator.
func (iterator *OrderedIterator[TKey, TValue]) Index() (key TKey, found bool) {
	if iterator.node == nil {
		return
	}

	return iterator.node.Key, true
}

// Node returns the current element's node.
// Does not modify the state of the iterator.
func (iterator *OrderedIterator[TKey, TValue]) Node() *Node[TKey, TValue] {
	return iterator.node
}
