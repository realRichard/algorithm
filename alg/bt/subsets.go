package bt

// 力扣第 78 题
// 子集（元素无重不可复选）

// https://leetcode.cn/problems/subsets/

// https://labuladong.github.io/algo/1/9/#子集元素无重不可复选

// 保存结果集
var res1 [][]int

// 保存路径
var trace []int

func Subsets(nums []int) [][]int {
	res1 = make([][]int, 0)
	trace = make([]int, 0, len(nums))
	subsets(nums, 0)
	return res1
}

func subsets(nums []int, start int) {
	res := make([]int, len(trace))
	copy(res, trace)
	res1 = append(res1, res)
	for i := start; i < len(nums); i++ {
		// 做选择
		trace = append(trace, nums[i])
		// 通过 start 参数控制树枝的遍历，避免产生重复的子集
		subsets(nums, i+1)
		// 撤销选择
		trace = trace[:len(trace)-1]
	}
}
