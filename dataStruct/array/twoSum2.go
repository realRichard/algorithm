package array

// 力扣第 167 题
// 两数之和 II - 输入有序数组

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/

func TwoSum2(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		result := numbers[left] + numbers[right]
		if result == target {
			break
		} else if result < target {
			left++
		} else {
			right--
		}
	}
	return []int{left + 1, right + 1}
}
