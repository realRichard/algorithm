package tree_test

import (
	"algorithm/dataStruct/tree"
	"testing"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	root := tree.TreeNode{
		Val: 1,
	}
	root.Left = &tree.TreeNode{
		Val: 2,
	}
	// root.Right = &tree.TreeNode{
	// 	Val: 3,
	// }
	// root.Left.Left = &tree.TreeNode{
	// 	Val: 4,
	// }
	// root.Left.Right = &tree.TreeNode{
	// 	Val: 5,
	// }
	diameter := tree.DiameterOfBinaryTree(&root)
	expected := 1
	if diameter != expected {
		t.Errorf("expected %v, result %v", expected, diameter)
	}
}
