package find

// 力扣第 704 题
// 二分查找

// https://leetcode.cn/problems/binary-search/description/

// https://labuladong.github.io/algo/1/11/#一寻找一个数基本的二分搜索

func Search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		// 防止可能 left right 太大，相加直接溢出
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left++
		} else if nums[mid] > target {
			right--
		}
	}
	return -1
}
