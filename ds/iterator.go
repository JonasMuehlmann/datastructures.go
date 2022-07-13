// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ds

// TODO: Implement Constructors for types, which take iterators to materialize

// Iterator defines the minimum functionality required for all other iterators.
type Iterator interface {
	// IsBegin checks if the iterator is pointing to one element before it's first element, unless Next() is called, the iterator is in an invalid state.
	IsBegin() bool
	// IsEnd checks if the iterator is pointing to one element past it's last element, unless Previous() is called, the iterator is in an invalid state.
	IsEnd() bool

	// IsFirst checks if the iterator is pointing to it's first element, when Next() is called, the iterator is in an invalid state.
	IsFirst() bool
	// IsBegin checks if the iterator is pointing to it's last element, when Previous() is called, the iterator is in an invalid state.
	IsLast() bool

	// IsValid checks if the iterator is in a valid position and not e.g. out of bounds.
	IsValid() bool
}

// OrderedIterator defines an Iterator, which can be said to be in a position before or after others's position.
type OrderedIterator interface {
	// *********************    Inherited methods    ********************//
	Iterator
	// ************************    Own methods    ***********************//

	// IsBefore checks if the iterator's position is said to come before other's position.
	IsBefore(other OrderedIterator) bool
	// IsBefore checks if the iterator's position is said to come after other's position.
	IsAfter(other OrderedIterator) bool

	// DistanceTo returns the signed distance of the iterator's position to other's position.
	DistanceTo(other OrderedIterator) int
}

// ComparableIterator defines an Iterator, which can be said to be in a position equal to other's position.
type ComparableIterator interface {
	// *********************    Inherited methods    ********************//
	Iterator
	// ************************    Own methods    ***********************//

	// IsEqualTo checks if the iterator's position is said to be equal to other's position.
	IsEqual(other ComparableIterator) bool
}

// SizedIterator defines an Iterator, which can be said to have a fixed size.
type SizedIterator interface {
	// *********************    Inherited methods    ********************//
	Iterator
	// ************************    Own methods    ***********************//

	// Size returns the number of elements in the iterator.
	Size() int
}

// CollectionIterator defines a SizedIterator, which can be said to reference a collection of elements.
type CollectionIterator[TIndex any] interface {
	// *********************    Inherited methods    ********************//
	SizedIterator
	// ************************    Own methods    ***********************//

	// Index returns the index of the iterator's position in the collection.
	Index() TIndex
}

// WritableIterator defines an Iterator, which can be used to write to the underlying values.
type WritableIterator[T any] interface {
	// *********************    Inherited methods    ********************//
	Iterator
	// ************************    Own methods    ***********************//

	// Set sets the value at the iterator's position.
	Set(value T) bool
}

// ReadableIterator defines an Iterator, which can be used to read the underlying values.
type ReadableIterator[T any] interface {
	// *********************    Inherited methods    ********************//
	Iterator
	// ************************    Own methods    ***********************//

	// Get returns the value at the iterator's position.
	// found will be false if the iterator is in an invalid state or the collection is empty.
	Get() (value T, found bool)
}

// ForwardIterator defines an ordered Iterator, which can be moved forward according to the indexes ordering.
type ForwardIterator interface {
	// *********************    Inherited methods    ********************//
	Iterator
	// ************************    Own methods    ***********************//

	// Next moves the iterator forward by one position.
	Next()

	// NextN moves the iterator forward by n positions.
	NextN(i int)

	// Advance() bool
	// Next() ForwardIterator

	// AdvanceN(n int) bool
	// NextN(n int) ForwardIterator
}

// UnorderedForwardIterator defines an unordered ForwardIterator, which can be moved forward according to  the indexes ordering.
// This iterator would allow you to e.g. iterate over a builtin map in the lexographical order of the keys.
type UnorderedForwardIterator interface {
	// *********************    Inherited methods    ********************//
	ForwardIterator
	// ************************    Own methods    ***********************//
}

// ReversedIterator defines a a ForwardIterator, whose iteration direction is reversed.
// This allows using ForwardIterator and ReversedIterator with the same API.
type ReversedIterator interface {
	// *********************    Inherited methods    ********************//
	ForwardIterator
	// ************************    Own methods    ***********************//
}

// UnorderedReversedIterator defines an UnorderedForwardIterator, whose iteration direction is reversed.
// This allows using UnorderedForwardIterator and UnorderedReversedIterator with the same API.
type UnorderedReversedIterator interface {
	// *********************    Inherited methods    ********************//
	UnorderedForwardIterator
	// ************************    Own methods    ***********************//
}

