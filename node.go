package tree

import "strings"

// Node is a node within the tree, which includes its parents and children.
type Node struct {
	Item    Item
	Parents []*Node
}

// AscendantLevels returns how many levels there are above the current Node.
// Pass zero in, if to count upwards.
func (tn *Node) AscendantLevels(current int) int {
	if len(tn.Parents) == 0 {
		return current
	}
	vals := make([]int, len(tn.Parents))
	for i, p := range tn.Parents {
		vals[i] = p.AscendantLevels(current + 1)
	}
	var max int
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

// Ascendants is a list of all nodes above the current node in the hierarchy.
func (tn *Node) Ascendants() (ascendants []*Node) {
	for _, p := range tn.Parents {
		ascendants = append(ascendants, p)
		ascendants = append(ascendants, p.Ascendants()...)
	}
	return
}

// NewNodeSorter creates a new node sorter which can organise the nodes parents first.
func NewNodeSorter(nodes []*Node) (ns *NodeSorter) {
	ns = &NodeSorter{
		Nodes:     nodes,
		nodeDepth: make(map[*Node]int, len(nodes)),
	}
	for _, n := range nodes {
		ns.nodeDepth[n] = n.AscendantLevels(0)
	}
	return ns
}

// NodeSorter sorts a slice of Nodes which can be sorted by the level in the hierarchy
// and alphabetically when nodes are at the same level.
type NodeSorter struct {
	Nodes     []*Node
	nodeDepth map[*Node]int
}

func (ns NodeSorter) Len() int      { return len(ns.Nodes) }
func (ns NodeSorter) Swap(i, j int) { ns.Nodes[i], ns.Nodes[j] = ns.Nodes[j], ns.Nodes[i] }
func (ns NodeSorter) Less(i, j int) bool {
	ni, nj := ns.Nodes[i], ns.Nodes[j]
	il, jl := ns.nodeDepth[ni], ns.nodeDepth[nj]
	if il < jl {
		return true
	}
	if il == jl {
		return strings.Compare(ni.Item.Name(), nj.Item.Name()) < 0
	}
	return false
}
