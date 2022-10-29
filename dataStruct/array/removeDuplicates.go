package array

// 力扣第 26 题
// 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/

// https://labuladong.github.io/algo/1/5/#一快慢指针技巧

func RemoveDuplicates(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	fast, slow := 0, 0
	for ; fast < length; fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
