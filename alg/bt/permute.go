package bt

// 力扣第 46 题
// 全排列

// https://leetcode.cn/problems/permutations/description/

// https://labuladong.github.io/algo/1/8/#一全排列问题

// 用于保存最终结果
var res [][]int

func Permute(nums []int) [][]int {
	// 多次调用的时候先清空结果
	res = make([][]int, 0)
	// 已走过路径
	trace := make([]int, 0, len(nums))
	// 「路径」中的元素会被标记为 true，避免重复使用
	used := make([]bool, len(nums))
	backTrace(nums, trace, used)
	return res
}

// 路径：记录在 track 中
// 选择列表：nums 中不存在于 track 的那些元素（used[i] 为 false）
// 结束条件：nums 中的元素全都在 track 中出现
func backTrace(nums []int, trace []int, used []bool) {
	// 触发条件
	if len(trace) == len(nums) {
		t := make([]int, len(nums))
		copy(t, trace)
		res = append(res, t)
	}
	for index, num := range nums {
		// 排除不合法的选择
		if used[index] {
			// nums[i] 已经在 track 中，跳过
			continue
		}
		// 做选择
		trace = append(trace, num)
		used[index] = true
		// 进入下一层决策树
		backTrace(nums, trace, used)
		// 取消选择
		used[index] = false
		trace = trace[:len(trace)-1]
	}
}
