package tree

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		tree     *Tree
		expected []string
	}{
		{
			name:     "single parent, single child",
			tree:     singleParentSingleChild(),
			expected: []string{"parent", "child"},
		},
		{
			name:     "single parent, two children",
			tree:     singleParentTwoChildren(),
			expected: []string{"parent", "child1", "child2"},
		},
		{
			name:     "single child, two parents",
			tree:     singleChildTwoParents(),
			expected: []string{"parentA", "parentB", "child1"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			sorted := test.tree.Sorted()
			actual := getNames(sorted)
			if !reflect.DeepEqual(test.expected, actual) {
				t.Errorf("expected: '%v', got: '%v'", test.expected, actual)
			}
		})
	}
}

func getNames(nodes []Item) (s []string) {
	s = make([]string, len(nodes))
	for i, n := range nodes {
		s[i] = n.Name()
	}
	return
}

func singleParentSingleChild() *Tree {
	t := New()
	t.AddParents(StringNode("child"), StringNode("parent"))
	return t
}

func singleParentTwoChildren() *Tree {
	t := New()
	t.AddParents(StringNode("child1"), StringNode("parent"))
	t.AddParents(StringNode("child2"), StringNode("parent"))
	return t
}

func singleChildTwoParents() *Tree {
	t := New()
	t.AddParents(StringNode("child1"), StringNode("parentA"))
	t.AddParents(StringNode("child1"), StringNode("parentB"))
	return t
}
