package array

// 力扣第 283 题
// 移动零

// https://leetcode.cn/problems/move-zeroes/description/

func MoveZores(nums []int) {
	length := RemoveElement(nums, 0)
	for i := length; i < len(nums); i++ {
		nums[i] = 0
	}
}
