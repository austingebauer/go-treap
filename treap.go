// Package Treap provides a balanced binary search tree
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
	t.insert(value, rand.Uint64())
}

// insert inserts a node with the passed value and priority into the Treap.
func (t *Treap) insert(value string, priority uint64) {
	n := &node{
		value:    value,
		priority: priority,
	}

	if t.root == nil {
		t.root = n
		return
	}

	// TODO: Not implemented
}

// Delete delete the given value from the Treap.
func (t *Treap) Delete(value string) {
	// TODO: Not implemented
}

// rotateRight
func rotateRight(root, pivot *node) {
	root.left = pivot.right
	pivot.right = root
}

// rotateLeft
func rotateLeft(root, pivot *node) {
	root.right = pivot.left
	pivot.left = root
}
