package binarytrees

// BST is an implementation of a Binary search tree
//
// BST stores Item instances in an ordered structure, allowing easy insertion,
// removal, and iteration.
//
// Read/Write operations are not safe for concurrent mutation by multiple
// goroutines.
type BST struct {
	root   *BNode
	length int
}

// New build a BST with the root.
func New(value int) *BST {
	return &BST{NewNode(value), 0}
}

// Root returns the root node of the tree.
func (t BST) Root() *BNode {
	return t.root
}

// Insert insert an item in the right position in the tree. Return true if the value was inserted and false otherwise
func (t *BST) Insert(value int) bool {
	node := NewNode(value)
	inserted := true
	if t.root == nil {
		t.root = node
		t.length++
		return inserted
	}

	if inserted = insertNode(t.root, node); inserted {
		t.length++
	}
	return inserted
}

func insertNode(root, newNode *BNode) bool {
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
	inOrderTraverse(t.root, f)
}

func inOrderTraverse(node *BNode, f func(int)) {
	if node == nil {
		return
	}

	inOrderTraverse(node.Left, f)
	f(node.Value)
	inOrderTraverse(node.Right, f)
}

// PreOrderTraverse visits all the nodes in pre order
func (t *BST) PreOrderTraverse(f func(int)) {
	preOrderTraverse(t.root, f)
}

func preOrderTraverse(node *BNode, f func(int)) {
	if node == nil {
		return
	}

	f(node.Value)
	preOrderTraverse(node.Left, f)
	preOrderTraverse(node.Right, f)
}

// PostOrderTraverse visits all the nodes in post order
func (t *BST) PostOrderTraverse(f func(int)) {
	postOrderTraverse(t.root, f)
}

func postOrderTraverse(node *BNode, f func(int)) {
	if node == nil {
		return
	}

	postOrderTraverse(node.Left, f)
	postOrderTraverse(node.Right, f)
	f(node.Value)
}

// Min returns the node with minimal value stored in the tree
func (t *BST) Min() *BNode {
	return minNode(t.root)
}

func minNode(node *BNode) *BNode {
	if node == nil {
		return nil
	}

	current := node
	for {
		if current.Left == nil {
			return current
		}
		current = current.Left
	}
}

// Max returns the node with maximum value stored in the tree
func (t *BST) Max() *BNode {
	return maxNode(t.root)
}

func maxNode(node *BNode) *BNode {
	if node == nil {
		return nil
	}

	current := node
	for {
		if current.Right == nil {
			return current
		}
		current = current.Right
	}
}

// Search returns the node if the value exists in the tree
func (t *BST) Search(value int) *BNode {
	return searchNode(t.root, value)
}

func searchNode(node *BNode, value int) *BNode {
	if node == nil {
		return nil
	}

	if node.Value == value {
		return node
	}

	if node.Value > value {
		return searchNode(node.Left, value)
	} else {
		return searchNode(node.Right, value)
	}
}

// Has returns true if if the value exists in the tree
func (t *BST) Has(value int) bool {
	return t.Search(value) != nil
}

// Remove remove an item from the tree. Return true if the value was removed and false otherwise.
func (t *BST) Remove(value int) bool {
	var removed bool
	t.root, removed = removeNode(t.root, value)
	if removed {
		t.length--
	}
	return removed
}

func removeNode(node *BNode, value int) (*BNode, bool) {
	var removed bool
	if node == nil {
		return nil, removed
	}

	// recursive flows to find the item. When it's found this is avoided
	if node.Value > value {
		node.Left, removed = removeNode(node.Left, value)
		return node, removed
	} else if node.Value < value {
		node.Right, removed = removeNode(node.Right, value)
		return node, removed
	}

	// after this point the node.Value == value and the node will be deleted.
	// for the 3 delete cases early return to no replace with the next one
	removed = true
	// delete case 1: delete leaf node. Remove the node
	if node.Left == nil && node.Right == nil {
		return nil, removed
	}

	// delete case 2: delete half-leaf node
	if node.Left == nil {
		node = node.Right
		return node, removed
	} else if node.Right == nil {
		node = node.Left
		return node, removed
	}

	// delete case 3: delete an inner node
	replacement := maxNode(node.Left)
	node.Value = replacement.Value
	node.Left, removed = removeNode(node.Left, node.Value)
	return node, removed
}

// Len returns the number of items currently in the tree.
func (t *BST) Len() int {
	return t.length
}

// Height return the height of a tree
func (t *BST) Height() int {
	return nodeHeight(t.root)
}

func nodeHeight(node *BNode) int {
	if node == nil {
		return 0
	}

	return intMax(nodeHeight(node.Left), nodeHeight(node.Right)) + 1
}