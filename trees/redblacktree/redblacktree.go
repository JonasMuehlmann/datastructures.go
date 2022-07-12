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

	"github.com/JonasMuehlmann/datastructures.go/trees"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Tree implementation
var _ trees.Tree[string, any] = (*Tree[string, any])(nil)

type color bool

const (
	black, red color = true, false
)

// Tree holds elements of the red-black tree
type Tree[TKey comparable, TValue any] struct {
	Root       *Node[TKey, TValue]
	size       int
	Comparator utils.Comparator[TKey]
}

// Node is a single element within the tree
type Node[TKey comparable, TValue any] struct {
	Key    TKey
	Value  TValue
	color  color
	Left   *Node[TKey, TValue]
	Right  *Node[TKey, TValue]
	Parent *Node[TKey, TValue]
}

// NewWith instantiates a red-black tree with the custom comparator.
func NewWith[TKey comparable, TValue any](comparator utils.Comparator[TKey]) *Tree[TKey, TValue] {
	return &Tree[TKey, TValue]{Comparator: comparator}
}

// Put inserts node into the tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree[TKey, TValue]) Put(key TKey, value TValue) {
	var insertedNode *Node[TKey, TValue]
	if tree.Root == nil {
		// Assert key is of comparator's type for initial tree
		tree.Comparator(key, key)
		tree.Root = &Node[TKey, TValue]{Key: key, Value: value, color: red}
		insertedNode = tree.Root
	} else {
		node := tree.Root
		loop := true
		for loop {
			compare := tree.Comparator(key, node.Key)
			switch {
			case compare == 0:
				node.Key = key
				node.Value = value
				return
			case compare < 0:
				if node.Left == nil {
					node.Left = &Node[TKey, TValue]{Key: key, Value: value, color: red}
					insertedNode = node.Left
					loop = false
				} else {
					node = node.Left
				}
			case compare > 0:
				if node.Right == nil {
					node.Right = &Node[TKey, TValue]{Key: key, Value: value, color: red}
					insertedNode = node.Right
					loop = false
				} else {
					node = node.Right
				}
			}
		}
		insertedNode.Parent = node
	}
	tree.insertCase1(insertedNode)
	tree.size++
}

// Get searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree[TKey, TValue]) Get(key TKey) (value TValue, found bool) {
	node := tree.lookup(key)
	if node != nil {
		return node.Value, true
	}
	return
}

// GetNode searches the node in the tree by key and returns its node or nil if key is not found in tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree[TKey, TValue]) GetNode(key TKey) *Node[TKey, TValue] {
	return tree.lookup(key)
}

// Remove remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree[TKey, TValue]) Remove(key TKey) {
	var child *Node[TKey, TValue]
	node := tree.lookup(key)
	if node == nil {
		return
	}
	if node.Left != nil && node.Right != nil {
		pred := node.Left.maximumNode()
		node.Key = pred.Key
		node.Value = pred.Value
		node = pred
	}
	if node.Left == nil || node.Right == nil {
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if node.color == black {
			node.color = nodeColor(child)
			tree.deleteCase1(node)
		}
		tree.replaceNode(node, child)
		if node.Parent == nil && child != nil {
			child.color = black
		}
	}
	tree.size--
}

// Empty returns true if tree does not contain any nodes
func (tree *Tree[TKey, TValue]) IsEmpty() bool {
	return tree.size == 0
}

// Size returns number of nodes in the tree.
func (tree *Tree[TKey, TValue]) Size() int {
	return tree.size
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

// GetKeys returns all keys in-order
func (tree *Tree[TKey, TValue]) GetKeys() []TKey {
	keys := make([]TKey, tree.size)
	it := tree.Iterator()
	for i := 0; it.Next(); i++ {
		keys[i] = it.Key[TKey]()
	}
	return keys
}

// Values returns all values in-order based on the key.
func (tree *Tree[TKey, TValue]) GetValues() []TValue {
	values := make([]TValue, tree.size)
	it := tree.Iterator()
	for i := 0; it.Next(); i++ {
		values[i] = it.Value[TValue]()
	}
	return values
}

// Left returns the left-most (min) node or nil if tree is empty.
func (tree *Tree[TKey, TValue]) Left() *Node[TKey, TValue] {
	var parent *Node[TKey, TValue]
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Left
	}
	return parent
}

// Right returns the right-most (max) node or nil if tree is empty.
func (tree *Tree[TKey, TValue]) Right() *Node[TKey, TValue] {
	var parent *Node[TKey, TValue]
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Right
	}
	return parent
}

// Floor Finds floor node of the input key, return the floor node or nil if no floor is found.
// Second return parameter is true if floor was found, otherwise false.
//
// Floor node is defined as the largest node that is smaller than or equal to the given node.
// A floor node may not be found, either because the tree is empty, or because
// all nodes in the tree are larger than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree[TKey, TValue]) Floor(key TKey) (floor *Node[TKey, TValue], found bool) {
	found = false
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node, true
		case compare < 0:
			node = node.Left
		case compare > 0:
			floor, found = node, true
			node = node.Right
		}
	}
	if found {
		return floor, true
	}
	return nil, false
}

// Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling is found.
// Second return parameter is true if ceiling was found, otherwise false.
//
// Ceiling node is defined as the smallest node that is larger than or equal to the given node.
// A ceiling node may not be found, either because the tree is empty, or because
// all nodes in the tree are smaller than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree[TKey, TValue]) Ceiling(key TKey) (ceiling *Node[TKey, TValue], found bool) {
	found = false
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node, true
		case compare < 0:
			ceiling, found = node, true
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	if found {
		return ceiling, true
	}
	return nil, false
}

// Clear removes all nodes from the tree.
func (tree *Tree[TKey, TValue]) Clear() {
	tree.Root = nil
	tree.size = 0
}

// String returns a string representation of container
func (tree *Tree[TKey, TValue]) ToString() string {
	str := "RedBlackTree\n"
	if !tree.IsEmpty() {
		output(tree.Root, "", true, &str)
	}
	return str
}

func (node *Node[TKey, TValue]) String() string {
	return fmt.Sprintf("%v", node.Key)
}

func output[TKey comparable, TValue any](node *Node[TKey, TValue], prefix string, isTail bool, str *string) {
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Right, newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Left, newPrefix, true, str)
	}
}

func (tree *Tree[TKey, TValue]) lookup(key TKey) *Node[TKey, TValue] {
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node
		case compare < 0:
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	return nil
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

func (tree *Tree[TKey, TValue]) rotateLeft(node *Node[TKey, TValue]) {
	right := node.Right
	tree.replaceNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

func (tree *Tree[TKey, TValue]) rotateRight(node *Node[TKey, TValue]) {
	left := node.Left
	tree.replaceNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}

func (tree *Tree[TKey, TValue]) replaceNode(old *Node[TKey, TValue], new *Node[TKey, TValue]) {
	if old.Parent == nil {
		tree.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

func (tree *Tree[TKey, TValue]) insertCase1(node *Node[TKey, TValue]) {
	if node.Parent == nil {
		node.color = black
	} else {
		tree.insertCase2(node)
	}
}

func (tree *Tree[TKey, TValue]) insertCase2(node *Node[TKey, TValue]) {
	if nodeColor(node.Parent) == black {
		return
	}
	tree.insertCase3(node)
}

func (tree *Tree[TKey, TValue]) insertCase3(node *Node[TKey, TValue]) {
	uncle := node.uncle()
	if nodeColor(uncle) == red {
		node.Parent.color = black
		uncle.color = black
		node.grandparent().color = red
		tree.insertCase1(node.grandparent())
	} else {
		tree.insertCase4(node)
	}
}

func (tree *Tree[TKey, TValue]) insertCase4(node *Node[TKey, TValue]) {
	grandparent := node.grandparent()
	if node == node.Parent.Right && node.Parent == grandparent.Left {
		tree.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandparent.Right {
		tree.rotateRight(node.Parent)
		node = node.Right
	}
	tree.insertCase5(node)
}

func (tree *Tree[TKey, TValue]) insertCase5(node *Node[TKey, TValue]) {
	node.Parent.color = black
	grandparent := node.grandparent()
	grandparent.color = red
	if node == node.Parent.Left && node.Parent == grandparent.Left {
		tree.rotateRight(grandparent)
	} else if node == node.Parent.Right && node.Parent == grandparent.Right {
		tree.rotateLeft(grandparent)
	}
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

func (tree *Tree[TKey, TValue]) deleteCase1(node *Node[TKey, TValue]) {
	if node.Parent == nil {
		return
	}
	tree.deleteCase2(node)
}

func (tree *Tree[TKey, TValue]) deleteCase2(node *Node[TKey, TValue]) {
	sibling := node.sibling()
	if nodeColor(sibling) == red {
		node.Parent.color = red
		sibling.color = black
		if node == node.Parent.Left {
			tree.rotateLeft(node.Parent)
		} else {
			tree.rotateRight(node.Parent)
		}
	}
	tree.deleteCase3(node)
}

func (tree *Tree[TKey, TValue]) deleteCase3(node *Node[TKey, TValue]) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == black &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		tree.deleteCase1(node.Parent)
	} else {
		tree.deleteCase4(node)
	}
}

func (tree *Tree[TKey, TValue]) deleteCase4(node *Node[TKey, TValue]) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == red &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		node.Parent.color = black
	} else {
		tree.deleteCase5(node)
	}
}

func (tree *Tree[TKey, TValue]) deleteCase5(node *Node[TKey, TValue]) {
	sibling := node.sibling()
	if node == node.Parent.Left &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == red &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		sibling.Left.color = black
		tree.rotateRight(sibling)
	} else if node == node.Parent.Right &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Right) == red &&
		nodeColor(sibling.Left) == black {
		sibling.color = red
		sibling.Right.color = black
		tree.rotateLeft(sibling)
	}
	tree.deleteCase6(node)
}

func (tree *Tree[TKey, TValue]) deleteCase6(node *Node[TKey, TValue]) {
	sibling := node.sibling()
	sibling.color = nodeColor(node.Parent)
	node.Parent.color = black
	if node == node.Parent.Left && nodeColor(sibling.Right) == red {
		sibling.Right.color = black
		tree.rotateLeft(node.Parent)
	} else if nodeColor(sibling.Left) == red {
		sibling.Left.color = black
		tree.rotateRight(node.Parent)
	}
}

func nodeColor[TKey comparable, TValue any](node *Node[TKey, TValue]) color {
	if node == nil {
		return black
	}
	return node.color
}
