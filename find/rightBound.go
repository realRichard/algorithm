package find

// 寻找右侧边界的二分查找

// https://labuladong.github.io/algo/1/11/#三寻找右侧边界的二分查找

func RightBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
	}
	// 当 target 比 nums 中所有元素都小的时候会造成 right 一直往左移动，直到退出时 right = left -1(0 - 1 = -1)
	if right < 0 {
		return -1
	}
	if nums[right] == target {
		return right
	}
	return -1
}
