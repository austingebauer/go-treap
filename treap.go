// Package treap provides a balanced binary search tree
// by using binary heap properties and randomization.
package treap

import (
	"math"
	"math/rand"
	"time"
)

const (
	// Random priorities in the Treap will be in the range of [1, MaxInt64].
	maxPriority = math.MaxInt64
	minPriority = 1

	// The priority given to Treap nodes during deletion.
	deletePriority = 0
)

// Treap is a balanced binary search tree.
type Treap struct {
	root *node
}

// node represents a value and its priority in a Treap.
type node struct {
	value    string
	priority int64
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

	return binarySearch(t.root, value) != nil
}

// Insert inserts the given value into the Treap.
func (t *Treap) Insert(value string) {
	rand.Seed(time.Now().UnixNano())
	t.root = insert(t.root, value,
		rand.Int63n(maxPriority-minPriority)+minPriority)
}

// insert inserts a node with the passed value and priority into the Treap.
func insert(n *node, value string, priority int64) *node {
	if n == nil {
		return &node{
			value:    value,
			priority: priority,
		}
	}

	if value == n.value {
		return n
	} else if value < n.value {
		n.left = insert(n.left, value, priority)
		if n.priority < n.left.priority {
			n = rotateRight(n, n.left)
		}
	} else {
		n.right = insert(n.right, value, priority)
		if n.priority < n.right.priority {
			n = rotateLeft(n, n.right)
		}
	}

	return n
}

// Delete deletes the given value from the Treap.
func (t *Treap) Delete(value string) {
	t.root = delete(t.root, value)
}

// delete finds and deletes the node with the given value from the Treap.
func delete(n *node, value string) *node {
	if n == nil {
		return nil
	}

	// delete the node with value after it's been rotated down to a leaf
	if n.left == nil && n.right == nil && n.value == value {
		return nil
	}

	if n.value == value {
		n.priority = deletePriority

		if n.right == nil && n.left != nil {
			pivot := rotateRight(n, n.left)
			pivot.right = delete(n, value)
			return pivot
		} else if n.left == nil && n.right != nil {
			pivot := rotateLeft(n, n.right)
			pivot.left = delete(n, value)
			return pivot
		} else if n.right.priority > n.left.priority {
			pivot := rotateLeft(n, n.right)
			pivot.left = delete(n, value)
			return pivot
		} else {
			pivot := rotateRight(n, n.left)
			pivot.right = delete(n, value)
			return pivot
		}
	}

	if value < n.value {
		n.left = delete(n.left, value)
	} else {
		n.right = delete(n.right, value)
	}

	return n
}

// binarySearch performs a binary search starting from the
// passed node for the passed value.
// If the passed value is found, a pointer to the node with
// the value is returned. Otherwise, nil is returned.
func binarySearch(n *node, value string) *node {
	for n != nil {
		if value == n.value {
			return n
		}

		if value < n.value {
			n = n.left
		} else {
			n = n.right
		}
	}

	return nil
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
