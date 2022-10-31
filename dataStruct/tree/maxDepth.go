package tree

// 保存结果
var res int

// 记录当前遍历到的节点深度
var depth int

// 力扣第 104 题
// 二叉树的最大深度

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/

func MaxDepthV1(root *TreeNode) int {
	travesal(root)
	return res
}

// 遍历一遍二叉树，用一个外部变量记录每个节点所在的深度，
// 取最大值就可以得到最大深度，这就是遍历二叉树计算答案的思路
func travesal(root *TreeNode) {
	if root == nil {
		return
	}
	// 进入一个节点的时候
	depth++
	// 到达叶子结点，更新最大长度
	if root.Left == nil && root.Right == nil {
		if depth > res {
			res = depth
		}
	}
	travesal(root.Left)
	travesal(root.Right)
	// 离开一个节点的时候
	depth--
}

// 一棵二叉树的最大深度可以通过子树的最大深度推导出来，这就是分解问题计算答案的思路
// 定义：输入根节点，返回这颗二叉树的最大深度
func MaxDepthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 利用定义，计算出左右子树的最大深度
	left := MaxDepthV2(root.Left)
	right := MaxDepthV2(root.Right)
	// 整棵树的最大深度等于左右子树的最大深度取最大值，
	// 然后再加上根节点自己
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
