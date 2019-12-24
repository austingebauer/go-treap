// Package treap provides a balanced binary search tree
// by using binary heap properties and randomization.
package treap

import (
	"math/rand"
	"time"
)

// Treap is a balanced binary search tree.
type Treap struct {
	root *node
}

// node represents a value and its priority in a Treap.
type node struct {
	value    string
	priority uint64
	left     *node
	right    *node
}

// NewTreap returns a new Treap.
func NewTreap() *Treap {
	return &Treap{}
}

// Search returns true if the given value is in the Treap.
// Otherwise, returns false.
func (t *Treap) Search(value string) bool {
	if t.root == nil {
		return false
	}

	c := t.root
	for c != nil {
		if value == c.value {
			return true
		}

		if value < c.value {
			c = c.left
		} else {
			c = c.right
		}
	}

	return false
}

// Insert inserts the given value into the Treap.
func (t *Treap) Insert(value string) {
	rand.Seed(time.Now().UnixNano())
	t.root = insert(t.root, value, rand.Uint64())
}

// insert inserts a node with the passed value and priority into the Treap.
func insert(root *node, value string, priority uint64) *node {
	if root == nil {
		return &node{
			value:    value,
			priority: priority,
		}
	}

	if value < root.value {
		root.left = insert(root.left, value, priority)
		if root.priority < root.left.priority {
			root = rotateRight(root, root.left)
		}
	} else {
		root.right = insert(root.right, value, priority)
		if root.priority < root.right.priority {
			root = rotateLeft(root, root.right)
		}
	}

	return root
}

// Delete delete the given value from the Treap.
func (t *Treap) Delete(value string) {
	// TODO: Not implemented
}

// rotateRight does a tree rotation to the right given the passed root and pivot.
// After the rotation, the root will be the right child of the pivot.
// The pivot will be returned.
func rotateRight(root, pivot *node) *node {
	root.left = pivot.right
	pivot.right = root
	return pivot
}

// rotateLeft does a tree rotation to the left given the passed root and pivot.
// After the rotation, the root will be the left child of the pivot.
// The pivot will be returned.
func rotateLeft(root, pivot *node) *node {
	root.right = pivot.left
	pivot.left = root
	return pivot
}
