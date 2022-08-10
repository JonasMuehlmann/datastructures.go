// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ds

// NOTE: Abbreviations and order:
// ReadableIterator = Read
// WritableIterator = Write

// OrderedIterator = Ord
// ComparableIterator = Comp

// ForwardsIterator = For
// BackwardIterator = Back
// ReverseIterator = Rev
// BidirectionalIterator = Bid

// RandomAccessIterator = Rand
// IndexedIterator = Index
// CollectionIterator = Coll

const (
	CanOnlyCompareEqualIteratorTypes = "Can only compare iterators of equal concrete type"
)

type ReadWriteCompForRandCollIterator[TKey any, TValue any] interface {
	ComparableIterator
	CollectionIterator[TKey]
	ForwardIterator
	RandomAccessReadableIterator[TKey, TValue]
	RandomAccessWriteableIterator[TKey, TValue]
}

type ReadWriteOrdCompForRandCollIterator[TKey any, TValue any] interface {
	OrderedIterator
	ComparableIterator
	ForwardIterator
	ReadableIterator[TValue]
	WritableIterator[TValue]
	RandomAccessIterator[TKey]
}

type ReadWriteOrdCompBidRandCollIterator[TKey any, TValue any] interface {
	ComparableIterator
	OrderedIterator
	BidirectionalIterator
	RandomAccessReadableIterator[TKey, TValue]
	RandomAccessWriteableIterator[TKey, TValue]
}

type ReadForIterator[TValue any] interface {
	ReadableIterator[TValue]
	ForwardIterator
}

type ReadCompForIterator[TValue any] interface {
	ReadableIterator[TValue]
	ComparableIterator
	ForwardIterator
}

type ReadCompForIndexIterator[TKey any, TValue any] interface {
	ReadableIterator[TValue]
	IndexedIterator[TKey]
	ComparableIterator
	ForwardIterator
}

type CompIndexIterator[TKey any] interface {
	IndexedIterator[TKey]
	ComparableIterator
}

type ReadCompIterator[TValue any] interface {
	ReadableIterator[TValue]
	ComparableIterator
}
