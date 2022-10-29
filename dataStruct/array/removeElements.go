package array

// 力扣第 27 题
// 移除元素

// https://leetcode.cn/problems/remove-element/description/

func RemoveElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return length
	}
	fast, slow := 0, 0
	for ; fast < length; fast++ {
		// 将不等于 val 的元素往前搬
		if nums[fast] != val {
			nums[slow] = nums[fast]
			// 下一次待搬迁的位置
			slow++
		}
	}
	// 搬迁数组的长度
	return slow
}
