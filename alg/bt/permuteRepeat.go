package bt

// 排列（元素无重可复选）

// https://labuladong.github.io/algo/1/9/#排列元素无重可复选

var res7 [][]int
var trace7 []int

func PermuteRepeat(nums []int) [][]int {
	res7 = make([][]int, 0)
	trace7 = make([]int, 0, len(nums))
	permuteRepeat(nums, 0)
	return res7
}

func permuteRepeat(nums []int, row int) {
	if row == len(nums) {
		foo := make([]int, len(trace7))
		copy(foo, trace7)
		res7 = append(res7, foo)
		return
	}
	for i := 0; i < len(nums); i++ {
		trace7 = append(trace7, nums[i])
		permuteRepeat(nums, row+1)
		trace7 = trace7[:len(trace7)-1]
	}
}
