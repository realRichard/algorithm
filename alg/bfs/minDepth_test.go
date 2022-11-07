package bfs_test

import (
	"algorithm/alg/bfs"
	"algorithm/dataStruct/tree"
	"testing"
)

func TestMinDepth(t *testing.T) {
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
	minDepth := bfs.MinDepth(&root)
	expected := 2
	if minDepth != expected {
		t.Errorf("expected %v, result %v\n", expected, minDepth)
	}
}
