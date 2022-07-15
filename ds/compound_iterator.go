// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ds

// NOTE: Abbreviations and order:
// ReadableIterator = Read
// WritableIterator = Write

// OrderedIterator = Ord
// UnorderedIterator =Unord

// ComparableIterator = Comp

// ForwardsIterator = For
// BackwardIterator = Back
// ReverseIterator = Rev
// BidirectionalIterator = Bid

// RandomAccessIterator = Rand
// CollectionIterator = Coll

const (
	CanOnlyCompareEqualIteratorTypes = "Can only compare iterators of equal concrete type"
)

type ReadWriteOrdCompBidRandCollIterator[TIndex any, TValue any] interface {
	OrderedIterator
	ComparableIterator

	BidirectionalIterator

	RandomAccessReadableIterator[TIndex, TValue]
	RandomAccessWriteableIterator[TIndex, TValue]
}

type ReadWriteCompForRandCollIterator[TIndex any, TValue any] interface {
	ComparableIterator
	CollectionIterator[TIndex]
	ForwardIterator
	RandomAccessReadableIterator[TIndex, TValue]
	RandomAccessWriteableIterator[TIndex, TValue]
}

type ReadWriteUnordCompBidRandCollIterator[TIndex any, TValue any] interface {
	ComparableIterator
	OrderedIterator
	CollectionIterator[TIndex]
	BidirectionalIterator
	RandomAccessReadableIterator[TIndex, TValue]
	RandomAccessWriteableIterator[TIndex, TValue]
}
