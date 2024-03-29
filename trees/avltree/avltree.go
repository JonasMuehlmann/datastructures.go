// Copyright (c) 2017, Benjamin Scher Purcell. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package avltree implements an AVL balanced binary tree.
//
// Structure is not thread safe.
//
// References: https://en.wikipedia.org/wiki/AVL_tree
package avltree

import (
	"github.com/JonasMuehlmann/datastructures.go/ds"
	"github.com/JonasMuehlmann/datastructures.go/trees"
	"github.com/JonasMuehlmann/datastructures.go/utils"
)

// Assert Tree implementation
var _ trees.Tree[string, any] = new(Tree[string, any])

// Tree holds elements of the AVL tree.
type Tree[TKey comparable, TValue any] struct {
	Root       *Node[TKey, TValue]    // Root node
	Comparator utils.Comparator[TKey] // Key comparator
	size       int                    // Total number of keys in the tree
}

// NewWith instantiates an AVL tree with the custom comparator.
func New[TKey comparable, TValue any](comparator utils.Comparator[TKey]) *Tree[TKey, TValue] {
	return &Tree[TKey, TValue]{Comparator: comparator}
}

// NewFromMap instantiates a new tree containing the provided map.
func NewFromMap[TKey comparable, TValue any](comparator utils.Comparator[TKey], map_ map[TKey]TValue) *Tree[TKey, TValue] {
	tree := New[TKey, TValue](comparator)

	for k, v := range map_ {
		tree.Put(k, v)
	}

	return tree
}

// NewFromIterator instantiates a new tree containing the elements provided by the passed iterator.
func NewFromIterator[TKey comparable, TValue any](comparator utils.Comparator[TKey], begin ds.ReadForIndexIterator[TKey, TValue]) *Tree[TKey, TValue] {
	tree := New[TKey, TValue](comparator)

	for begin.Next() {
		newKey, _ := begin.GetKey()
		newValue, _ := begin.Get()

		tree.Put(newKey, newValue)
	}

	return tree
}

// NewFromIterators instantiates a new tree containing the elements provided by first, until it is equal to end.
// end is a sentinel and not included.
func NewFromIterators[TKey comparable, TValue any](comparator utils.Comparator[TKey], begin ds.ReadCompForIndexIterator[TKey, TValue], end ds.CompIndexIterator[TKey]) *Tree[TKey, TValue] {
	tree := New[TKey, TValue](comparator)

	for !begin.IsEqual(end) && begin.Next() {
		newKey, _ := begin.GetKey()
		newValue, _ := begin.Get()

		tree.Put(newKey, newValue)
	}

	return tree
}

// Put inserts node into the tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[TKey, TValue]) Put(key TKey, value TValue) {
	t.put(key, value, nil, &t.Root)
}

// Get searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[TKey, TValue]) Get(key TKey) (value TValue, found bool) {
	n := t.GetNode(key)
	if n != nil {
		return n.Value, true
	}
	return
}

// GetNode searches the node in the tree by key and returns its node or nil if key is not found in tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[TKey, TValue]) GetNode(key TKey) *Node[TKey, TValue] {
	n := t.Root
	for n != nil {
		cmp := t.Comparator(key, n.Key)
		switch {
		case cmp == 0:
			return n
		case cmp < 0:
			n = n.Children[0]
		case cmp > 0:
			n = n.Children[1]
		}
	}
	return n
}

// Remove remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[TKey, TValue]) Remove(key TKey) {
	t.remove(key, &t.Root)
}

// Empty returns true if tree does not contain any nodes.
func (t *Tree[TKey, TValue]) IsEmpty() bool {
	return t.size == 0
}

// Size returns the number of elements stored in the tree.
func (t *Tree[TKey, TValue]) Size() int {
	return t.size
}

