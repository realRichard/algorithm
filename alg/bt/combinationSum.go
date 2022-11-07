package bt

// 力扣第 39 题
// 组合总和

// https://leetcode.cn/problems/combination-sum/description/

// https://labuladong.github.io/algo/1/9/#子集组合元素无重可复选

var res6 [][]int
var trace6 []int
var checkSum6 int

func CombinationSum(candidates []int, target int) [][]int {
	res6 = make([][]int, 0)
	trace6 = make([]int, 0)
	checkSum6 = 0
	combinationSum(candidates, target, 0)
	return res6
}

func combinationSum(candidates []int, target int, start int) {
	// base case，找到目标和，记录结果
	if checkSum6 == target {
		foo := make([]int, len(trace6))
		copy(foo, trace6)
		res6 = append(res6, foo)
		return
	}
	// base case，超过目标和，停止向下遍历
	if checkSum6 > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		trace6 = append(trace6, candidates[i])
		checkSum6 += candidates[i]
		// 同一元素可重复使用，注意参数
		combinationSum(candidates, target, i)
		trace6 = trace6[:len(trace6)-1]
		checkSum6 -= candidates[i]

	}
}
