// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ds

// TODO: Implement Constructors for types, which take iterators to materialize

// Iterator defines the minimum functionality required for all other iterators.
type Iterator interface {
	// Begin returns an initialized iterator, which points to one element before it's first.
	// Unless Next() is called, the iterator is in an invalid state.
	Begin() Iterator
	// End returns an initialized iterator, which points to one element afrer it's last.
	// Unless Previous() is called, the iterator is in an invalid state.
	End() Iterator

	// IsBegin checks if the iterator is pointing to one element before it's first element, unless Next() is called, the iterator is in an invalid state.
	IsBegin() bool
	// IsEnd checks if the iterator is pointing to one element past it's last element, unless Previous() is called, the iterator is in an invalid state.
	IsEnd() bool

	// First returns an initialized iterator, which points to it's first element.
	First() Iterator
	// Last returns an initialized iterator, which points to it's last element.
	Last() Iterator

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
type SizedIterator[T any] interface {
	// *********************    Inherited methods    ********************//
	Iterator
	// ************************    Own methods    ***********************//

	// Size returns the number of elements in the iterator.
	Size() int
}

// CollectionIterator defines a SizedIterator, which can be said to reference a collection of elements.
type CollectionIterator[T any] interface {
	// *********************    Inherited methods    ********************//
	SizedIterator[T]
	// ************************    Own methods    ***********************//

	// Index returns the index of the iterator's position in the collection.
	Index() T
}

// WritableIterator defines an Iterator, which can be used to write to the underlying values.
type WritableIterator[T any] interface {
	// *********************    Inherited methods    ********************//
	Iterator
	// ************************    Own methods    ***********************//

	// Set sets the value at the iterator's position.
	Set(value T)
}

// ReadableIterator defines an Iterator, which can be used to read the underlying values.
type ReadableIterator[T any] interface {
	// *********************    Inherited methods    ********************//
	Iterator
	// ************************    Own methods    ***********************//

	// Get returns the value at the iterator's position.
	Get() T
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

// UnorderedForwardIterator defines an unordered ForwardIterator, which can be moved forward according to  the indexes ordering in addition to the underlying data structure's ordering.
// This iterator would allow you to e.g. iterate over a builtin map in the lexographical order of the keys.
type UnorderedForwardIterator interface {
	// *********************    Inherited methods    ********************//
	ForwardIterator
	// ************************    Own methods    ***********************//

	// NextOrdered moves the iterator forward by one position according to the indexes lexographical ordering instead of the underlying data structure's ordering.
	NextOrdered()

	// NextOrdered moves the iterator forward by n positions according to the indexes lexographical ordering instead of the underlying data structure's ordering.
	NextOrderedN(n int)

	// AdvanceOrdered() bool
	// NextOrdered() UnorderedForwardIterator

	// AdvanceOrderedN(n int) bool
	// NextOrderedN(n int) UnorderedForwardIterator
}

// ReversedIterator defines a a ForwardIterator, whose iteration direction is reversed.
// This allows using ForwardIterator and ReversedIterator with the same API.
type ReversedIterator interface {
	// *********************    Inherited methods    ********************//
	ForwardIterator
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

// UnorderedBackwardIterator defines an unordered BackwardIterator, which can be moved backward according to  the indexes ordering in addition to the underlying data structure's ordering.
// This iterator would allow you to e.g. iterate over a builtin map in the reverse lexographical order of the keys.
type UnorderedBackwardIterator interface {
	// *********************    Inherited methods    ********************//
	BackwardIterator
	// ************************    Own methods    ***********************//

	// PreviousOrdered moves the iterator backward by one position according to the indexes lexographical ordering instead of the underlying data structure's ordering.
	PreviousOrdered()

	// PreviousOrdered moves the iterator backward by n positions according to the indexes lexographical ordering instead of the underlying data structure's ordering.
	PreviousOrderedN(n int)

	// RecedeOrdered() bool
	// PreviousOrdered() UnorderedBackwardIterator

	// RecedeOrderedN(n int) bool
	// PreviousOrderedN(n int) UnorderedBackwardIterator
}

// BidirectionalIterator defines a ForwardIterator and BackwardIterator, which can be moved forward and backward according to the underlying data structure's ordering.
type BidirectionalIterator interface {
	// *********************    Inherited methods    ********************//
	ForwardIterator
	BackwardIterator
	// ************************    Own methods    ***********************//

	// Next moves the iterator forward/backward by n positions.
	MoveBy(n int) bool
	// Nth(n int) BidirectionalIterator
}

// RandomAccessIterator defines a BidirectionalIterator and CollectionIterator, which can be moved to every position in the iterator.
type RandomAccessIterator[T any] interface {
	// *********************    Inherited methods    ********************//
	BidirectionalIterator
	CollectionIterator[T]
	// ************************    Own methods    ***********************//

	// MoveTo moves the iterator to the given index.
	MoveTo(i T) bool
}

// RandomAccessReadableIterator defines a RandomAccessIterator and ReadableIterator, which can read from every index in the iterator.
type RandomAccessReadableIterator[T any, V any] interface {
	// *********************    Inherited methods    ********************//
	RandomAccessIterator[T]
	ReadableIterator[T]
	// ************************    Own methods    ***********************//

	// GetAt returns the value at the given index of the iterator.
	GetAt(i T) V
}

// RandomAccessReadableIterator defines a RandomAccessIterator and WritableIterator, which can write from every index in the iterator.
type RandomAccessWriteableIterator[T any, V any] interface {
	// *********************    Inherited methods    ********************//
	RandomAccessIterator[T]
	WritableIterator[V]
	// ************************    Own methods    ***********************//

	// GetAt sets the value at the given index of the iterator.
	SetAt(i T, value V)
}
