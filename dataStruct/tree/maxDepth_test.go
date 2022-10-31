package tree_test

import (
	"algorithm/dataStruct/tree"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	root := tree.TreeNode{
		Val: 3,
	}
	root.Left = &tree.TreeNode{
		Val: 9,
	}
	root.Right = &tree.TreeNode{
		Val: 20,
	}

	root.Right.Left = &tree.TreeNode{
		Val: 15,
	}
	root.Right.Right = &tree.TreeNode{
		Val: 7,
	}
	// root := tree.TreeNode{
	// 	Val: 1,
	// }
	// root.Right = &tree.TreeNode{
	// 	Val: 2,
	// }

	res := tree.MaxDepthV2(&root)
	expected := 3
	if res != expected {
		t.Errorf("expected %v, result %v", expected, res)
	}
}
