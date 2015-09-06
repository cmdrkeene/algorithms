package algorithms

import "testing"

func TestBinaryTree_IsBalanced(t *testing.T) {
	// Null Tree
	Expect(t).Equal(
		(&BinaryTreeNode{}).IsBalanced(),
		true,
	)

	// Simple Blanced Tree
	Expect(t).Equal(
		NewBinaryTree(1).Left(2).Right(3).IsBalanced(),
		true,
	)

	// Simple Unbalanced Tree
	Expect(t).Equal(
		NewBinaryTree(1).Left(
			NewBinaryTree(2).Left(
				NewBinaryTree(3).Right(4),
			),
		).Right(5).IsBalanced(),
		false,
	)
}

func TestBinaryTree_IsBalancedOptimized(t *testing.T) {
	// Null Tree
	Expect(t).Equal(
		(&BinaryTreeNode{}).IsBalancedOptimized(),
		true,
	)

	// Simple Blanced Tree
	Expect(t).Equal(
		NewBinaryTree(1).Left(2).Right(3).IsBalancedOptimized(),
		true,
	)

	// Simple Unbalanced Tree
	Expect(t).Equal(
		NewBinaryTree(1).Left(
			NewBinaryTree(2).Left(
				NewBinaryTree(3).Right(4),
			),
		).Right(5).IsBalancedOptimized(),
		false,
	)
}
