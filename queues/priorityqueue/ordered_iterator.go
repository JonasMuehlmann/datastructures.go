// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package priorityqueue

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/trees/binaryheap"
)

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[int, any] = (*OrderedIterator[any])(nil)

// Iterator holding the iterator's state
type OrderedIterator[T any] struct {
	*binaryheap.OrderedIterator[T]
}

// NewIterator returns a stateful iterator whose values can be fetched by an index.
func (list *Queue[T]) NewOrderedIterator(q *Queue[T], index int) *OrderedIterator[T] {
	return &OrderedIterator[T]{q.heap.NewOrderedIterator(q.heap, index)}
}

// DistanceTo implements ds.ReadWriteOrdCompBidRandCollIterator
// If other is of type IndexedIterator, IndexedIterator.Index() will be used, possibly executing in O(1)
func (it *OrderedIterator[T]) DistanceTo(other ds.OrderedIterator) int {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	thisIndex, _ := it.Index()
	otherThisIndex, _ := otherThis.Index()

	return thisIndex - otherThisIndex
}

// IsAfter implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) IsAfter(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) > 0
}

// IsBefore implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) IsBefore(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) < 0
}

// IsEqual implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *OrderedIterator[T]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*OrderedIterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) == 0
}
