// Package Treap provides a balanced binary search tree
// by using binary heap properties and randomization.
package treap

// Treap is a balanced binary search tree.
type Treap struct {
	root *node
}

// node represents a value and its priority in a Treap.
type node struct {
	value    string
	priority int
	left     *node
	right    *node
}

// NewTreap returns a new Treap.
func NewTreap() *Treap {
	return &Treap{}
}

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

func (t *Treap) Insert(value string) {
	// TODO: Not implemented
}

func (t *Treap) Delete(value string) {
	// TODO: Not implemented
}