// GetKeys returns all keys in-order
func (tree *Tree[TKey, TValue]) GetKeys() []TKey {
	keys := make([]TKey, 0, tree.size)

	it := tree.OrderedBegin()

	for it.Next() {
		newIndex, _ := it.GetKey()
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

// Left returns the minimum element of the AVL tree
// or nil if the tree is empty.
func (t *Tree[TKey, TValue]) Left() *Node[TKey, TValue] {
	return t.bottom(0)
}

// Right returns the maximum element of the AVL tree
// or nil if the tree is empty.
func (t *Tree[TKey, TValue]) Right() *Node[TKey, TValue] {
	return t.bottom(1)
}

// Floor Finds floor node of the input key, return the floor node or nil if no ceiling is found.
// Second return parameter is true if floor was found, otherwise false.
//
// Floor node is defined as the largest node that is smaller than or equal to the given node.
// A floor node may not be found, either because the tree is empty, or because
// all nodes in the tree is larger than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[TKey, TValue]) Floor(key TKey) (floor *Node[TKey, TValue], found bool) {
	found = false

	n := t.Root
	for n != nil {
		c := t.Comparator(key, n.Key)

		switch {
		case c == 0:
			return n, true
		case c < 0:
			n = n.Children[0]
		case c > 0:
			floor, found = n, true
			n = n.Children[1]
		}
	}

	if found {
		return
	}

	return nil, false
}

// Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling is found.
// Second return parameter is true if ceiling was found, otherwise false.
//
// Ceiling node is defined as the smallest node that is larger than or equal to the given node.
// A ceiling node may not be found, either because the tree is empty, or because
// all nodes in the tree is smaller than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (t *Tree[TKey, TValue]) Ceiling(key TKey) (floor *Node[TKey, TValue], found bool) {
	found = false

	n := t.Root
	for n != nil {
		c := t.Comparator(key, n.Key)

		switch {
		case c == 0:
			return n, true
		case c < 0:
			floor, found = n, true
			n = n.Children[0]
		case c > 0:
			n = n.Children[1]
		}
	}

	if found {
		return
	}

	return nil, false
}

// Clear removes all nodes from the tree.
func (t *Tree[TKey, TValue]) Clear() {
	t.Root = nil
	t.size = 0
}

// String returns a string representation of container
func (t *Tree[TKey, TValue]) ToString() string {
	str := "AVLTree\n"

	if !t.IsEmpty() {
		output(t.Root, "", true, &str)
	}

	return str
}

func (t *Tree[TKey, TValue]) put(key TKey, value TValue, p *Node[TKey, TValue], qp **Node[TKey, TValue]) bool {
	q := *qp
	if q == nil {
		t.size++
		*qp = &Node[TKey, TValue]{Key: key, Value: value, Parent: p}

		return true
	}

	c := t.Comparator(key, q.Key)
	if c == 0 {
		q.Key = key
		q.Value = value

		return false
	}

	if c < 0 {
		c = -1
	} else {
		c = 1
	}

	a := (c + 1) / 2
	var fix bool

	fix = t.put(key, value, q, &q.Children[a])
	if fix {
		return putFix(int8(c), qp)
	}

	return false
}

func (t *Tree[TKey, TValue]) remove(key TKey, qp **Node[TKey, TValue]) bool {
	q := *qp
	if q == nil {
		return false
	}

	c := t.Comparator(key, q.Key)
	if c == 0 {
		t.size--
		if q.Children[1] == nil {
			if q.Children[0] != nil {
				q.Children[0].Parent = q.Parent
			}

			*qp = q.Children[0]

			return true
		}

		fix := removeMin(&q.Children[1], &q.Key, &q.Value)
		if fix {
			return removeFix(-1, qp)
		}

		return false
	}

	if c < 0 {
		c = -1
	} else {
		c = 1
	}
	a := (c + 1) / 2

	fix := t.remove(key, &q.Children[a])
	if fix {
		return removeFix(int8(-c), qp)
	}
	return false
}

func removeMin[TKey comparable, TValue any](qp **Node[TKey, TValue], minKey *TKey, minVal *TValue) bool {
	q := *qp
	if q.Children[0] == nil {
		*minKey = q.Key
		*minVal = q.Value

		if q.Children[1] != nil {
			q.Children[1].Parent = q.Parent
		}

		*qp = q.Children[1]

		return true
	}

	fix := removeMin(&q.Children[0], minKey, minVal)

	if fix {
		return removeFix(1, qp)
	}

	return false
}

func putFix[TKey comparable, TValue any](c int8, t **Node[TKey, TValue]) bool {
	s := *t
	if s.b == 0 {
		s.b = c

		return true
	}

	if s.b == -c {
		s.b = 0

		return false
	}

	if s.Children[(c+1)/2].b == c {
		s = singlerot(c, s)
	} else {
		s = doublerot(c, s)
	}

	*t = s

	return false
}

func removeFix[TKey comparable, TValue any](c int8, t **Node[TKey, TValue]) bool {
	s := *t
	if s.b == 0 {
		s.b = c

		return false
	}

	if s.b == -c {
		s.b = 0

		return true
	}

	a := (c + 1) / 2
	if s.Children[a].b == 0 {
		s = rotate(c, s)
		s.b = -c
		*t = s

		return false
	}

	if s.Children[a].b == c {
		s = singlerot(c, s)
	} else {
		s = doublerot(c, s)
	}

	*t = s

	return true
}

func singlerot[TKey comparable, TValue any](c int8, s *Node[TKey, TValue]) *Node[TKey, TValue] {
	s.b = 0
	s = rotate(c, s)
	s.b = 0

	return s
}

