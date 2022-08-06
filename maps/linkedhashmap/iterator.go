// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashmap

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/lists/doublylinkedlist"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Iterator implementation
var _ ds.ReadWriteOrdCompBidRandCollMapIterator[string, any] = (*Iterator[string, any])(nil)

type Iterator[TKey comparable, TValue any] struct {
	s             *Map[TKey, TValue]
	orderIterator *doublylinkedlist.Iterator[TKey]
	// Redundant but has better locality
	value TValue
	key   TKey
	index int
	size  int
}

func (m *Map[TKey, TValue]) NewIterator(position int) *Iterator[TKey, TValue] {
	it := &Iterator[TKey, TValue]{
		s:             m,
		orderIterator: m.ordering.NewIterator(position),
		index:         position,
		size:          m.Size(),
	}

	it.key, _ = it.orderIterator.Get()
	it.value, _ = m.Get(it.key)

	return it
}

func (it *Iterator[TKey, TValue]) IsBegin() bool {
	return it.index == -1
}

func (it *Iterator[TKey, TValue]) IsEnd() bool {
	return it.index == it.size
}

func (it *Iterator[TKey, TValue]) IsFirst() bool {
	return it.index == 0
}

func (it *Iterator[TKey, TValue]) IsLast() bool {
	return it.index == it.size-1
}

func (it *Iterator[TKey, TValue]) IsValid() bool {
	return it.size > 0 && !it.IsBegin() && !it.IsEnd()
}

func (it *Iterator[TKey, TValue]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*Iterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.index == otherThis.index

}

func (it *Iterator[TKey, TValue]) DistanceTo(other ds.OrderedIterator) int {
	otherThis, ok := other.(*Iterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.index - otherThis.index
}

func (it *Iterator[TKey, TValue]) IsAfter(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*Iterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) > 0
}

func (it *Iterator[TKey, TValue]) IsBefore(other ds.OrderedIterator) bool {
	otherThis, ok := other.(*Iterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.DistanceTo(otherThis) < 0
}

func (it *Iterator[TKey, TValue]) Size() int {
	return it.size
}

func (it *Iterator[TKey, TValue]) Index() (index int, found bool) {
	return it.index, it.IsValid()
}

func (it *Iterator[TKey, TValue]) Next() bool {
	it.orderIterator.Next()
	it.index = utils.Min(it.index+1, it.size)

	var found bool

	it.key, found = it.orderIterator.Get()

	return found
}

func (it *Iterator[TKey, TValue]) NextN(i int) bool {
	it.orderIterator.NextN(i)
	it.index = utils.Min(it.index+i, it.size)

	var found bool

	it.key, found = it.orderIterator.Get()

	return found
}

func (it *Iterator[TKey, TValue]) Previous() bool {
	it.orderIterator.Previous()
	it.index = utils.Max(it.index-1, -1)

	var found bool

	it.key, found = it.orderIterator.Get()

	return found
}

func (it *Iterator[TKey, TValue]) PreviousN(n int) bool {
	it.orderIterator.PreviousN(n)
	it.index = utils.Max(it.index-n, -1)

	var found bool

	it.key, found = it.orderIterator.Get()

	return found
}

func (it *Iterator[TKey, TValue]) MoveBy(n int) bool {
	if n > 0 {
		return it.NextN(n)
	} else if n < 0 {
		return it.PreviousN(-n)
	}

	return it.IsValid()
}

func (it *Iterator[TKey, TValue]) MoveToKey(key TKey) bool {
	// PERF: This can be optimized
	for it.Next() {
		if it.key == key {
			return true
		}
	}
	for it.Previous() {
		if it.key == key {
			return true
		}
	}

	return false
}

func (it *Iterator[TKey, TValue]) Get() (value TValue, found bool) {
	return it.s.Get(it.key)
}

func (it *Iterator[TKey, TValue]) GetKey() (index TKey, found bool) {
	return it.key, it.IsValid()
}

func (it *Iterator[TKey, TValue]) Set(value TValue) bool {
	if !it.IsValid() {
		return false
	}

	it.value = value
	it.s.Put(it.key, value)

	return true
}

func (it *Iterator[TKey, TValue]) GetAt(i int) (value TValue, found bool) {
	if !it.IsValid() {
		return
	}

	return it.s.NewIterator(i).Get()
}

func (it *Iterator[TKey, TValue]) SetAt(i int, value TValue) bool {
	if !it.IsValid() {
		return false
	}

	return it.s.NewIterator(i).Set(value)
}

func (it *Iterator[TKey, TValue]) GetAtKey(key TKey) (value TValue, found bool) {
	return it.s.Get(key)
}

func (it *Iterator[TKey, TValue]) SetAtKey(i TKey, value TValue) bool {
	it.s.Put(it.key, value)

	return true
}
