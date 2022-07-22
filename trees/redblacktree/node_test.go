// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package redblacktree implements a red-black tree.
//
// Used by TreeSet and TreeMap.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Red%E2%80%93black_tree

package redblacktree

import (
	"testing"

	"github.com/JonasMuehlmann/datastructures.go/utils"
	"github.com/stretchr/testify/assert"
)

type Pair[TKey any, TValue any] struct {
	Key   TKey
	Value TValue
}

func TestNodeSize(t *testing.T) {
	tests := []struct {
		name  string
		pairs []Pair[int, string]
		size  int
	}{
		{
			name:  "empty list",
			pairs: []Pair[int, string]{},
			size:  0,
		},
		{
			name:  "only left children",
			pairs: []Pair[int, string]{{5, "foo"}, {3, "bar"}, {1, "baz"}},
			size:  3,
		},
		{
			name:  "only right children",
			pairs: []Pair[int, string]{{1, "foo"}, {3, "bar"}, {5, "baz"}},
			size:  3,
		},
		{
			name:  "mixed children",
			pairs: []Pair[int, string]{{1, "foo"}, {5, "bar"}, {2, "baz"}},
			size:  3,
		},
	}

	for _, test := range tests {
		tree := NewWith[int, string](utils.BasicComparator[int])

		for _, pair := range test.pairs {
			tree.Put(pair.Key, pair.Value)
		}

		size := tree.Size()

		assert.Equal(t, test.size, size, test.name)
	}
}
