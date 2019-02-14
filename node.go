package tree

import "strings"

// Node is a node within the tree, which includes its parents and children.
type Node struct {
	Item     Item
	Parents  []*Node
	Children []*Node
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
func (tn *Node) Ascendants() (ascendants Nodes) {
	for _, p := range tn.Parents {
		ascendants = append(ascendants, p)
		ascendants = append(ascendants, p.Ascendants()...)
	}
	return
}

// Descendants is a list of all nodes below the current node in the hierarchy.
func (tn *Node) Descendants() (descendants Nodes) {
	for _, c := range tn.Children {
		descendants = append(descendants, c)
		descendants = append(descendants, c.Descendants()...)
	}
	return
}

// Nodes is a slice of Nodes which can be sorted by the level in the hierarchy
// and alphabetically when nodes are at the same level.
type Nodes []*Node

func (tn Nodes) Len() int      { return len(tn) }
func (tn Nodes) Swap(i, j int) { tn[i], tn[j] = tn[j], tn[i] }
func (tn Nodes) Less(i, j int) bool {
	il, jl := tn[i].AscendantLevels(0), tn[j].AscendantLevels(0)
	if il < jl {
		return true
	}
	if il == jl {
		return strings.Compare(tn[i].Item.Name(), tn[j].Item.Name()) < 0
	}
	return false
}
