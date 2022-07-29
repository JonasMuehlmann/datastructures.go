// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package btree implements a B tree.
//
// According to Knuth's definition, a B-tree of order m is a tree which satisfies the following properties:
// - Every node has at most m children.
// - Every non-leaf node (except root) has at least ⌈m/2⌉ children.
// - The root has at least two children if it is not a leaf node.
// - A non-leaf node with k children contains k−1 keys.
// - All leaves appear in the same level
//
// Structure is not thread safe.
//
// References: https://en.wikipedia.org/wiki/B-tree
package btree

// Node is a single element within the tree
type Node[TKey comparable, TValue any] struct {
	Parent   *Node[TKey, TValue]
	Entries  []*Entry[TKey, TValue] // Contained keys in node
	Children []*Node[TKey, TValue]  // Children nodes
}

// Size returns the number of elements stored in the subtree.
// Computed dynamically on each call, i.e. the subtree is traversed to count the number of the nodes.
func (node *Node[TKey, TValue]) Size() int {
	if node == nil {
		return 0
	}
	size := 1
	for _, child := range node.Children {
		size += child.Size()
	}
	return size
}

func (node *Node[TKey, TValue]) height() int {
	height := 0
	for ; node != nil; node = node.Children[0] {
		height++
		if len(node.Children) == 0 {
			break
		}
	}
	return height
}
