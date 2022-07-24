// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedliststack

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/lists/singlylinkedlist"
)

// Assert Iterator implementation.
var _ ds.ReadWriteOrdCompForRandCollIterator[int, any] = (*Iterator[any])(nil)

type Iterator[T any] struct {
	*singlylinkedlist.Iterator[T]
}

// NewIterator returns a stateful iterator whose values can be fetched by an index.
func (list *Stack[T]) NewIterator(list_ *Stack[T], index int) *Iterator[T] {
	return &Iterator[T]{list_.list.NewIterator(list_.list, index)}
}

// DistanceTo implements ds.ReadWriteOrdCompBidRandCollIterator
// If other is of type IndexedIterator, IndexedIterator.Index() will be used, possibly executing in O(1)
func (it *Iterator[T]) DistanceTo(other ds.OrderedIterator) int {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	thisIndex, _ := it.Index()
	otherThisIndex, _ := otherThis.Index()

	return thisIndex - otherThisIndex
}

// IsAfter implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsAfter(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) > 0
}

// IsBefore implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsBefore(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) < 0
}

// IsEqual implements ds.ReadWriteOrdCompBidRandCollIterator
func (it *Iterator[T]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) == 0
}

// NOTE: Turns out that struct embedings remove the need for the redundant implementation below.
// In case it is needed again, it will probably stay here

// // Iterator holding the iterator's state
// type Iterator[T any] struct {
// 	stack   *Stack[T]
// 	index   int
// 	element *singlylinkedlist.Element[T]
// 	// Redundant but has better locality
// 	value T
// 	size  int
// }

// // NewIterator returns a stateful iterator whose values can be fetched by an index.
// func (list *Stack[T]) NewIterator(list_ *Stack[T], index int) *Iterator[T] {
// 	it := &Iterator[T]{stack: list_, index: index, size: list.Size()}

// 	it.element, _ = it.stack.list.First().Get()

// 	for i := 0; i < index; i++ {
// 		it.element = it.element.Next
// 	}

// 	return it
// }

// // IsValid implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) IsValid() bool {
// 	return it.size > 0 && !it.IsBegin() && !it.IsEnd()
// }

// // Get implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) Get() (value T, found bool) {
// 	return it.value, it.IsValid()
// }

// // Set implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) Set(value T) bool {
// 	if !it.IsValid() {
// 		return false
// 	}

// 	it.element.Value = value
// 	it.value = value

// 	return true
// }

// // Size implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) Size() int {
// 	return it.size
// }

// // Index implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) Index() (int, bool) {
// 	return it.index, true
// }

// // Next implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) Next() bool {
// 	it.index = utils.Min(it.index+1, it.size)

// 	if !it.IsValid() {
// 		return false
// 	}

// 	it.element = it.element.Next
// 	it.value = it.element.Value

// 	return true
// }

// // NextN implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) NextN(n int) bool {
// 	it.index = utils.Min(it.index+n, it.size)

// 	if !it.IsValid() {
// 		return false
// 	}

// 	for i := 0; i < n; i++ {
// 		it.element = it.element.Next
// 	}

// 	it.value = it.element.Value

// 	return true
// }

// // Previous implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) Previous() bool {
// 	it.index = utils.Max(it.index-1, -1)

// 	if !it.IsValid() {
// 		return false
// 	}

// 	it.value, _ = it.stack.list.Get(it.index)

// 	return true
// }

// // PreviousN implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) PreviousN(n int) bool {
// 	it.index = utils.Max(it.index-n, -1)

// 	if !it.IsValid() {
// 		return false
// 	}

// 	it.value, _ = it.stack.list.Get(it.index)

// 	return true
// }

// // MoveBy implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) MoveBy(n int) bool {
// 	if n > 0 {
// 		return it.NextN(n)
// 	}

// 	return it.PreviousN(-n)
// }

// // MoveTo implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) MoveTo(i int) bool {
// 	return it.MoveBy(i - it.index)
// }

// // IsBegin implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) IsBegin() bool {
// 	return it.index == -1
// }

// // IsEnd implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) IsEnd() bool {
// 	return it.size == 0 || it.index == it.size
// }

// // IsFirst implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) IsFirst() bool {
// 	return it.index == 0
// }

// // IsLast implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) IsLast() bool {
// 	return it.index == it.size-1
// }

// // GetAt implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) GetAt(i int) (value T, found bool) {
// 	if !it.IsValid() {
// 		return
// 	}

// 	return it.stack.list.Get(i)
// }

// // SetAt implements ds.ReadWriteOrdCompBidRandCollIterator
// func (it *Iterator[T]) SetAt(i int, value T) bool {
// 	if !it.IsValid() {
// 		return false
// 	}
// 	it.stack.list.Set(i, value)

// 	return true
// }
