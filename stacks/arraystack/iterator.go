// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arraystack

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/lists/arraylist"
)

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollIterator[int, any] = (*Iterator[any])(nil)

// Iterator holding the iterator's state
type Iterator[T any] struct {
	*arraylist.Iterator[T]
}

// NewIterator returns a stateful iterator whose values can be fetched by an index.
func (list *Stack[T]) NewIterator(index int) *Iterator[T] {
	return &Iterator[T]{list.list.NewIterator(index)}
}

// NOTE: The following methods need to be reimplemented because of the type assertions they contain

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

func (it *Iterator[T]) IsAfter(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) > 0
}

func (it *Iterator[T]) IsBefore(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*Iterator[T])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) < 0
}

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
// 	stack *Stack[T]
// 	index int
// }

// // NewIterator returns a stateful iterator whose values can be fetched by an index.
// func (list *Stack[T]) NewIterator(list_ *Stack[T], index int) *Iterator[T] {
// 	return &Iterator[T]{stack: list_, index: index}
// }

// func (it *Iterator[T]) IsValid() bool {
// 	return it.stack.withinRange(it.index)
// }

// func (it *Iterator[T]) Get() (value T, found bool) {
// 	if it.stack.Size() == 0 || !it.IsValid() {
// 		return
// 	}

// 	return it.stack.list.Get(it.index)
// }

// func (it *Iterator[T]) Set(value T) bool {
// 	if it.stack.Size() == 0 || !it.IsValid() {
// 		return false
// 	}

// 	it.stack.list.Set(it.index, value)

// 	return true
// }

// func (it *Iterator[T]) Next() {
// 	it.index++
// }

// func (it *Iterator[T]) NextN(i int) {
// 	it.index += i
// }

// func (it *Iterator[T]) Previous() {
// 	it.index--
// }

// func (it *Iterator[T]) PreviousN(n int) {
// 	it.index -= n
// }

// func (it *Iterator[T]) MoveBy(n int) {
// 	it.index += n
// }

// func (it *Iterator[T]) Size() int {
// 	return it.stack.Size()
// }

// func (it *Iterator[T]) Index() (int, bool) {
// 	return it.index, true
// }

// func (it *Iterator[T]) MoveTo(i int) bool {
// 	it.index = i

// 	return true
// }

// func (it *Iterator[T]) IsBegin() bool {
// 	return it.index == -1
// }

// func (it *Iterator[T]) IsEnd() bool {
// 	return it.stack.Size() == 0 || it.index == it.stack.Size()
// }

// func (it *Iterator[T]) IsFirst() bool {
// 	return it.index == 0
// }

// func (it *Iterator[T]) IsLast() bool {
// 	return it.index == it.stack.Size()-1
// }

// func (it *Iterator[T]) GetAt(i int) (value T, found bool) {
// 	if it.stack.Size() == 0 || !it.stack.withinRange(i) {
// 		return
// 	}

// 	return it.stack.list.Get(i)
// }

// func (it *Iterator[T]) SetAt(i int, value T) bool {
// 	if it.stack.Size() == 0 || !it.stack.withinRange(i) {
// 		return false
// 	}
// 	it.stack.list.Set(i, value)

// 	return true
// }
