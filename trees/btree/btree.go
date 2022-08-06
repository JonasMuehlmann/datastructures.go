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

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/trees"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Tree implementation
var _ trees.Tree[string, any] = (*Tree[string, any])(nil)

// Tree holds elements of the B-tree
type Tree[TKey comparable, TValue any] struct {
	Root       *Node[TKey, TValue]    // Root node
	Comparator utils.Comparator[TKey] // Key comparator
	size       int                    // Total number of keys in the tree
	m          int                    // order (maximum number of children)
}

// NewWith instantiates a B-tree with the order (maximum number of children) and a custom key comparator.
func New[TKey comparable, TValue any](order int, comparator utils.Comparator[TKey]) *Tree[TKey, TValue] {
	if order < 3 {
		panic("Invalid order, should be at least 3")
	}
	return &Tree[TKey, TValue]{m: order, Comparator: comparator}
}

// NewFromMap instantiates a new tree containing the provided map.
func NewFromMap[TKey comparable, TValue any](order int, comparator utils.Comparator[TKey], map_ map[TKey]TValue) *Tree[TKey, TValue] {
	tree := New[TKey, TValue](order, comparator)

	for k, v := range map_ {
		tree.Put(k, v)
	}

	return tree
}

// NewFromIterator instantiates a new tree containing the elements provided by the passed iterator.
func NewFromIterator[TKey comparable, TValue any](order int, comparator utils.Comparator[TKey], begin ds.ReadCompForIndexIterator[TKey, TValue]) *Tree[TKey, TValue] {
	tree := New[TKey, TValue](order, comparator)

	for begin.Next() {
		newKey, _ := begin.Index()
		newValue, _ := begin.Get()

		tree.Put(newKey, newValue)
	}

	return tree
}

// NewFromIterators instantiates a new tree containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[TKey comparable, TValue any](order int, comparator utils.Comparator[TKey], begin ds.ReadCompForIndexIterator[TKey, TValue], end ds.CompIndexIterator[TKey]) *Tree[TKey, TValue] {
	tree := New[TKey, TValue](order, comparator)

	for !begin.IsEqual(end) && begin.Next() {
		newKey, _ := begin.Index()
		newValue, _ := begin.Get()

		tree.Put(newKey, newValue)
	}

	return tree
}

// Put inserts key-value pair node into the tree.
// If key already exists, then its value is updated with the new value.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree[TKey, TValue]) Put(key TKey, value TValue) {
	entry := &Entry[TKey, TValue]{Key: key, Value: value}

	if tree.Root == nil {
		tree.Root = &Node[TKey, TValue]{Entries: []*Entry[TKey, TValue]{entry}, Children: []*Node[TKey, TValue]{}}
		tree.size++
		return
	}

	if tree.insert(tree.Root, entry) {
		tree.size++
	}
}

// Get searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree[TKey, TValue]) Get(key TKey) (value TValue, found bool) {
	node, index, found := tree.searchRecursively(tree.Root, key)
	if found {
		return node.Entries[index].Value, true
	}
	return
}

// GetNode searches the node in the tree by key and returns its node or nil if key is not found in tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree[TKey, TValue]) GetNode(key TKey) *Node[TKey, TValue] {
	node, _, _ := tree.searchRecursively(tree.Root, key)
	return node
}

// Remove remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree[TKey, TValue]) Remove(key TKey) {
	node, index, found := tree.searchRecursively(tree.Root, key)
	if found {
		tree.delete(node, index)
		tree.size--
	}
}

// Empty returns true if tree does not contain any nodes
func (tree *Tree[TKey, TValue]) IsEmpty() bool {
	return tree.size == 0
}

// Size returns number of nodes in the tree.
func (tree *Tree[TKey, TValue]) Size() int {
	return tree.size
}

// GetKeys returns all keys in-order
func (tree *Tree[TKey, TValue]) GetKeys() []TKey {
	keys := make([]TKey, 0, tree.size)

	it := tree.OrderedBegin()

	for it.Next() {
		newIndex, _ := it.Index()
		keys = append(keys, newIndex)
	}

	return keys
}

