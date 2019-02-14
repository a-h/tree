package tree

// An Item to add to a tree.
type Item interface {
	Name() string
}

type StringNode string

func (sn StringNode) Name() string {
	return string(sn)
}
