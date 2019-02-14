package tree

import "strings"

// TreeNode is a node within the tree, which includes its parents and children.
type TreeNode struct {
	Node     Item
	Parents  []*TreeNode
	Children []*TreeNode
}

// AscendantLevels returns how many levels there are above the current TreeNode.
// Pass zero in, if to count upwards.
func (tn *TreeNode) AscendantLevels(current int) int {
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
func (tn *TreeNode) Ascendants() (ascendants TreeNodes) {
	for _, p := range tn.Parents {
		ascendants = append(ascendants, p)
		ascendants = append(ascendants, p.Ascendants()...)
	}
	return
}

// Descendants is a list of all nodes below the current node in the hierarchy.
func (tn *TreeNode) Descendants() (descendants TreeNodes) {
	for _, c := range tn.Children {
		descendants = append(descendants, c)
		descendants = append(descendants, c.Descendants()...)
	}
	return
}

// TreeNodes is a slice of TreeNodes which can be sorted by the level in the hierarchy
// and alphabetically when nodes are at the same level.
type TreeNodes []*TreeNode

func (tn TreeNodes) Len() int      { return len(tn) }
func (tn TreeNodes) Swap(i, j int) { tn[i], tn[j] = tn[j], tn[i] }
func (tn TreeNodes) Less(i, j int) bool {
	il, jl := tn[i].AscendantLevels(0), tn[j].AscendantLevels(0)
	if il < jl {
		return true
	}
	if il == jl {
		return strings.Compare(tn[i].Node.Name(), tn[j].Node.Name()) < 0
	}
	return false
}
