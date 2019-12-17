package go_treap

type Treap struct {
	root *node
}

type node struct {
	value string
	priority int
}

func NewTreap() *Treap {
	return &Treap{}
}
