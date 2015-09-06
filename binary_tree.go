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

// IsBalanced returns if subtree heights do not vary by more than 1
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

func absInt(a int) int {
	return int(math.Abs(float64(a)))
}

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
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
