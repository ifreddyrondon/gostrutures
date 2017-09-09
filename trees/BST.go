package trees

// A single node that compose a tree
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil, nil}
}

// Binary search tree
type BST struct {
	Root *Node
}

// Constructor, return a Binary Search Tree with the root
func New(value int) *BST {
	return &BST{NewNode(value)}
}

// Insert insert an item in the right position in the tree. Return true if the value was inserted and false otherwise
func (t *BST) Insert(value int) bool {
	node := NewNode(value)
	if t.Root == nil {
		t.Root = node
		return true
	}
	return insertNode(t.Root, node)
}

func insertNode(root, newNode *Node) bool {
	if root.Value == newNode.Value {
		return false
	}

	if newNode.Value < root.Value {
		if root.Left == nil {
			root.Left = newNode
			return true
		}
		return insertNode(root.Left, newNode)
	} else {
		if root.Right == nil {
			root.Right = newNode
			return true
		}
		return insertNode(root.Right, newNode)
	}
}

// InOrderTraverse visits all the nodes in order
func (t *BST) InOrderTraverse(f func(int)) {
	if t.Root == nil {
		return
	}

	inOrderTraverse(t.Root, f)
}

// internal recursive function to traverse in order
func inOrderTraverse(node *Node, f func(int)) {
	if node == nil {
		return
	}

	inOrderTraverse(node.Left, f)
	f(node.Value)
	inOrderTraverse(node.Right, f)
}

// PreOrderTraverse visits all the nodes in pre order
func (t *BST) PreOrderTraverse(f func(int)) {
	if t.Root == nil {
		return
	}

	preOrderTraverse(t.Root, f)
}

// internal recursive function to traverse in pre order
func preOrderTraverse(node *Node, f func(int)) {
	if node == nil {
		return
	}

	f(node.Value)
	preOrderTraverse(node.Left, f)
	preOrderTraverse(node.Right, f)
}
