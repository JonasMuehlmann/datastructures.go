// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ds

type BoundsSentinel int

const (
	SentinelBegin BoundsSentinel = iota
	SentinelEnd
	SentinelFirst
	SentinelLast
	SentinelInRange
)

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
type CollectionIterator interface {
	// *********************    Inherited methods    ********************//
	SizedIterator
	IndexedIterator
	// ************************    Own methods    ***********************//
}

// IndexedIterator defines an Iterator, which defines an an iterator with an index.
// This iterator can be combined with a ReadableIterator to hold key-value or index-value pairs.
type IndexedIterator interface {
	// *********************    Inherited methods    ********************//
	SizedIterator
	// ************************    Own methods    ***********************//

	// Index returns the index of the iterator's position in the collection.
	Index() (int, bool)
}

// MappingIterator defines an Iterator, which defines an an iterator holding a key and a value.
// This iterator can be combined with a ReadableIterator to hold key-value or index-value pairs.
type MappingIterator[TKey any] interface {
	// *********************    Inherited methods    ********************//
	SizedIterator
	// ************************    Own methods    ***********************//

	// GetKey returns the key of the iterator's position in the collection.
	GetKey() (TKey, bool)
}

// WritableIterator defines an Iterator, which can be used to write to the underlying values.
type WritableIterator[TValue any] interface {
	// *********************    Inherited methods    ********************//
	Iterator
	// ************************    Own methods    ***********************//

	// Set sets the value at the iterator's position.
	Set(value TValue) bool
}

// ReadableIterator defines an Iterator, which can be used to read the underlying values.
type ReadableIterator[TValue any] interface {
	// *********************    Inherited methods    ********************//
	Iterator
	// ************************    Own methods    ***********************//

	// Get returns the value at the iterator's position.
	// found will be false if the iterator is in an invalid state or the collection is empty.
	Get() (value TValue, found bool)
}

// ForwardIterator defines an ordered Iterator, which can be moved forward according to the indexes ordering.
type ForwardIterator interface {
	// *********************    Inherited methods    ********************//
	Iterator
	// ************************    Own methods    ***********************//

	// Next moves the iterator forward by one position.
	Next() bool

	// NextN moves the iterator forward by n positions.
	NextN(i int) bool

	// Advance() bool
	// Next() ForwardIterator

	// AdvanceN(n int) bool
	// NextN(n int) ForwardIterator
}

// BackwardIterator defines an Iterator, which can be moved backward according to the indexes ordering.
type BackwardIterator interface {
	// *********************    Inherited methods    ********************//
	Iterator
	// ************************    Own methods    ***********************//

	// Next moves the iterator backward  by one position.
	Previous() bool

	// NextN moves the iterator backward by n positions.
	PreviousN(n int) bool

	// Recede() bool
	// Previous() BackwardIterator

	// RecedeN(n int) bool
	// PreviousN(n int) BackwardIterator
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

// RandomAccessIterator defines a CollectionIterator, which can be moved to every position in the iterable direction.
// A RandomAccessIterator does not imply bidirectional iteration.
type RandomAccessIterator interface {
	// *********************    Inherited methods    ********************//
	CollectionIterator
	// ************************    Own methods    ***********************//

	// MoveTo moves the iterator to the given index, if it is reachable.
	MoveTo(i int) bool
}

// RandomAccessReadableIterator defines a RandomAccessIterator and ReadableIterator, which can read from every index in the iterator.
type RandomAccessReadableIterator[TIndex any, TValue any] interface {
	// *********************    Inherited methods    ********************//
	RandomAccessIterator
	ReadableIterator[TValue]
	// ************************    Own methods    ***********************//

	// GetAt returns the value at the given index of the iterator.
	GetAt(i TIndex) (value TValue, found bool)
}

// RandomAccessWriteableIterator defines a RandomAccessIterator and WritableIterator, which can write from every index in the iterator.
type RandomAccessWriteableIterator[TIndex any, TValue any] interface {
	// *********************    Inherited methods    ********************//
	RandomAccessIterator
	WritableIterator[TValue]
	// ************************    Own methods    ***********************//

	// GetAt sets the value at the given index of the iterator.
	SetAt(i TIndex, value TValue) bool
}

// RandomAccessMappingIterator defines a CollectionIterator, which can be moved to every position in the iterable direction.
// A RandomAccessMappingIterator does not imply bidirectional iteration.
type RandomAccessMappingIterator[TKey any] interface {
	// *********************    Inherited methods    ********************//
	CollectionIterator
	MappingIterator[TKey]
	// ************************    Own methods    ***********************//

	// MoveTo moves the iterator to the given index, if it is reachable.
	MoveToKey(i TKey) bool
}

// RandomAccessReadableMappingIterator defines a RandomAccessMappingIterator and ReadableIterator, which can read from every key in the iterator.
type RandomAccessReadableMappingIterator[TKey any, TValue any] interface {
	// *********************    Inherited methods    ********************//
	RandomAccessMappingIterator[TKey]
	ReadableIterator[TValue]
	// ************************    Own methods    ***********************//

	// GetAt returns the value at the given index of the iterator.
	GetAtKey(i TKey) (value TValue, found bool)
}

// RandomAccessWriteableMappingIterator defines a RandomAccessMappingIterator and WritableIterator, which can write to every key in the iterator.
type RandomAccessWriteableMappingIterator[TKey any, TValue any] interface {
	// *********************    Inherited methods    ********************//
	RandomAccessMappingIterator[TKey]
	WritableIterator[TValue]
	// ************************    Own methods    ***********************//

	// GetAt sets the value at the given index of the iterator.
	SetAtKey(i TKey, value TValue) bool
}
