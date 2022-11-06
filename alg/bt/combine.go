package bt

// 力扣第 77 题
// 组合（元素无重不可复选）

// https://leetcode.cn/problems/combinations/description/

// https://labuladong.github.io/algo/1/9/#组合元素无重不可复选

var res2 [][]int
var trace2 []int

func Combine(n int, k int) [][]int {
	res2 = make([][]int, 0)
	trace2 = make([]int, 0, n)
	combine(n, k, 0)
	return res2
}

func combine(n int, k int, start int) {
	// 类似子集问题，只收集第 k 层节点的值
	if len(trace2) == k {
		res := make([]int, len(trace2))
		copy(res, trace2)
		res2 = append(res2, res)
	}
	for i := start; i < n; i++ {
		trace2 = append(trace2, i+1)
		combine(n, k, i+1)
		trace2 = trace2[:len(trace2)-1]
	}
}
