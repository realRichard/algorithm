package dp

// 力扣第 213 题
// 打家劫舍 II

// https://leetcode.cn/problems/house-robber-ii/description/

// 首尾房间不能同时被抢，那么只可能有三种不同情况：
// 	要么都不被抢；
// 	要么第一间房子被抢最后一间不抢；
// 	要么最后一间房子被抢第一间不抢
// 这三种情况，哪种的结果最大，就是最终答案
// 由于房子里的钱是非负，所以选择的余地越多，结果越大，所以直接从情况二三中选
func Rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	state1 := RobV4(nums[1:])
	state2 := RobV4(nums[:len(nums)-1])
	if state1 > state2 {
		return state1
	} else {
		return state2
	}
}