func doublerot[TKey comparable, TValue any](c int8, s *Node[TKey, TValue]) *Node[TKey, TValue] {
	a := (c + 1) / 2
	r := s.Children[a]
	s.Children[a] = rotate(-c, s.Children[a])
	p := rotate(c, s)

	switch {
	default:
		s.b = 0
		r.b = 0
	case p.b == c:
		s.b = -c
		r.b = 0
	case p.b == -c:
		s.b = 0
		r.b = c
	}

	p.b = 0

	return p
}

func rotate[TKey comparable, TValue any](c int8, s *Node[TKey, TValue]) *Node[TKey, TValue] {
	a := (c + 1) / 2
	r := s.Children[a]

	s.Children[a] = r.Children[a^1]
	if s.Children[a] != nil {
		s.Children[a].Parent = s
	}

	r.Children[a^1] = s
	r.Parent = s.Parent
	s.Parent = r

	return r
}

func (t *Tree[TKey, TValue]) bottom(d int) *Node[TKey, TValue] {
	n := t.Root
	if n == nil {
		return nil
	}

	for c := n.Children[d]; c != nil; c = n.Children[d] {
		n = c
	}

	return n
}

func output[TKey comparable, TValue any](node *Node[TKey, TValue], prefix string, isTail bool, str *string) {
	if node.Children[1] != nil {
		newPrefix := prefix

		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}

		output(node.Children[1], newPrefix, false, str)
	}

	*str += prefix

	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}

	*str += node.String() + "\n"

	if node.Children[0] != nil {
		newPrefix := prefix

		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}

		output(node.Children[0], newPrefix, true, str)
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
			node = node.Children[0]
		case compare > 0:
			node = node.Children[1]
		}
	}

	return nil
}

func findLowestCommonAncestor[TKey comparable, TValue any](start, node1, node2 *Node[TKey, TValue]) *Node[TKey, TValue] {
	if start == nil {
		return nil
	}

	if start == node1 || start == node2 {
		return start
	}

	leftChild := findLowestCommonAncestor(start.Children[0], node1, node2)
	rightChild := findLowestCommonAncestor(start.Children[1], node1, node2)

	if leftChild != nil && rightChild != nil {
		return start
	}

	if leftChild == nil {
		return rightChild
	}

	return leftChild
}

func getDistanceFromLCA[TKey comparable, TValue any](comparator utils.Comparator[TKey], lca, child *Node[TKey, TValue], distance int, originalToTargetOrdering int, isOriginal bool) int {
	if lca == child {
		return distance
	}

	if child == lca.Children[0] || child == lca.Children[1] {
		return distance + 1
	}

	newDistance := distance

	if comparator(child.Key, lca.Key) < 0 {
		if originalToTargetOrdering < 0 {
			newDistance += 1
		}

		return getDistanceFromLCA(comparator, lca.Children[0], child, newDistance, originalToTargetOrdering, true)
	}
	if originalToTargetOrdering > 0 {
		newDistance += 1
	}

	return getDistanceFromLCA(comparator, lca.Children[1], child, newDistance, originalToTargetOrdering, true)

}

func distanceBetween[TKey comparable, TValue any](comparator utils.Comparator[TKey], root, original, target *Node[TKey, TValue]) int {
	distance := 0

	lca := findLowestCommonAncestor(root, original, target)

	return getDistanceFromLCA(comparator, lca, original, distance, comparator(original.Key, target.Key), true) + getDistanceFromLCA(comparator, lca, target, distance, comparator(original.Key, target.Key), false)
}

//******************************************************************//
//                             Iterator                             //
//******************************************************************//

// Begin returns an initialized iterator, which points to one element before it's first.
// Unless Next() is called, the iterator is in an invalid state.

func (tree *Tree[TKey, TValue]) OrderedBegin() ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return tree.NewOrderedIterator(-1, tree.Size())
}

// End returns an initialized iterator, which points to one element afrer it's last.
// Unless Previous() is called, the iterator is in an invalid state.

func (tree *Tree[TKey, TValue]) OrderedEnd() ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return tree.NewOrderedIterator(tree.Size(), tree.Size())
}

// First returns an initialized iterator, which points to it's first element.

func (tree *Tree[TKey, TValue]) OrderedFirst() ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return tree.NewOrderedIterator(0, tree.Size())
}

// Last returns an initialized iterator, which points to it's last element.

func (tree *Tree[TKey, TValue]) OrderedLast() ds.ReadWriteOrdCompBidRandCollIterator[TKey, TValue] {
	return tree.NewOrderedIterator(tree.Size()-1, tree.Size())
}