// Values returns all values in-order based on the key.
func (tree *Tree[TKey, TValue]) GetValues() []TValue {
	values := make([]TValue, 0, tree.size)

	it := tree.OrderedBegin()

	for it.Next() {
		newValue, _ := it.Get()
		values = append(values, newValue)
	}

	return values
}

// Clear removes all nodes from the tree.
func (tree *Tree[TKey, TValue]) Clear() {
	tree.Root = nil
	tree.size = 0
}

// Height returns the height of the tree.
func (tree *Tree[TKey, TValue]) Height() int {
	return tree.Root.height()
}

// Left returns the left-most (min) node or nil if tree is empty.
func (tree *Tree[TKey, TValue]) Left() *Node[TKey, TValue] {
	return tree.left(tree.Root)
}

// LeftKey returns the left-most (min) key or nil if tree is empty.
func (tree *Tree[TKey, TValue]) LeftKey() (key TKey, found bool) {
	if left := tree.Left(); left != nil {
		return left.Entries[0].Key, true
	}
	return
}

// LeftValue returns the left-most value or nil if tree is empty.
func (tree *Tree[TKey, TValue]) LeftValue() (value TValue, found bool) {
	if left := tree.Left(); left != nil {
		return left.Entries[0].Value, true
	}
	return
}

// Right returns the right-most (max) node or nil if tree is empty.
func (tree *Tree[TKey, TValue]) Right() *Node[TKey, TValue] {
	return tree.right(tree.Root)
}

// RightKey returns the right-most (max) key or nil if tree is empty.
func (tree *Tree[TKey, TValue]) RightKey() (key TKey, found bool) {
	if right := tree.Right(); right != nil {
		return right.Entries[len(right.Entries)-1].Key, true
	}
	return
}

// RightValue returns the right-most value or nil if tree is empty.
func (tree *Tree[TKey, TValue]) RightValue() (value TValue, found bool) {
	if right := tree.Right(); right != nil {
		return right.Entries[len(right.Entries)-1].Value, true
	}
	return
}

// String returns a string representation of container (for debugging purposes)
func (tree *Tree[TKey, TValue]) ToString() string {
	var buffer bytes.Buffer
	buffer.WriteString("BTree\n")
	if !tree.IsEmpty() {
		tree.output(&buffer, tree.Root, 0, true)
	}
	return buffer.String()
}

func (tree *Tree[TKey, TValue]) output(buffer *bytes.Buffer, node *Node[TKey, TValue], level int, isTail bool) {
	for e := 0; e < len(node.Entries)+1; e++ {
		if e < len(node.Children) {
			tree.output(buffer, node.Children[e], level+1, true)
		}
		if e < len(node.Entries) {
			buffer.WriteString(strings.Repeat("    ", level))
			buffer.WriteString(fmt.Sprintf("%v", node.Entries[e].Key) + "\n")
		}
	}
}

func (tree *Tree[TKey, TValue]) isLeaf(node *Node[TKey, TValue]) bool {
	return len(node.Children) == 0
}

func (tree *Tree[TKey, TValue]) isFull(node *Node[TKey, TValue]) bool {
	return len(node.Entries) == tree.maxEntries()
}

func (tree *Tree[TKey, TValue]) shouldSplit(node *Node[TKey, TValue]) bool {
	return len(node.Entries) > tree.maxEntries()
}

func (tree *Tree[TKey, TValue]) maxChildren() int {
	return tree.m
}

func (tree *Tree[TKey, TValue]) minChildren() int {
	return (tree.m + 1) / 2 // ceil(m/2)
}

func (tree *Tree[TKey, TValue]) maxEntries() int {
	return tree.maxChildren() - 1
}

func (tree *Tree[TKey, TValue]) minEntries() int {
	return tree.minChildren() - 1
}

func (tree *Tree[TKey, TValue]) middle() int {
	return (tree.m - 1) / 2 // "-1" to favor right nodes to have more keys when splitting
}

