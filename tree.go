package tree

import (
	"sort"
	"sync"
)

// New creates a new tree.
func New() *Tree {
	return &Tree{
		nodes:   make(map[string]Item),
		parents: make(map[string][]string),
	}
}

// Tree of nodes.
type Tree struct {
	nodes   map[string]Item
	parents map[string][]string
	m       sync.Mutex
}

// AddItems adds items to the tree.
func (t *Tree) AddItems(nodes ...Item) {
	t.m.Lock()
	defer t.m.Unlock()
	for _, n := range nodes {
		if _, exists := t.nodes[n.Name()]; exists {
			continue
		}
		t.nodes[n.Name()] = n
		t.parents[n.Name()] = []string{}
	}
}

// AddParents adds both items (if required) then adds a relationship from the child to the parent.
func (t *Tree) AddParents(from Item, to ...Item) {
	t.AddItems(from)
	t.AddItems(to...)
	t.m.Lock()
	defer t.m.Unlock()
	for _, tt := range to {
		t.parents[from.Name()] = append(t.parents[from.Name()], tt.Name())
	}
}

// Nodes returns all available nodes, in a random order.
func (t *Tree) Nodes() (tn []*Node) {
	tn = make([]*Node, len(t.nodes))
	var i int
	for name := range t.nodes {
		n, _ := t.GetNode(name)
		tn[i] = n
		i++
	}
	return
}

// GetNodes gets nodes by their names. It will return false if the node cannot be found.
func (t *Tree) GetNodes(names ...string) (tn []*Node, ok bool) {
	ok = true
	tn = make([]*Node, len(names))
	for i, n := range names {
		tn[i], ok = t.GetNode(n)
		if !ok {
			return
		}
	}
	return
}

// GetNode returns a node by its name.
func (t *Tree) GetNode(name string) (tn *Node, ok bool) {
	n, ok := t.nodes[name]
	if !ok {
		return
	}
	parents, ok := t.GetNodes(t.parents[name]...)
	if !ok {
		return
	}
	tn = &Node{
		Item:    n,
		Parents: parents,
	}
	return
}

// Sorted returns the nodes in sorted order, where the nodes with no parents come first, then
// worked through the levels of the tree.
func (t *Tree) Sorted() (items []Item) {
	ns := NewNodeSorter(t.Nodes())
	sort.Sort(ns)
	items = make([]Item, len(ns.Nodes))
	for i, nn := range ns.Nodes {
		items[i] = nn.Item
	}
	return
}
