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
		{
			name:     "multiple levels",
			tree:     multipleLevels(),
			expected: []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"},
		},
	}
	for _, test := range tests {
		test := test
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
	t.AddParents(StringItem("child"), StringItem("parent"))
	return t
}

func singleParentTwoChildren() *Tree {
	t := New()
	t.AddParents(StringItem("child1"), StringItem("parent"))
	t.AddParents(StringItem("child2"), StringItem("parent"))
	return t
}

func singleChildTwoParents() *Tree {
	t := New()
	t.AddParents(StringItem("child1"), StringItem("parentA"))
	t.AddParents(StringItem("child1"), StringItem("parentB"))
	return t
}

func multipleLevels() *Tree {
	//    A  B
	//   C  D
	//  E  F
	// G H  I
	t := New()
	t.AddParents(StringItem("G"), StringItem("E"))
	t.AddParents(StringItem("H"), StringItem("E"), StringItem("F"))
	t.AddParents(StringItem("I"), StringItem("F"))
	t.AddParents(StringItem("E"), StringItem("C"))
	t.AddParents(StringItem("F"), StringItem("D"))
	t.AddParents(StringItem("C"), StringItem("A"))
	t.AddParents(StringItem("D"), StringItem("B"))
	return t
}

func BenchmarkMultipleLevels(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		tree := multipleLevels()
		tree.Sorted()
	}
}
