package array

import "sort"

// 力扣第 15 题
// 三数之和

// https://leetcode.cn/problems/3sum/description/

func ThreeSum(nums []int) [][]int {
	return ThreeSumTarget(nums, 0)
}

func ThreeSumTarget(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		foo := target - nums[i]
		twoSumResult := TwoSumTarget(nums[i+1:], foo)
		for index, item := range twoSumResult {
			if len(item) != 0 {
				twoSumResult[index] = append(twoSumResult[index], nums[i])
				res = append(res, twoSumResult[index])
			}
		}
		// 去重
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
