// Copyright (c) 2022, Jonas Muehlmann. All rights reserved.
// Copyright (c) 2015, Emir Pasic. All rights reserved.
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
	"fmt"
)

// Node is a single element within the tree
type Node[TKey any, TValue any] struct {
	Key    TKey
	Value  TValue
	color  color
	Left   *Node[TKey, TValue]
	Right  *Node[TKey, TValue]
	Parent *Node[TKey, TValue]
}

// Size returns the number of elements stored in the subtree.
// Computed dynamically on each call, i.e. the subtree is traversed to count the number of the nodes.
func (node *Node[TKey, TValue]) Size() int {
	if node == nil {
		return 0
	}
	size := 1
	if node.Left != nil {
		size += node.Left.Size()
	}
	if node.Right != nil {
		size += node.Right.Size()
	}
	return size
}

func (node *Node[TKey, TValue]) String() string {
	return fmt.Sprintf("%v", node.Key)
}

func (node *Node[TKey, TValue]) grandparent() *Node[TKey, TValue] {
	if node != nil && node.Parent != nil {
		return node.Parent.Parent
	}
	return nil
}

func (node *Node[TKey, TValue]) uncle() *Node[TKey, TValue] {
	if node == nil || node.Parent == nil || node.Parent.Parent == nil {
		return nil
	}
	return node.Parent.sibling()
}

func (node *Node[TKey, TValue]) sibling() *Node[TKey, TValue] {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	}
	return node.Parent.Left
}

func (node *Node[TKey, TValue]) maximumNode() *Node[TKey, TValue] {
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}
