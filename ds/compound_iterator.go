// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
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
// MappingIterator = Map
// CollectionIterator = Coll

const (
	CanOnlyCompareEqualIteratorTypes = "Can only compare iterators of equal concrete type"
)

type ReadWriteCompForRandCollIterator[TIndex any, TValue any] interface {
	ComparableIterator
	CollectionIterator
	ForwardIterator
	RandomAccessReadableIterator[TIndex, TValue]
	RandomAccessWriteableIterator[TIndex, TValue]
}

type ReadWriteOrdCompForRandCollIterator[TIndex any, TValue any] interface {
	OrderedIterator
	ComparableIterator
	ForwardIterator
	ReadableIterator[TValue]
	WritableIterator[TValue]
	RandomAccessIterator
}

type ReadWriteOrdCompBidRandCollIterator[TIndex any, TValue any] interface {
	ComparableIterator
	OrderedIterator
	BidirectionalIterator
	RandomAccessReadableIterator[TIndex, TValue]
	RandomAccessWriteableIterator[TIndex, TValue]
}

type ReadWriteOrdCompBidRandCollMapIterator[TIndex any, TValue any] interface {
	ComparableIterator
	OrderedIterator
	BidirectionalIterator
	RandomAccessReadableMappingIterator[TIndex, TValue]
	RandomAccessWriteableMappingIterator[TIndex, TValue]
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

type ReadCompForIndexIterator[TIndex any, TValue any] interface {
	ReadableIterator[TValue]
	IndexedIterator
	ComparableIterator
	ForwardIterator
}

type ReadCompForIndexMapIterator[TIndex any, TValue any] interface {
	ReadableIterator[TValue]
	IndexedIterator
	MappingIterator[TIndex]
	ComparableIterator
	ForwardIterator
}

type CompIndexIterator[TIndex any] interface {
	IndexedIterator
	ComparableIterator
}

type CompIndexMapIterator[TIndex any] interface {
	IndexedIterator
	ComparableIterator
}

type ReadCompIterator[TValue any] interface {
	ReadableIterator[TValue]
	ComparableIterator
}