// search searches only within the single node among its entries
func (tree *Tree[TKey, TValue]) search(node *Node[TKey, TValue], key TKey) (index int, found bool) {
	low, high := 0, len(node.Entries)-1
	var mid int
	for low <= high {
		mid = (high + low) / 2
		compare := tree.Comparator(key, node.Entries[mid].Key)
		switch {
		case compare > 0:
			low = mid + 1
		case compare < 0:
			high = mid - 1
		case compare == 0:
			return mid, true
		}
	}
	return low, false
}

// searchRecursively searches recursively down the tree starting at the startNode
func (tree *Tree[TKey, TValue]) searchRecursively(startNode *Node[TKey, TValue], key TKey) (node *Node[TKey, TValue], index int, found bool) {
	if tree.IsEmpty() {
		return nil, -1, false
	}
	node = startNode
	for {
		index, found = tree.search(node, key)
		if found {
			return node, index, true
		}
		if tree.isLeaf(node) {
			return nil, -1, false
		}
		node = node.Children[index]
	}
}

func (tree *Tree[TKey, TValue]) insert(node *Node[TKey, TValue], entry *Entry[TKey, TValue]) (inserted bool) {
	if tree.isLeaf(node) {
		return tree.insertIntoLeaf(node, entry)
	}
	return tree.insertIntoInternal(node, entry)
}

func (tree *Tree[TKey, TValue]) insertIntoLeaf(node *Node[TKey, TValue], entry *Entry[TKey, TValue]) (inserted bool) {
	insertPosition, found := tree.search(node, entry.Key)
	if found {
		node.Entries[insertPosition] = entry
		return false
	}
	// Insert entry's key in the middle of the node
	node.Entries = append(node.Entries, nil)
	copy(node.Entries[insertPosition+1:], node.Entries[insertPosition:])
	node.Entries[insertPosition] = entry
	tree.split(node)
	return true
}

func (tree *Tree[TKey, TValue]) insertIntoInternal(node *Node[TKey, TValue], entry *Entry[TKey, TValue]) (inserted bool) {
	insertPosition, found := tree.search(node, entry.Key)
	if found {
		node.Entries[insertPosition] = entry
		return false
	}
	return tree.insert(node.Children[insertPosition], entry)
}

func (tree *Tree[TKey, TValue]) split(node *Node[TKey, TValue]) {
	if !tree.shouldSplit(node) {
		return
	}

	if node == tree.Root {
		tree.splitRoot()
		return
	}

	tree.splitNonRoot(node)
}

func (tree *Tree[TKey, TValue]) splitNonRoot(node *Node[TKey, TValue]) {
	middle := tree.middle()
	parent := node.Parent

	left := &Node[TKey, TValue]{Entries: append([]*Entry[TKey, TValue](nil), node.Entries[:middle]...), Parent: parent}
	right := &Node[TKey, TValue]{Entries: append([]*Entry[TKey, TValue](nil), node.Entries[middle+1:]...), Parent: parent}

	// Move children from the node to be split into left and right nodes
	if !tree.isLeaf(node) {
		left.Children = append([]*Node[TKey, TValue](nil), node.Children[:middle+1]...)
		right.Children = append([]*Node[TKey, TValue](nil), node.Children[middle+1:]...)
		setParent(left.Children, left)
		setParent(right.Children, right)
	}

	insertPosition, _ := tree.search(parent, node.Entries[middle].Key)

	// Insert middle key into parent
	parent.Entries = append(parent.Entries, nil)
	copy(parent.Entries[insertPosition+1:], parent.Entries[insertPosition:])
	parent.Entries[insertPosition] = node.Entries[middle]

	// Set child left of inserted key in parent to the created left node
	parent.Children[insertPosition] = left

	// Set child right of inserted key in parent to the created right node
	parent.Children = append(parent.Children, nil)
	copy(parent.Children[insertPosition+2:], parent.Children[insertPosition+1:])
	parent.Children[insertPosition+1] = right

	tree.split(parent)
}

