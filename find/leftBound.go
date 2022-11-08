package find

// 寻找左侧边界的二分搜索

// https://labuladong.github.io/algo/1/11/#二寻找左侧边界的二分搜索

func LeftBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	// 当 target 比 nums 中所有元素都大的时候，left 一直往右移动，到退出的时候， left = right + 1(len(nums) - 1 + 1 = len(nums))
	if left == len(nums) {
		return -1
	}
	if nums[left] == target {
		return left
	}
	return -1
}
