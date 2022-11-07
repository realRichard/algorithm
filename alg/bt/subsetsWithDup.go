package bt

import "sort"

// 力扣第 90 题
// 子集 II

// https://leetcode.cn/problems/subsets-ii/description/

// https://labuladong.github.io/algo/1/9/#子集组合元素可重不可复选

var res3 [][]int
var trace3 []int

func SubsetsWithDup(nums []int) [][]int {
	res3 = make([][]int, 0)
	trace3 = make([]int, 0, len(nums))
	// 先排序，让相同的元素靠在一起
	sort.Ints(nums)
	subsetsWithDup(nums, 0)
	return res3
}

func subsetsWithDup(nums []int, start int) {
	res := make([]int, len(trace3))
	copy(res, trace3)
	res3 = append(res3, res)
	for i := start; i < len(nums); i++ {
		// 剪枝逻辑，值相同的相邻树枝，只遍历第一条
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		trace3 = append(trace3, nums[i])
		subsetsWithDup(nums, i+1)
		trace3 = trace3[:len(trace3)-1]
	}
}