func (tree *Tree[TKey, TValue]) splitRoot() {
	middle := tree.middle()

	left := &Node[TKey, TValue]{Entries: append([]*Entry[TKey, TValue](nil), tree.Root.Entries[:middle]...)}
	right := &Node[TKey, TValue]{Entries: append([]*Entry[TKey, TValue](nil), tree.Root.Entries[middle+1:]...)}

	// Move children from the node to be split into left and right nodes
	if !tree.isLeaf(tree.Root) {
		left.Children = append([]*Node[TKey, TValue](nil), tree.Root.Children[:middle+1]...)
		right.Children = append([]*Node[TKey, TValue](nil), tree.Root.Children[middle+1:]...)
		setParent(left.Children, left)
		setParent(right.Children, right)
	}

	// Root is a node with one entry and two children (left and right)
	newRoot := &Node[TKey, TValue]{
		Entries:  []*Entry[TKey, TValue]{tree.Root.Entries[middle]},
		Children: []*Node[TKey, TValue]{left, right},
	}

	left.Parent = newRoot
	right.Parent = newRoot
	tree.Root = newRoot
}

func setParent[TKey comparable, TValue any](nodes []*Node[TKey, TValue], parent *Node[TKey, TValue]) {
	for _, node := range nodes {
		node.Parent = parent
	}
}

func (tree *Tree[TKey, TValue]) left(node *Node[TKey, TValue]) *Node[TKey, TValue] {
	if tree.IsEmpty() {
		return nil
	}
	current := node
	for {
		if tree.isLeaf(current) {
			return current
		}
		current = current.Children[0]
	}
}

func (tree *Tree[TKey, TValue]) right(node *Node[TKey, TValue]) *Node[TKey, TValue] {
	if tree.IsEmpty() {
		return nil
	}
	current := node
	for {
		if tree.isLeaf(current) {
			return current
		}
		current = current.Children[len(current.Children)-1]
	}
}

// leftSibling returns the node's left sibling and child index (in parent) if it exists, otherwise (nil,-1)
// key is any of keys in node (could even be deleted).
func (tree *Tree[TKey, TValue]) leftSibling(node *Node[TKey, TValue], key TKey) (*Node[TKey, TValue], int) {
	if node.Parent != nil {
		index, _ := tree.search(node.Parent, key)
		index--
		if index >= 0 && index < len(node.Parent.Children) {
			return node.Parent.Children[index], index
		}
	}
	return nil, -1
}

// rightSibling returns the node's right sibling and child index (in parent) if it exists, otherwise (nil,-1)
// key is any of keys in node (could even be deleted).
func (tree *Tree[TKey, TValue]) rightSibling(node *Node[TKey, TValue], key TKey) (*Node[TKey, TValue], int) {
	if node.Parent != nil {
		index, _ := tree.search(node.Parent, key)
		index++
		if index < len(node.Parent.Children) {
			return node.Parent.Children[index], index
		}
	}
	return nil, -1
}

// delete deletes an entry in node at entries' index
// ref.: https://en.wikipedia.org/wiki/B-tree#Deletion
func (tree *Tree[TKey, TValue]) delete(node *Node[TKey, TValue], index int) {
	// deleting from a leaf node
	if tree.isLeaf(node) {
		deletedKey := node.Entries[index].Key
		tree.deleteEntry(node, index)
		tree.rebalance(node, deletedKey)
		if len(tree.Root.Entries) == 0 {
			tree.Root = nil
		}
		return
	}

	// deleting from an internal node
	leftLargestNode := tree.right(node.Children[index]) // largest node in the left sub-tree (assumed to exist)
	leftLargestEntryIndex := len(leftLargestNode.Entries) - 1
	node.Entries[index] = leftLargestNode.Entries[leftLargestEntryIndex]
	deletedKey := leftLargestNode.Entries[leftLargestEntryIndex].Key
	tree.deleteEntry(leftLargestNode, leftLargestEntryIndex)
	tree.rebalance(leftLargestNode, deletedKey)
}

