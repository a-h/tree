package tree

import (
	"sort"
	"sync"
)

// New creates a new tree.
func New() *Tree {
	return &Tree{
		nodes:    make(map[string]Item),
		parents:  make(map[string][]string),
		children: make(map[string][]string),
	}
}

// Tree of nodes.
type Tree struct {
	nodes    map[string]Item
	parents  map[string][]string
	children map[string][]string
	m        sync.Mutex
}

// AddNodes adds nodes to the tree.
func (t *Tree) AddNodes(nodes ...Item) {
	t.m.Lock()
	defer t.m.Unlock()
	for _, n := range nodes {
		if _, exists := t.nodes[n.Name()]; exists {
			continue
		}
		t.nodes[n.Name()] = n
		t.parents[n.Name()] = []string{}
		t.children[n.Name()] = []string{}
	}
}

// AddParents adds both items (if required) then adds a relationship from the child to the parent.
func (t *Tree) AddParents(from Item, to ...Item) {
	t.AddNodes(from)
	t.AddNodes(to...)
	t.m.Lock()
	defer t.m.Unlock()
	for _, tt := range to {
		edges := t.children[from.Name()]
		t.children[from.Name()] = append(edges, tt.Name())
	}
	t.parents[from.Name()] = append(t.parents[from.Name()], nodeNames(to)...)
}

func nodeNames(nodes []Item) (names []string) {
	names = make([]string, len(nodes))
	for i, n := range nodes {
		names[i] = n.Name()
	}
	return
}

// Nodes returns all available nodes, in a random order.
func (t *Tree) Nodes() (tn TreeNodes) {
	for name := range t.nodes {
		n, _ := t.getNode(name)
		tn = append(tn, n)
	}
	return
}

// GetNodes gets nodes by their names. It will return false if the node cannot be found.
func (t *Tree) GetNodes(names ...string) (tn TreeNodes, ok bool) {
	ok = true
	tn = make(TreeNodes, len(names))
	for i, n := range names {
		tn[i], ok = t.getNode(n)
		if !ok {
			return
		}
	}
	return
}

func (t *Tree) getNode(name string) (tn *TreeNode, ok bool) {
	n, ok := t.nodes[name]
	if !ok {
		return
	}
	parents, ok := t.GetNodes(t.parents[name]...)
	if !ok {
		return
	}
	children, ok := t.GetNodes(t.children[name]...)
	if !ok {
		return
	}
	tn = &TreeNode{
		Node:     n,
		Children: children,
		Parents:  parents,
	}
	return
}

// Sorted returns the nodes in sorted order, where the nodes with no parents come first, then
// worked through the levels of the tree.
func (t *Tree) Sorted() (nodes []Item) {
	n := t.Nodes()
	sort.Sort(n)
	nodes = make([]Item, len(n))
	for i, nn := range n {
		nodes[i] = nn.Node
	}
	return
}
