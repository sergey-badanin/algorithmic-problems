package binarytreediameter

import "container/list"

/*
 * - Develop heap implementation with array and linked list
 * - Implement diameter search algorithm
 * - Prepare benchmarking suit for different heap implementations
 * - Implements the same in java
 * - Develop performance benchmarking and comparing suit
 */
func diameterOfBinaryTree(root *TreeNode) int {
	return 0
}

func toStack(root *TreeNode) *list.List {
	return nil
}

/*TreeNode - structure
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stackNode struct {
	treeNode  *TreeNode
	height    *int
	lDiameter *int
	rDiameter *int
}