// rebalance rebalances the tree after deletion if necessary and returns true, otherwise false.
// Note that we first delete the entry and then call rebalance, thus the passed deleted key as reference.
func (tree *Tree[TKey, TValue]) rebalance(node *Node[TKey, TValue], deletedKey TKey) {
	// check if rebalancing is needed
	if node == nil || len(node.Entries) >= tree.minEntries() {
		return
	}

	// try to borrow from left sibling
	leftSibling, leftSiblingIndex := tree.leftSibling(node, deletedKey)
	if leftSibling != nil && len(leftSibling.Entries) > tree.minEntries() {
		// rotate right
		node.Entries = append([]*Entry[TKey, TValue]{node.Parent.Entries[leftSiblingIndex]}, node.Entries...) // prepend parent's separator entry to node's entries
		node.Parent.Entries[leftSiblingIndex] = leftSibling.Entries[len(leftSibling.Entries)-1]
		tree.deleteEntry(leftSibling, len(leftSibling.Entries)-1)
		if !tree.isLeaf(leftSibling) {
			leftSiblingRightMostChild := leftSibling.Children[len(leftSibling.Children)-1]
			leftSiblingRightMostChild.Parent = node
			node.Children = append([]*Node[TKey, TValue]{leftSiblingRightMostChild}, node.Children...)
			tree.deleteChild(leftSibling, len(leftSibling.Children)-1)
		}
		return
	}

	// try to borrow from right sibling
	rightSibling, rightSiblingIndex := tree.rightSibling(node, deletedKey)
	if rightSibling != nil && len(rightSibling.Entries) > tree.minEntries() {
		// rotate left
		node.Entries = append(node.Entries, node.Parent.Entries[rightSiblingIndex-1]) // append parent's separator entry to node's entries
		node.Parent.Entries[rightSiblingIndex-1] = rightSibling.Entries[0]
		tree.deleteEntry(rightSibling, 0)
		if !tree.isLeaf(rightSibling) {
			rightSiblingLeftMostChild := rightSibling.Children[0]
			rightSiblingLeftMostChild.Parent = node
			node.Children = append(node.Children, rightSiblingLeftMostChild)
			tree.deleteChild(rightSibling, 0)
		}
		return
	}

	// merge with siblings
	if rightSibling != nil {
		// merge with right sibling
		node.Entries = append(node.Entries, node.Parent.Entries[rightSiblingIndex-1])
		node.Entries = append(node.Entries, rightSibling.Entries...)
		deletedKey = node.Parent.Entries[rightSiblingIndex-1].Key
		tree.deleteEntry(node.Parent, rightSiblingIndex-1)
		tree.appendChildren(node.Parent.Children[rightSiblingIndex], node)
		tree.deleteChild(node.Parent, rightSiblingIndex)
	} else if leftSibling != nil {
		// merge with left sibling
		entries := append([]*Entry[TKey, TValue](nil), leftSibling.Entries...)
		entries = append(entries, node.Parent.Entries[leftSiblingIndex])
		node.Entries = append(entries, node.Entries...)
		deletedKey = node.Parent.Entries[leftSiblingIndex].Key
		tree.deleteEntry(node.Parent, leftSiblingIndex)
		tree.prependChildren(node.Parent.Children[leftSiblingIndex], node)
		tree.deleteChild(node.Parent, leftSiblingIndex)
	}

	// make the merged node the root if its parent was the root and the root is empty
	if node.Parent == tree.Root && len(tree.Root.Entries) == 0 {
		tree.Root = node
		node.Parent = nil
		return
	}

	// parent might underflow, so try to rebalance if necessary
	tree.rebalance(node.Parent, deletedKey)
}

func (tree *Tree[TKey, TValue]) prependChildren(fromNode *Node[TKey, TValue], toNode *Node[TKey, TValue]) {
	children := append([]*Node[TKey, TValue](nil), fromNode.Children...)
	toNode.Children = append(children, toNode.Children...)
	setParent(fromNode.Children, toNode)
}

func (tree *Tree[TKey, TValue]) appendChildren(fromNode *Node[TKey, TValue], toNode *Node[TKey, TValue]) {
	toNode.Children = append(toNode.Children, fromNode.Children...)
	setParent(fromNode.Children, toNode)
}

