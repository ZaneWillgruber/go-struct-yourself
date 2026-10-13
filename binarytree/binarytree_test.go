package binarytree

import (
	"slices"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	t.Run("Create new binary tree", func(t *testing.T) {
		tree := NewBinaryTree()
		if tree.Root != nil {
			t.Errorf("Expected root to be nil, got %v", tree.Root)
		}
	})
}

func TestInsert(t *testing.T) {
	t.Run("Insert into empty tree", func(t *testing.T) {
		tree := NewBinaryTree()
		tree.Insert(5)

		if tree.Root == nil {
			t.Errorf("Expected root to be non-nil, got nil")
		}

		if tree.Root.value != 5 {
			t.Errorf("Expected root value to be 5, got %v", tree.Root.value)
		}

		if tree.Root.left != nil {
			t.Errorf("Expected left child to be nil, got %v", tree.Root.left)
		}

		if tree.Root.right != nil {
			t.Errorf("Expected right child to be nil, got %v", tree.Root.right)
		}
	})

	t.Run("Insert into left", func(t *testing.T) {
		tree := NewBinaryTree()
		tree.Insert(5)
		tree.Insert(3)

		if tree.Root.left == nil {
			t.Errorf("Expected left child to be non-nil, got nil")
		}

		if tree.Root.left.value != 3 {
			t.Errorf("Expected left child value to be 3, got %v", tree.Root.left.value)
		}
	})

	t.Run("Insert into right", func(t *testing.T) {
		tree := NewBinaryTree()
		tree.Insert(5)
		tree.Insert(7)

		if tree.Root.right == nil {
			t.Errorf("Expected right child to be non-nil, got nil")
		}

		if tree.Root.right.value != 7 {
			t.Errorf("Expected right child value to be 7, got %v", tree.Root.right.value)
		}
	})

	t.Run("Insert multiple values", func(t *testing.T) {
		tree := NewBinaryTree()
		values := []int{5, 3, 7, 2, 4, 6, 8}
		for _, v := range values {
			tree.Insert(v)
		}

		if tree.Root.value != 5 {
			t.Errorf("Expected root value to be 5, got %v", tree.Root.value)
		}

		if tree.Root.left.value != 3 {
			t.Errorf("Expected left child value to be 3, got %v", tree.Root.left.value)
		}

		if tree.Root.right.value != 7 {
			t.Errorf("Expected right child value to be 7, got %v", tree.Root.right.value)
		}

		if tree.Root.left.left.value != 2 {
			t.Errorf("Expected left-left child value to be 2, got %v", tree.Root.left.left.value)
		}

		if tree.Root.left.right.value != 4 {
			t.Errorf("Expected left-right child value to be 4, got %v", tree.Root.left.right.value)
		}

		if tree.Root.right.left.value != 6 {
			t.Errorf("Expected right-left child value to be 6, got %v", tree.Root.right.left.value)
		}

		if tree.Root.right.right.value != 8 {
			t.Errorf("Expected right-right child value to be 8, got %v", tree.Root.right.right.value)
		}
	})
}

type traversalTestCase struct {
	name     string
	values   []int
	expected []int
}

func testTraversal(t *testing.T, traversalOrder TraversalOrder, cases []traversalTestCase) {

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tree := NewBinaryTree()
			for _, v := range tc.values {
				tree.Insert(v)
			}

			result := tree.Traverse(traversalOrder)

			if len(result) != len(tc.expected) {
				t.Errorf("Expected result length %d, got %d", len(tc.expected), len(result))
			}

			for i, v := range tc.expected {
				if result[i] != v {
					t.Errorf("Expected result[%d] to be %d, got %d", i, v, result[i])
				}
			}
		})
	}
}

func TestInOrderTraversal(t *testing.T) {
	cases := []traversalTestCase{
		{
			name:     "Empty tree",
			values:   []int{},
			expected: []int{},
		},
		{
			name:     "Single node",
			values:   []int{5},
			expected: []int{5},
		},
		{
			name:     "Multiple nodes",
			values:   []int{5, 3, 7, 2, 4, 6, 8},
			expected: []int{2, 3, 4, 5, 6, 7, 8},
		},
	}

	testTraversal(t, InOrder, cases)
}

func TestPreOrderTraversal(t *testing.T) {
	cases := []traversalTestCase{
		{
			name:     "Empty tree",
			values:   []int{},
			expected: []int{},
		},
		{
			name:     "Single node",
			values:   []int{5},
			expected: []int{5},
		},
		{
			name:     "Multiple nodes",
			values:   []int{5, 3, 7, 2, 4, 6, 8},
			expected: []int{5, 3, 2, 4, 7, 6, 8},
		},
	}

	testTraversal(t, PreOrder, cases)
}

