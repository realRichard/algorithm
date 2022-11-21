package array

// 力扣第 303 题
// 区域和检索 - 数组不可变

// https://leetcode.cn/problems/range-sum-query-immutable/description/

// https://labuladong.github.io/algo/2/20/24/#一维数组中的前缀和

// 前缀和
type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	numArray := NumArray{}
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	numArray.preSum = preSum
	return numArray
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.preSum[right+1] - this.preSum[left]
}
