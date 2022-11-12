package dp

import "algorithm/dataStruct/tree"

// 力扣第 337 题
// 打家劫舍 III

// https://leetcode.cn/problems/house-robber-iii/description/

var memory map[*tree.TreeNode]int = make(map[*tree.TreeNode]int)

func Rob3V1(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	// 利用备忘录消除重叠子问题
	if val, ok := memory[root]; ok {
		return val
	}
	// 抢，然后去下下家
	current := 0
	if root.Left != nil {
		current += Rob3V1(root.Left.Left) + Rob3V1(root.Left.Right)
	}
	if root.Right != nil {
		current += Rob3V1(root.Right.Left) + Rob3V1(root.Right.Right)
	}
	current = root.Val + current
	// 不抢，然后去下家
	next := Rob3V1(root.Left) + Rob3V1(root.Right)
	res := 0
	if current > next {
		res = current
	} else {
		res = next
	}
	memory[root] = res
	return res
}

// 定义：输入一个节点，返回一个 result []int
// result[0] 表示不抢当前节点能得到最大的钱
// result[1] 表示抢当前节点能得到最大的钱
func rob3V2(root *tree.TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := rob3V2(root.Left)
	right := rob3V2(root.Right)
	// 抢当前节点就不能抢下一个节点
	rob := root.Val + left[0] + right[0]
	// 不抢当前节点，那么下一个节点可抢也可不抢，取较大者
	notRob := 0
	if left[0] > left[1] {
		notRob += left[0]
	} else {
		notRob += left[1]
	}
	if right[0] > right[1] {
		notRob += right[0]
	} else {
		notRob += right[1]
	}
	return []int{notRob, rob}
}

func Rob3V2(root *tree.TreeNode) int {
	result := rob3V2(root)
	if result[0] > result[1] {
		return result[0]
	} else {
		return result[1]
	}
}
