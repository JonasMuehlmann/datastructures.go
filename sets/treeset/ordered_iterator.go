// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treeset

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/trees/redblacktree"
)

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[int, string] = (*OrderedIterator[string])(nil)

// Iterator holding the iterator's state
type OrderedIterator[T comparable] struct {
	*redblacktree.OrderedIterator[T, struct{}]
	set *Set[T]
}

// NewIterator returns a stateful iterator whose values can be fetched by an index.
func (set *Set[T]) NewOrderedIterator(index int, size int) *OrderedIterator[T] {
	return &OrderedIterator[T]{set.tree.NewOrderedIterator(index, size), set}
}

// NOTE: The following methods need to be reimplemented because of the type assertions they contain

// If other is of type IndexedIterator, IndexedIterator.Index() will be used, possibly executing in O(1)
func (it *OrderedIterator[T]) DistanceTo(other ds.OrderedIterator) int {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	// thisIndex, _ := it.Index()
	// otherThisIndex, _ := otherThis.Index()

	// return thisIndex - otherThisIndex

	return it.OrderedIterator.DistanceTo(otherThis.OrderedIterator)
}

func (it *OrderedIterator[T]) IsAfter(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) > 0
}

func (it *OrderedIterator[T]) IsBefore(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) < 0
}

func (it *OrderedIterator[T]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) == 0
}

// PERF: These methods are inefficient, but the API is limiting here
func (it *OrderedIterator[T]) Get() (value T, found bool) {
	return it.OrderedIterator.GetKey()
}

func (it *OrderedIterator[T]) GetKey() (value int, found bool) {
	return it.Index()
}

func (it *OrderedIterator[T]) Set(value T) bool {
	if !it.IsValid() {
		return false
	}

	curKey, _ := it.OrderedIterator.GetKey()
	// FIX: This probably does not change the cached value in the tree iterator
	it.set.tree.Remove(curKey)
	it.set.tree.Put(value, struct{}{})

	return true
}

func (it *OrderedIterator[T]) GetAt(i int) (value T, found bool) {
	treeIteratorCopy := it.set.tree.OrderedFirst()
	valid := treeIteratorCopy.MoveBy(i)
	if !valid {
		found = false

		return
	}

	return treeIteratorCopy.GetKey()
}

func (it *OrderedIterator[T]) GetAtKey(i int) (value T, found bool) {
	return it.GetAt(i)
}

func (it *OrderedIterator[T]) SetAt(i int, value T) bool {
	treeIteratorCopy := it.set.tree.OrderedFirst()
	valid := treeIteratorCopy.MoveBy(i)
	if !valid {
		return false
	}

	keyToRemove, _ := treeIteratorCopy.GetKey()

	// FIX: This probably does not change the cached value in the tree iterator
	it.set.tree.Remove(keyToRemove)
	it.set.tree.Put(value, struct{}{})

	return true
}

func (it *OrderedIterator[T]) SetAtKey(i int, value T) bool {
	return it.SetAt(i, value)
}

func (it *OrderedIterator[T]) Index() (value int, found bool) {
	treeIteratorCopy := it.set.tree.OrderedFirst()

	return treeIteratorCopy.DistanceTo(it.OrderedIterator), it.IsValid()
}

func (it *OrderedIterator[T]) MoveTo(i int) bool {
	treeIteratorCopy := it.set.tree.OrderedFirst()
	treeIteratorCopy.MoveBy(i)

	wantedKey, valid := treeIteratorCopy.Index()
	if !valid {
		return false
	}

	it.OrderedIterator.MoveTo(wantedKey)

	return it.IsValid()
}

func (it *OrderedIterator[T]) MoveToKey(i int) bool {
	return it.MoveTo(i)
}
