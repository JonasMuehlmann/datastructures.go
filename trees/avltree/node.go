// Copyright (c) 2017, Benjamin Scher Purcell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package avltree implements an AVL balanced binary tree.
//
// Structure is not thread safe.
//
// References: https://en.wikipedia.org/wiki/AVL_tree
package avltree

import "fmt"

// Node is a single element within the tree.
type Node[TKey comparable, TValue any] struct {
	Key      TKey
	Value    TValue
	Parent   *Node[TKey, TValue]    // Parent node
	Children [2]*Node[TKey, TValue] // Children nodes
	b        int8
}

// Size returns the number of elements stored in the subtree.
// Computed dynamically on each call, i.e. the subtree is traversed to count the number of the nodes.
func (n *Node[TKey, TValue]) Size() int {
	if n == nil {
		return 0
	}

	size := 1

	if n.Children[0] != nil {
		size += n.Children[0].Size()
	}

	if n.Children[1] != nil {
		size += n.Children[1].Size()
	}

	return size
}

func (n *Node[TKey, TValue]) String() string {
	return fmt.Sprintf("%v", n.Key)
}

// Prev returns the previous element in an inorder
// walk of the AVL tree.
func (n *Node[TKey, TValue]) Prev() *Node[TKey, TValue] {
	return n.walk1(0)
}

// Next returns the next element in an inorder
// walk of the AVL tree.
func (n *Node[TKey, TValue]) Next() *Node[TKey, TValue] {
	return n.walk1(1)
}

func (n *Node[TKey, TValue]) walk1(a int) *Node[TKey, TValue] {
	if n == nil {
		return nil
	}

	if n.Children[a] != nil {
		n = n.Children[a]
		for n.Children[a^1] != nil {
			n = n.Children[a^1]
		}

		return n
	}

	p := n.Parent
	for p != nil && p.Children[a] == n {
		n = p
		p = p.Parent
	}

	return p
}
