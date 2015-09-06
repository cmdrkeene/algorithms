package algorithms

import "fmt"

func NewTree(data interface{}) *TreeNode {
	return &TreeNode{
		Data:    data,
		Visited: false,
	}
}

type TreeNode struct {
	Adjacent []*TreeNode
	Data     interface{}
	Visited  bool
}

func (self *TreeNode) Leaves(leaves ...interface{}) *TreeNode {
	for _, leaf := range leaves {
		switch t := leaf.(type) {
		case *TreeNode:
			self.Adjacent = append(self.Adjacent, t)
		default:
			self.Adjacent = append(self.Adjacent, NewTree(leaf))
		}
	}
	return self
}

func (self *TreeNode) String() string {
	return fmt.Sprintf("%#v", self.Data)
}

type visitFunc func(*TreeNode)

func DepthFirstSearch(root *TreeNode, visit visitFunc) {
	if root == nil {
		return
	}

	visit(root)
	root.Visited = true

	for _, n := range root.Adjacent {
		if !n.Visited {
			DepthFirstSearch(n, visit)
		}
	}
}
