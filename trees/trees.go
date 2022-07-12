// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package trees provides an abstract Tree interface.
//
// In computer science, a tree is a widely used abstract data type (ADT) or data structure implementing this ADT that simulates a hierarchical tree structure, with a root value and subtrees of children with a parent node, represented as a set of linked nodes.
//
// Reference: https://en.wikipedia.org/wiki/Tree_%28data_structure%29
package trees

import "github.com/JonasMuehlmann/datastructures.go/ds"

// Tree interface that all trees implement.
type Tree[Tkey comparable, TValue any] interface {
	ds.Container
	// IsEmpty() bool
	// Size() int
	// Clear()
	// GetValues() []T
	// ToString() string
}
