package algorithms

import "math"

func NewBinaryTree(data interface{}) *BinaryTreeNode {
	return &BinaryTreeNode{data: data}
}

// BinaryTreeNode implements a binary tree
type BinaryTreeNode struct {
	data  interface{}
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func (self *BinaryTreeNode) Left(node interface{}) *BinaryTreeNode {
	self.left = self.node(node)
	return self
}

func (self *BinaryTreeNode) Right(node interface{}) *BinaryTreeNode {
	self.right = self.node(node)
	return self
}

func (self *BinaryTreeNode) node(value interface{}) *BinaryTreeNode {
	switch n := value.(type) {
	case *BinaryTreeNode:
		return n
	default:
		return NewBinaryTree(node)
	}
}

// IsBalancedOptimized runs in O(N) time and O(TreeHeight) space
func (self *BinaryTreeNode) IsBalancedOptimized() bool {
	if checkHeight(self) == NotBalanced {
		return false
	} else {
		return true
	}
}

// IsBalanced returns if subtree heights do not vary by more than 1
// Runs in O(N log N) time
func (self *BinaryTreeNode) IsBalanced() bool {
	if self == nil {
		return true
	}

	heightDifference := absInt(
		getHeight(self.left) - getHeight(self.right),
	)

	if heightDifference > 1 {
		return false
	} else {
		return self.left.IsBalanced() && self.right.IsBalanced()
	}
}

const NotBalanced = -1

func checkHeight(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := checkHeight(root.left)
	if leftHeight == NotBalanced {
		return NotBalanced
	}

	rightHeight := checkHeight(root.right)
	if rightHeight == NotBalanced {
		return NotBalanced
	}

	diff := absInt(leftHeight - rightHeight)
	if diff > 1 {
		return NotBalanced
	} else {
		return maxInt(leftHeight, rightHeight) + 1
	}
}

func getHeight(root *BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	return maxInt(
		getHeight(root.left),
		getHeight(root.right),
	) + 1
}

func absInt(a int) int {
	return int(math.Abs(float64(a)))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
