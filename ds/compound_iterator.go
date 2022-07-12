// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ds

// NOTE: Abbreviations:
// OrderedIterator = Ord
// UnorderedIterator =Unord
// ComparableIterator = Comp
// CollectionIterator = Coll
// WritableIterator = W
// ReadableIterator = R
// ForwardsIterator = For
// BackwardIterator = Back
// ReverseIterator = Rev
// BidirectionalIterator = Bid
// RandomAccessIterator = Rand

type RWOrdCompBidRandCollIterator[T any, TIndex any] interface {
	OrderedIterator
	ComparableIterator

	BidirectionalIterator

	RandomAccessReadableIterator[TIndex, T]
	RandomAccessWriteableIterator[TIndex, T]
}

type RWOrdCompBidRevRandCollIterator[T any, TIndex any] interface {
	OrderedIterator
	ComparableIterator

	ReversedIterator
	BackwardIterator

	RandomAccessReadableIterator[TIndex, T]
	RandomAccessWriteableIterator[TIndex, T]
}

const (
	CanOnlyCompareEqualIteratorTypes = "Can only compare iterators of equal concrete type"
)
