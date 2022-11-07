package bt

import "sort"

// 力扣第 47 题
// 全排列 II

// https://leetcode.cn/problems/permutations-ii/description/

// https://labuladong.github.io/algo/1/9/#排列元素可重不可复选

var res5 [][]int
var trace5 []int
var used5 []bool

func PermuteUnique(nums []int) [][]int {
	res5 = make([][]int, 0)
	trace5 = make([]int, 0, len(nums))
	used5 = make([]bool, len(nums))
	// 先排序，让相同的元素靠在一起
	sort.Ints(nums)
	permuteUnique(nums, 0)
	return res5
}

func permuteUnique(nums []int, row int) {
	if row == len(nums) {
		foo := make([]int, len(trace5))
		copy(foo, trace5)
		res5 = append(res5, foo)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used5[i] {
			continue
		}
		// 新添加的剪枝逻辑，固定相同的元素在排列中的相对位置
		// 如 nums[i] == nums[i-1]，!used5[i-1] 为 false 表示必须要 nums[i-1] 使用过了才能使用 nums[i]
		// 从保保证相同元素在排列中的相对位置保持不变，就可以达到去重的目的
		if i > 0 && nums[i] == nums[i-1] && !used5[i-1] {
			continue
		}
		trace5 = append(trace5, nums[i])
		used5[i] = true
		permuteUnique(nums, row+1)
		trace5 = trace5[:len(trace5)-1]
		used5[i] = false
	}

}