// BackwardIterator defines an Iterator, which can be moved backward according to the indexes ordering.
type BackwardIterator interface {
	// *********************    Inherited methods    ********************//
	Iterator
	// ************************    Own methods    ***********************//

	// Next moves the iterator backward  by one position.
	Previous()

	// NextN moves the iterator backward by n positions.
	PreviousN(n int)

	// Recede() bool
	// Previous() BackwardIterator

	// RecedeN(n int) bool
	// PreviousN(n int) BackwardIterator
}

// UnorderedBackwardIterator defines an unordered BackwardIterator, which can be moved backward according to  the indexes ordering.
// This iterator would allow you to e.g. iterate over a builtin map in the reverse lexographical order of the keys.
type UnorderedBackwardIterator interface {
	// *********************    Inherited methods    ********************//
	BackwardIterator
	// ************************    Own methods    ***********************//
}

// BidirectionalIterator defines a ForwardIterator and BackwardIterator, which can be moved forward and backward according to the underlying data structure's ordering.
type BidirectionalIterator interface {
	// *********************    Inherited methods    ********************//
	ForwardIterator
	BackwardIterator
	// ************************    Own methods    ***********************//

	// Next moves the iterator forward/backward by n positions.
	MoveBy(n int)
	// Nth(n int) BidirectionalIterator
}

// UnorderedBidirectionalIterator defines an UnorderedForwardIterator and UnorderedBackwardIterator, which can be moved forward and backward according to the indexes ordering.
type UnorderedBidirectionalIterator interface {
	// *********************    Inherited methods    ********************//
	UnorderedForwardIterator
	UnorderedBackwardIterator
	// ************************    Own methods    ***********************//

	// Next moves the iterator forward/backward by n positions.
	MoveBy(n int)
	// Nth(n int) BidirectionalIterator
}

// RandomAccessIterator defines a BidirectionalIterator and CollectionIterator, which can be moved to every position in the iterator.
type RandomAccessIterator[TIndex any] interface {
	// *********************    Inherited methods    ********************//
	BidirectionalIterator
	CollectionIterator[TIndex]
	// ************************    Own methods    ***********************//

	// MoveTo moves the iterator to the given index.
	MoveTo(i TIndex)
}

// UnorderedRandomAccessIterator defines an UnorderedBidirectionalIterator and CollectionIterator, which can be moved to every position in the iterator.
type UnorderedRandomAccessIterator[TIndex any] interface {
	// *********************    Inherited methods    ********************//
	UnorderedBidirectionalIterator
	CollectionIterator[TIndex]
	// ************************    Own methods    ***********************//

	// MoveTo moves the iterator to the given index.
	MoveTo(i TIndex)
}

// RandomAccessReadableIterator defines a RandomAccessIterator and ReadableIterator, which can read from every index in the iterator.
type RandomAccessReadableIterator[T any, V any] interface {
	// *********************    Inherited methods    ********************//
	RandomAccessIterator[T]
	ReadableIterator[V]
	// ************************    Own methods    ***********************//

	// GetAt returns the value at the given index of the iterator.
	GetAt(i T) (value V, found bool)
}

// RandomAccessWriteableIterator defines a RandomAccessIterator and WritableIterator, which can write from every index in the iterator.
type RandomAccessWriteableIterator[T any, V any] interface {
	// *********************    Inherited methods    ********************//
	RandomAccessIterator[T]
	WritableIterator[V]
	// ************************    Own methods    ***********************//

	// GetAt sets the value at the given index of the iterator.
	SetAt(i T, value V) bool
}

// UnorderedRandomAccessReadableIterator defines a RandomAccessIterator and ReadableIterator, which can read from every index in the iterator.
type UnorderedRandomAccessReadableIterator[T any, V any] interface {
	// *********************    Inherited methods    ********************//
	UnorderedRandomAccessIterator[T]
	ReadableIterator[V]
	// ************************    Own methods    ***********************//

	// GetAt returns the value at the given index of the iterator.
	GetAt(i T) (value V, found bool)
}

// UnorderedRandomAccessWriteableIterator defines a RandomAccessIterator and WritableIterator, which can write from every index in the iterator.
type UnorderedRandomAccessWriteableIterator[T any, V any] interface {
	// *********************    Inherited methods    ********************//
	UnorderedRandomAccessIterator[T]
	WritableIterator[V]
	// ************************    Own methods    ***********************//

	// GetAt sets the value at the given index of the iterator.
	SetAt(i T, value V) bool
}
