package array

import "sort"

// 如果假设输入一个数组 nums 和一个目标和 target，请你返回 nums 中能够凑出 target 的两个元素的值
// nums 中可能有多对儿元素之和都等于 target，请你的算法返回所有和为 target 的元素对儿，其中不能出现重复

// https://mp.weixin.qq.com/s/fSyJVvggxHq28a0SdmZm6Q

func TwoSumTarget(nums []int, target int) [][]int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	res := make([][]int, 0)
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			res = append(res, []int{nums[left], nums[right]})
			left++
			// 去重
			for left < right && nums[left] == nums[left-1] {
				left++
			}
			right--
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		} else if sum < target {
			left++
			// 优化
			for left < right && nums[left] == nums[left-1] {
				left++
			}
		} else if sum > target {
			right--
			for left < right && nums[right] == nums[right+1] {
				right--
			}
		}
	}
	return res
}
