package dp_test

import (
	"algorithm/alg/dp"
	"algorithm/dataStruct/tree"
	"testing"
)

func TestRob3(t *testing.T) {
	root := tree.TreeNode{
		Val: 3,
	}
	root.Left = &tree.TreeNode{
		Val: 2,
	}
	root.Right = &tree.TreeNode{
		Val: 3,
	}
	root.Left.Right = &tree.TreeNode{
		Val: 3,
	}
	root.Right.Right = &tree.TreeNode{
		Val: 1,
	}
	expected := 7
	result := dp.Rob3V2(&root)
	if result != expected {
		t.Errorf("expected %v, result %v\n", expected, result)
	}
}
