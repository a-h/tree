package tree

// An Item to add to a tree.
type Item interface {
	Name() string
}

// A StringItem can be added to the tree.
type StringItem string

// Name of the StringItem is the value of the underlying string.
func (sn StringItem) Name() string {
	return string(sn)
}
