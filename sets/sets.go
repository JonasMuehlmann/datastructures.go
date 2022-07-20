// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sets provides an abstract Set interface.
//
// In computer science, a set is an abstract data type that can store certain values and no repeated values. It is a computer implementation of the mathematical concept of a finite set. Unlike most other collection types, rather than retrieving a specific element from a set, one typically tests a value for membership in a set.
//
// Reference: https://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package sets

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Set interface that all sets implement.
type Set[T any] interface {
	Add(elements ...T)
	Remove(comparator utils.Comparator[T], elements ...T)
	Contains(elements ...T) bool
	MakeIntersectionWith(other Set[T]) Set[T]
	MakeUnionWith(other Set[T]) Set[T]
	MakeDifferenceWith(other Set[T]) Set[T]

	ds.Container[T]
	// IsEmpty() bool
	// Size() int
	// Clear()
	// GetValues() []T
	// ToString() string
}