func TestPostOrderTraversal(t *testing.T) {
	cases := []traversalTestCase{
		{
			name:     "Empty tree",
			values:   []int{},
			expected: []int{},
		},
		{
			name:     "Single node",
			values:   []int{5},
			expected: []int{5},
		},
		{
			name:     "Multiple nodes",
			values:   []int{5, 3, 7, 2, 4, 6, 8},
			expected: []int{2, 4, 3, 6, 8, 7, 5},
		},
	}

	testTraversal(t, PostOrder, cases)
}

func TestMin(t *testing.T) {
	type minResult struct {
		value int
		found bool
	}

	cases := []struct {
		name     string
		values   []int
		expected minResult
	}{
		{
			name:   "Empty tree",
			values: []int{},
			expected: minResult{
				value: 0,
				found: false,
			},
		},
		{
			name:   "Single node",
			values: []int{5},
			expected: minResult{
				value: 5,
				found: true,
			},
		},
		{
			name:   "Multiple nodes",
			values: []int{5, 3, 7, 2, 4, 6, 8},
			expected: minResult{
				value: 2,
				found: true,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tree := NewBinaryTree()
			for _, num := range tc.values {
				tree.Insert(num)
			}

			got, found := tree.Min()

			if got != tc.expected.value {
				t.Errorf("Expected min value %d, got %d", tc.expected.value, got)
			}

			if found != tc.expected.found {
				t.Errorf("Expected found to be %v, got %v", tc.expected.found, found)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type maxResult struct {
		value int
		found bool
	}

	cases := []struct {
		name     string
		values   []int
		expected maxResult
	}{
		{
			name:   "Empty tree",
			values: []int{},
			expected: maxResult{
				value: 0,
				found: false,
			},
		},
		{
			name:   "Single node",
			values: []int{5},
			expected: maxResult{
				value: 5,
				found: true,
			},
		},
		{
			name:   "Multiple nodes",
			values: []int{5, 3, 7, 2, 4, 6, 8},
			expected: maxResult{
				value: 8,
				found: true,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tree := NewBinaryTree()
			for _, num := range tc.values {
				tree.Insert(num)
			}

			got, found := tree.Max()

			if got != tc.expected.value {
				t.Errorf("Expected max value %d, got %d", tc.expected.value, got)
			}

			if found != tc.expected.found {
				t.Errorf("Expected found to be %v, got %v", tc.expected.found, found)
			}
		})
	}
}

func TestFind(t *testing.T) {
	cases := []struct {
		name     string
		values   []int
		search   int
		expected bool
	}{
		{
			name:     "Empty tree",
			values:   []int{},
			search:   5,
			expected: false,
		},
		{
			name:     "Single node found",
			values:   []int{5},
			search:   5,
			expected: true,
		},
		{
			name:     "Single node not found",
			values:   []int{5},
			search:   3,
			expected: false,
		},
		{
			name:     "Multiple nodes found",
			values:   []int{5, 3, 7, 2, 4, 6, 8},
			search:   4,
			expected: true,
		},
		{
			name:     "Multiple nodes not found",
			values:   []int{5, 3, 7, 2, 4, 6, 8},
			search:   10,
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tree := NewBinaryTree()
			for _, num := range tc.values {
				tree.Insert(num)
			}

			node := tree.Find(tc.search)

			if (node != nil) != tc.expected {
				t.Errorf("Expected find result to be %v, got %v", tc.expected, node != nil)
			}
		})
	}
}

func TestSuccessor(t *testing.T) {
	cases := []struct {
		name     string
		values   []int
		search   int
		expected *int
	}{
		{
			name:     "Empty tree",
			values:   []int{},
			search:   5,
			expected: nil,
		},
		{
			name:     "Single node",
			values:   []int{5},
			search:   5,
			expected: nil,
		},
		{
			name:     "Multiple nodes with successor",
			values:   []int{5, 3, 7, 2, 4, 6, 8},
			search:   4,
			expected: new(5),
		},
		{
			name:     "Multiple nodes without successor",
			values:   []int{5, 3, 7, 2, 4, 6, 8},
			search:   8,
			expected: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tree := NewBinaryTree()
			for _, num := range tc.values {
				tree.Insert(num)
			}

			successorNode := tree.Successor(tc.search)

			if successorNode == nil && tc.expected != nil {
				t.Errorf("Expected successor to be %v, got nil", *tc.expected)
			} else if successorNode != nil && tc.expected == nil {
				t.Errorf("Expected successor to be nil, got %v", successorNode.value)
			} else if successorNode != nil && tc.expected != nil && successorNode.value != *tc.expected {
				t.Errorf("Expected successor to be %v, got %v", *tc.expected, successorNode.value)
			}
		})
	}
}

func TestPredecessor(t *testing.T) {
	cases := []struct {
		name     string
		values   []int
		search   int
		expected *int
	}{
		{
			name:     "Empty tree",
			values:   []int{},
			search:   5,
			expected: nil,
		},
		{
			name:     "Single node",
			values:   []int{5},
			search:   5,
			expected: nil,
		},
		{
			name:     "Multiple nodes with predecessor",
			values:   []int{5, 3, 7, 2, 4, 6, 8},
			search:   4,
			expected: new(3),
		},
		{
			name:     "Multiple nodes without predecessor",
			values:   []int{5, 3, 7, 2, 4, 6, 8},
			search:   2,
			expected: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tree := NewBinaryTree()
			for _, num := range tc.values {
				tree.Insert(num)
			}

			predecessorNode := tree.Predecessor(tc.search)

			if predecessorNode == nil && tc.expected != nil {
				t.Errorf("Expected predecessor to be %v, got nil", *tc.expected)
			} else if predecessorNode != nil && tc.expected == nil {
				t.Errorf("Expected predecessor to be nil, got %v", predecessorNode.value)
			} else if predecessorNode != nil && tc.expected != nil && predecessorNode.value != *tc.expected {
				t.Errorf("Expected predecessor to be %v, got %v", *tc.expected, predecessorNode.value)
			}
		})
	}
}

func checkParents(t *testing.T, n *node, parent *node) {
	t.Helper()
	if n == nil {
		return
	}

	if n.parent != parent {
		t.Errorf("Expected parent of %v to be %v, got %v", n.value, parent, n.parent)
	}

	checkParents(t, n.left, n)
	checkParents(t, n.right, n)
}

func TestDelete(t *testing.T) {
	cases := []struct {
		name     string
		values   []int
		delete   int
		deleted  bool
		expected []int // pre-order, so the tree's shape is checked too
	}{
		{
			name:     "Empty tree",
			values:   []int{},
			delete:   5,
			deleted:  false,
			expected: []int{},
		},
		{
			name:     "Single node",
			values:   []int{5},
			delete:   5,
			deleted:  true,
			expected: []int{},
		},
		{
			name:     "Value not found",
			values:   []int{5, 3, 7, 2, 4, 6, 8},
			delete:   10,
			deleted:  false,
			expected: []int{5, 3, 2, 4, 7, 6, 8},
		},
		{
			name:     "Leaf node",
			values:   []int{5, 3, 7, 2, 4, 6, 8},
			delete:   2,
			deleted:  true,
			expected: []int{5, 3, 4, 7, 6, 8},
		},
		{
			name:     "Node with only left child",
			values:   []int{5, 3, 7, 2},
			delete:   3,
			deleted:  true,
			expected: []int{5, 2, 7},
		},
		{
			name:     "Node with only right child",
			values:   []int{5, 3, 7, 8},
			delete:   7,
			deleted:  true,
			expected: []int{5, 3, 8},
		},
		{
			name:     "Root with one child",
			values:   []int{5, 7, 6, 8},
			delete:   5,
			deleted:  true,
			expected: []int{7, 6, 8},
		},
		{
			name:     "Node with two children",
			values:   []int{5, 3, 7, 2, 4, 6, 8},
			delete:   3,
			deleted:  true,
			expected: []int{5, 4, 2, 7, 6, 8},
		},
		{
			name:     "Root with two children, successor has right child",
			values:   []int{5, 3, 8, 6, 9, 7},
			delete:   5,
			deleted:  true,
			expected: []int{6, 3, 8, 7, 9},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tree := NewBinaryTree()
			for _, num := range tc.values {
				tree.Insert(num)
			}

			deleted := tree.Delete(tc.delete)
			if deleted != tc.deleted {
				t.Errorf("Expected delete result to be %v, got %v", tc.deleted, deleted)
			}

			result := tree.Traverse(PreOrder)
			if !slices.Equal(result, tc.expected) {
				t.Errorf("Expected pre-order %v, got %v", tc.expected, result)
			}

			checkParents(t, tree.Root, nil)
		})
	}

	t.Run("Delete every node", func(t *testing.T) {
		values := []int{5, 3, 7, 2, 4, 6, 8}
		tree := NewBinaryTree()
		for _, num := range values {
			tree.Insert(num)
		}

		for _, num := range values {
			if !tree.Delete(num) {
				t.Fatalf("Expected delete of %v to succeed", num)
			}
			checkParents(t, tree.Root, nil)
		}

		if tree.Root != nil {
			t.Errorf("Expected root to be nil, got %v", tree.Root)
		}
	})
}
