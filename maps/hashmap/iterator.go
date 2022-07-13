// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashmap

import (
	"fmt"

	"github.com/JonasMuehlmann/datastructures.go/ds"
)

// Assert Iterator implementation
var _ ds.ReadWriteCompForRandCollIterator[string, any] = (*Iterator[string, any])(nil)

type hashMapItState int

type Iterator[TKey comparable, TValue any] struct {
	m                   *Map[TKey, TValue]
	key                 TKey
	completedIterations int

	skip1   chan struct{}
	skipN   chan int
	skipTo  chan TKey
	canRead chan struct{}
}

func (it *Iterator[TKey, TValue]) iterate() {
	var iterationsToSkip int
	var keyToSkipTo TKey
	var isSkippingToKey bool

	for k := range it.m.m {
		if isSkippingToKey && k != keyToSkipTo {
			fmt.Println("skipTo not found")
			continue
		} else if isSkippingToKey && k == keyToSkipTo {
			fmt.Println("skipTo found")
			isSkippingToKey = false
			it.key = k
			it.canRead <- struct{}{}
			continue
		} else if iterationsToSkip > 0 {
			fmt.Println("skipN not found")
			iterationsToSkip--
			continue
		} else if iterationsToSkip == -1 {
			fmt.Println("skipN found")
			it.key = k
			it.canRead <- struct{}{}
			continue
		}

		select {
		case <-it.skip1:
			fmt.Println("skip1")
			it.key = k
			it.canRead <- struct{}{}
			continue
		case iterationsToSkip = <-it.skipN:
			fmt.Println("skipN")
			continue
		case keyToSkipTo = <-it.skipTo:
			fmt.Println("skipTo")
			isSkippingToKey = true
			continue
		}

		it.completedIterations++
	}

	close(it.skip1)
	close(it.skipN)
	close(it.skipTo)
	close(it.canRead)
}

func (m *Map[TKey, TValue]) NewIterator(m_ *Map[TKey, TValue]) *Iterator[TKey, TValue] {
	it := &Iterator[TKey, TValue]{
		m:       m_,
		skip1:   make(chan struct{}),
		canRead: make(chan struct{}),
		skipN:   make(chan int),
		skipTo:  make(chan TKey),
	}

	for k := range m_.m {
		it.key = k
		break
	}

	go it.iterate()

	return it
}

// IsFirst implements ds.ReadWriteCompForRandCollIterator
func (it *Iterator[TKey, TValue]) IsFirst() bool {
	return it.completedIterations == 0
}

// IsLast implements ds.ReadWriteCompForRandCollIterator
func (it *Iterator[TKey, TValue]) IsLast() bool {
	return it.completedIterations == it.m.Size()
}

// IsBegin implements ds.ReadWriteCompForRandCollIterator
func (it *Iterator[TKey, TValue]) IsBegin() bool {
	return false
}

// IsEnd implements ds.ReadWriteCompForRandCollIterator
func (it *Iterator[TKey, TValue]) IsEnd() bool {
	return false
}

// IsValid implements ds.ReadWriteCompForRandCollIterator
func (it *Iterator[TKey, TValue]) IsValid() bool {
	return it.completedIterations < it.m.Size()
}

// IsEqual implements ds.ReadWriteCompForRandCollIterator
func (it *Iterator[TKey, TValue]) IsEqual(other ds.ComparableIterator) bool {
	otherThis, ok := other.(*Iterator[TKey, TValue])
	if !ok {
		panic(ds.CanOnlyCompareEqualIteratorTypes)
	}

	return it.key == otherThis.key
}

// Size implements ds.ReadWriteCompForRandCollIterator
func (it *Iterator[TKey, TValue]) Size() int {
	return it.m.Size()
}

// Index implements ds.ReadWriteCompForRandCollIterator
func (it *Iterator[TKey, TValue]) Index() TKey {
	return it.key
}

// Next implements ds.ReadWriteCompForRandCollIterator
func (it *Iterator[TKey, TValue]) Next() {
	if it.completedIterations == it.m.Size() {
		return
	}

	it.skip1 <- struct{}{}
	<-it.canRead
}

// NextN implements ds.ReadWriteCompForRandCollIterator
func (it *Iterator[TKey, TValue]) NextN(i int) {
	if it.completedIterations == it.m.Size() {
		return
	}

	it.skipN <- i
	<-it.canRead
}

// MoveTo implements ds.ReadWriteCompForRandCollIterator
func (it *Iterator[TKey, TValue]) MoveTo(i TKey) {
	// if it.completedIterations == it.m.Size() {
	// 	return
	// }

	// it.skipTo <- i
	// <-it.canRead

	it.key = i
}

// Get implements ds.ReadWriteCompForRandCollIterator
func (it *Iterator[TKey, TValue]) Get() (value TValue, found bool) {
	return it.m.Get(it.key)
}

// GetAt implements ds.ReadWriteCompForRandCollIterator
func (it *Iterator[TKey, TValue]) GetAt(i TKey) (value TValue, found bool) {
	return it.m.Get(i)
}

// Set implements ds.ReadWriteCompForRandCollIterator
func (it *Iterator[TKey, TValue]) Set(value TValue) bool {
	it.m.Put(it.key, value)

	return true
}

// SetAt implements ds.ReadWriteCompForRandCollIterator
func (it *Iterator[TKey, TValue]) SetAt(i TKey, value TValue) bool {
	it.m.Put(it.key, value)

	return true
}
