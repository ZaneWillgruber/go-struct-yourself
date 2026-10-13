package binarytree

// Node
type node struct {
	parent *node
	left   *node
	right  *node
	value  int
}

type TraversalOrder int

const (
	InOrder TraversalOrder = iota
	PreOrder
	PostOrder
)

func newNode(value int) *node {
	return &node{value: value}
}

func (n *node) isLeft() bool {
	return n.parent.left == n
}

func (n *node) isRight() bool {
	return !n.isLeft()
}

func (n *node) insert(value int) {
	if value < n.value {
		if n.left == nil {
			n.left = newNode(value)
			n.left.parent = n
			return
		} else {
			n.left.insert(value)
			return
		}
	}

	if n.right == nil {
		n.right = newNode(value)
		n.right.parent = n
		return
	} else {
		n.right.insert(value)
		return
	}
}

func (n *node) find(value int) *node {
	if n.value == value {
		return n
	}

	if n.left != nil && value < n.value {
		return n.left.find(value)
	}

	if n.right != nil && value > n.value {
		return n.right.find(value)
	}

	return nil
}

func (n *node) successor() *node {
	if n.right != nil {
		return n.right.min()
	}

	parent := n.parent
	currentNode := n
	for parent != nil && currentNode == parent.right {
		currentNode = parent
		parent = parent.parent
	}

	return parent
}

func (n *node) predecessor() *node {
	if n.left != nil {
		return n.left.max()
	}

	parent := n.parent
	currentNode := n
	for parent != nil && currentNode == parent.left {
		currentNode = parent
		parent = parent.parent
	}

	return parent
}

func (n *node) traverse(order TraversalOrder) []int {
	var result []int

	var walk func(*node)
	walk = func(current *node) {
		if current == nil {
			return
		}

		if order == PreOrder {
			result = append(result, current.value)
		}

		walk(current.left)

		if order == InOrder {
			result = append(result, current.value)
		}

		walk(current.right)

		if order == PostOrder {
			result = append(result, current.value)
		}
	}

	walk(n)
	return result
}

func (n *node) min() *node {
	if n.left != nil {
		return n.left.min()
	}

	return n
}

func (n *node) max() *node {
	if n.right != nil {
		return n.right.max()
	}

	return n
}

// Tree
type BinaryTree struct {
	Root *node
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) Insert(value int) {
	if t.Root == nil {
		t.Root = newNode(value)
	} else {
		t.Root.insert(value)
	}
}

func (t *BinaryTree) Find(value int) *node {
	if t.Root == nil {
		return nil
	}

	return t.Root.find(value)
}

func (t *BinaryTree) Successor(value int) *node {
	n := t.Find(value)
	if n == nil {
		return nil
	}

	return n.successor()
}

func (t *BinaryTree) Predecessor(value int) *node {
	n := t.Find(value)
	if n == nil {
		return nil
	}

	return n.predecessor()
}

func (t *BinaryTree) Delete(value int) bool {
	n := t.Find(value)
	if n == nil {
		return false
	}

	if n.left == nil && n.right == nil {
		if n.parent == nil {
			t.Root = nil
		} else if n.isLeft() {
			n.parent.left = nil
		} else {
			n.parent.right = nil
		}
		return true
	}

	if n.left != nil && n.right == nil {
		n.left.parent = n.parent
		if n.parent == nil {
			t.Root = n.left
		} else if n.isLeft() {
			n.parent.left = n.left
		} else {
			n.parent.right = n.left
		}
		return true
	}

	if n.right != nil && n.left == nil {
		n.right.parent = n.parent
		if n.parent == nil {
			t.Root = n.right
		} else if n.isLeft() {
			n.parent.left = n.right
		} else {
			n.parent.right = n.right
		}
		return true
	}

	successor := n.successor()
	n.value = successor.value

	if successor.right != nil {
		successor.right.parent = successor.parent
	}
	if successor.isLeft() {
		successor.parent.left = successor.right
	} else {
		successor.parent.right = successor.right
	}
	return true
}

func (t *BinaryTree) Traverse(order TraversalOrder) []int {
	if t.Root == nil {
		return []int{}
	}

	return t.Root.traverse(order)
}

func (t *BinaryTree) Min() (int, bool) {
	if t.Root == nil {
		return 0, false
	}

	return t.Root.min().value, true
}

func (t *BinaryTree) Max() (int, bool) {
	if t.Root == nil {
		return 0, false
	}

	return t.Root.max().value, true
}