func (tree *Tree[TKey, TValue]) deleteEntry(node *Node[TKey, TValue], index int) {
	copy(node.Entries[index:], node.Entries[index+1:])
	node.Entries[len(node.Entries)-1] = nil
	node.Entries = node.Entries[:len(node.Entries)-1]
}

func (tree *Tree[TKey, TValue]) deleteChild(node *Node[TKey, TValue], index int) {
	if index >= len(node.Children) {
		return
	}
	copy(node.Children[index:], node.Children[index+1:])
	node.Children[len(node.Children)-1] = nil
	node.Children = node.Children[:len(node.Children)-1]
}

// func (tree *Tree[TKey, TValue]) lookup(key TKey) *Node[TKey, TValue] {
// 	node := tree.Root

// 	for node != nil {
// 		compare := tree.Comparator(key, node.Key)

// 		switch {
// 		case compare == 0:
// 			return node
// 		case compare < 0:
// 			node = node.Children[0]
// 		case compare > 0:
// 			node = node.Children[1]
// 		}
// 	}

// 	return nil
// }

// func findLowestCommonAncestor[TKey comparable, TValue any](start, node1, node2 *Node[TKey, TValue]) *Node[TKey, TValue] {
// 	if start == nil {
// 		return nil
// 	}

// 	if start == node1 || start == node2 {
// 		return start
// 	}

// 	leftChild := findLowestCommonAncestor(start.Children[0], node1, node2)
// 	rightChild := findLowestCommonAncestor(start.Children[1], node1, node2)

// 	if leftChild != nil && rightChild != nil {
// 		return start
// 	}

// 	if leftChild == nil {
// 		return rightChild
// 	}

// 	return leftChild
// }

// func getDistanceFromLCA[TKey comparable, TValue any](comparator utils.Comparator[TKey], lca, child *Node[TKey, TValue], distance int, originalToTargetOrdering int, isOriginal bool) int {
// 	if lca == child {
// 		return distance
// 	}

// for child := range lca.Children{
//     if child.Key ==  lca {
//
//     }
//
//
// 		return distance + 1
// 	}

// 	newDistance := distance

// 	if comparator(child.Key, lca.Key) < 0 {
// 		if originalToTargetOrdering < 0 {
// 			newDistance += 1
// 		}

// 		return getDistanceFromLCA(comparator, lca.Children[0], child, newDistance, originalToTargetOrdering, true)
// 	}
// 	if originalToTargetOrdering > 0 {
// 		newDistance += 1
// 	}

// 	return getDistanceFromLCA(comparator, lca.Children[1], child, newDistance, originalToTargetOrdering, true)

// }

// func distanceBetween[TKey comparable, TValue any](comparator utils.Comparator[TKey], root, original *Node[TKey, TValue], indexOriginal int, target *Node[TKey, TValue], indexTargetNode int) int {
// 	distance := 0

// 	lca, index, found := findLowestCommonAncestor(root, original, target)

// 	return getDistanceFromLCA(comparator, lca, original, distance, comparator(original.Key, target.Key), true) + getDistanceFromLCA(comparator, lca, target, distance, comparator(original.Key, target.Key), false)
// }

//******************************************************************//
//                             Iterator                             //
//******************************************************************//

// Begin returns an initialized iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.
func (tree *Tree[TKey, TValue]) OrderedBegin() ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return tree.NewOrderedIterator(-1)
}

// End returns an initialized iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.
func (tree *Tree[TKey, TValue]) OrderedEnd() ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return tree.NewOrderedIterator(tree.Size())
}

// First returns an initialized iterator, which points to it's first element.
func (tree *Tree[TKey, TValue]) OrderedFirst() ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return tree.NewOrderedIterator(0)
}

// Last returns an initialized iterator, which points to it's last element.
func (tree *Tree[TKey, TValue]) OrderedLast() ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return tree.NewOrderedIterator(tree.Size() - 1)
}
