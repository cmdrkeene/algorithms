package algorithms

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTree_DepthFirstSearch(t *testing.T) {
	// given
	//         [1]
	//        / | \
	//     [2] [7] [8]
	//     / \     / \
	//   [3] [6] [9] [12]
	//   / \      | \
	// [4] [5]  [10] [11]
	// ====================
	node := func(data interface{}, leaves ...interface{}) *TreeNode {
		return NewTree(data).Leaves(leaves...)
	}
	tree := node(1,
		node(2,
			node(3,
				node(4, 5),
			),
			node(6),
		),
		node(7),
		node(8,
			node(9,
				node(10, 11),
			),
			node(12),
		),
	)

	recorder := newTreeRecorder()

	// when
	DepthFirstSearch(tree, recorder.Visit)

	// then
	Expect(t).Equal(
		recorder.String(),
		"1->2->3->4->5->6->7->8->9->10->11->12",
	)
}

func TestTree_BreadthFirstSearch(t *testing.T) {
}

func newTreeRecorder() *treeRecorder {
	return &treeRecorder{}
}

type treeRecorder struct {
	visited []*TreeNode
}

func (self *treeRecorder) Visit(node *TreeNode) {
	self.visited = append(self.visited, node)
}

func (self *treeRecorder) String() string {
	buf := bytes.NewBufferString("")
	last := self.visited[len(self.visited)-1]

	for _, node := range self.visited {
		fmt.Fprintf(buf, "%s", node)
		if node != last {
			buf.Write([]byte("->"))
		}
	}

	return buf.String()
}
