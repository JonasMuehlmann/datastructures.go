// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package treemap

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/trees/redblacktree"
)

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollMapIterator[string, any] = (*OrderedIterator[string, any])(nil)

// Iterator holding the iterator's state
type OrderedIterator[TKey comparable, TValue any] struct {
	*redblacktree.OrderedIterator[TKey, TValue]
}

// NewIterator returns a stateful iterator whose values can be fetched by an index.
func (list *Map[TKey, TValue]) NewOrderedIterator(index int) *OrderedIterator[TKey, TValue] {
	return &OrderedIterator[TKey, TValue]{list.tree.NewOrderedIterator(index)}
}

// NOTE: The following methods need to be reimplemented because of the type assertions they contain

// If other is of type IndexedIterator, IndexedIterator.Index() will be used, possibly executing in O(1)
func (it *OrderedIterator[TKey, TValue]) DistanceTo(other ds.OrderedIterator) int {
	otherThis, ok := other.(*OrderedIterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.OrderedIterator.DistanceTo(otherThis.OrderedIterator)
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

func (it *OrderedIterator[TKey, TValue]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*OrderedIterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) == 0
}
