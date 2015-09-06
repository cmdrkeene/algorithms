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

func (self *TreeNode) Leaves(nodes ...*TreeNode) *TreeNode {
	for _, node := range nodes {
		self.Adjacent = append(self.Adjacent, node)
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

	for _, node := range root.Adjacent {
		if !node.Visited {
			DepthFirstSearch(node, visit)
		}
	}
}

func BreadthFirstSearch(root *TreeNode, visit visitFunc) {
	if root == nil {
		return
	}

	queue := NewQueue()
	visit(root)
	root.Visited = true
	queue.Enqueue(root)

	for !queue.Empty() {
		next := queue.Dequeue()
		for _, node := range next.(*TreeNode).Adjacent {
			if !node.Visited {
				visit(node)
				node.Visited = true
				queue.Enqueue(node)
			}
		}
	}
}
