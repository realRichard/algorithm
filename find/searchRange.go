package find

// 力扣第 34 题
// 在排序数组中查找元素的第一个和最后一个位置

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/

func SearchRange(nums []int, target int) []int {
	left := LeftBound(nums, target)
	right := RightBound(nums, target)
	return []int{left, right}
}
