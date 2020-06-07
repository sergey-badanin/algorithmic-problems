package binarytreediameter

import "testing"

func TestToStack(t *testing.T) {
	root := testTree01()

	stakced := toStack(root)
	if stakced == nil || stakced.Len() != 7 {
		t.Error("Stack is too small")
		return
	}

	for i := 4; i > 0; i-- {
		elem := stakced.Front()
		node := elem.Value.(*TreeNode)
		if node.Val > 6 || node.Val < 3 {
			t.Errorf("Error in first 4 elements")
			return
		}
	}
}

func testTree01() *TreeNode {
	n0 := &TreeNode{0, nil, nil}
	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{2, nil, nil}

	n0.Left = n1
	n0.Right = n2

	n3 := &TreeNode{3, nil, nil}
	n4 := &TreeNode{4, nil, nil}
	n5 := &TreeNode{5, nil, nil}
	n6 := &TreeNode{6, nil, nil}

	n1.Left = n3
	n1.Right = n4

	n2.Left = n5
	n2.Right = n6

	return n0
}
