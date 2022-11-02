package tree

// 力扣第 543 题
//  二叉树的直径

// https://leetcode.cn/problems/diameter-of-binary-tree/description/

// 记录最大直径的长度
var maxDiameter int

func DiameterOfBinaryTree(root *TreeNode) int {
	// 由于 maxDiameter 是全局变量，每次新调用的时候先初始化为零，否则 leetcode 后面的 case 过不了
	maxDiameter = 0
	maxDepth(root)
	return maxDiameter
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	// 后序位置，顺便计算最大直径
	diameter := left + right
	if diameter > maxDiameter {
		maxDiameter = diameter
	}
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
