package bt

import "sort"

// 力扣第 40 题
// 组合总和 II

// https://leetcode.cn/problems/combination-sum-ii/description/

// https://labuladong.github.io/algo/1/9/#子集组合元素可重不可复选

var res4 [][]int
var trace4 []int
var checkSum int

func CombinationSum2(candidates []int, target int) [][]int {
	res4 = make([][]int, 0)
	trace4 = make([]int, 0, len(candidates))
	checkSum = 0
	sort.Ints(candidates)
	combinationSum2(candidates, target, 0)
	return res4
}

func combinationSum2(candidates []int, target int, start int) {
	// base case
	// 达到目标和，找到符合条件的组合，也结束，后面的和只会更大
	if checkSum == target {
		foo := make([]int, len(trace4))
		copy(foo, trace4)
		res4 = append(res4, foo)
		return
	}
	// base case
	// 超过目标和，直接结束，不然会超时
	if checkSum > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		trace4 = append(trace4, candidates[i])
		checkSum += candidates[i]
		combinationSum2(candidates, target, i+1)
		trace4 = trace4[:len(trace4)-1]
		checkSum -= candidates[i]
	}
}
