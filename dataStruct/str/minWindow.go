package str

// 力扣第 76 题
// 最小覆盖子串

// https://leetcode.cn/problems/minimum-window-substring/description/

// https://labuladong.github.io/algo/1/12/#一最小覆盖子串

func MinWindow(s string, t string) string {
	// 滑动窗口左右指针，窗口区间 [left, right)，左闭右开
	left, right := 0, 0
	// 最小覆盖子串的左右索引，end 初始化为一个比较大的值，要大于 len(s)，因为最后窗口结束时的条件是 right == len(s)，此时也无法判断是否问题有解
	start, end := 0, len(s)+1
	needs := make(map[string]int)
	for index := range t {
		c := string(t[index])
		needs[c]++
	}
	windows := make(map[string]int)
	// valid 表示窗口中满足 needs 中的字符个数，当 valid == len(needs) 时，窗口包含子串
	var valid int
	// 寻找满足的解
	for right < len(s) {
		char1 := string(s[right])
		windows[char1]++
		if val1, ok1 := needs[char1]; ok1 {
			if val2 := windows[char1]; val2 == val1 {
				valid++
			}
		}
		right++
		// 寻找最优解
		for valid == len(needs) {
			// 更新最优解
			if right-left < end-start {
				start, end = left, right
			}
			char2 := string(s[left])
			if val1, ok1 := needs[char2]; ok1 {
				if val2 := windows[char2]; val1 == val2 {
					valid--
				}
			}
			windows[char2]--
			left++
		}
	}
	if end != len(s)+1 {
		return s[start:end]
	}
	return ""
}
