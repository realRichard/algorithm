package bfs

import (
	"algorithm/dataStruct/tree"
)

// 力扣第 111 题
// 二叉树的最小深度

// https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/

// https://labuladong.github.io/algo/1/10/#二二叉树的最小高度

func MinDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*tree.TreeNode, 0)
	queue = append(queue, root)
	depth := 1
	for len(queue) != 0 {
		for _, node := range queue {
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		depth++
	}
	return depth
}
