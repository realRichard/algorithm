package str

// 力扣第 567 题
// 字符串的排列

// https://leetcode.cn/problems/permutation-in-string/description/

// https://labuladong.github.io/algo/1/12/#二字符串排列

// 相当给你一个 S 和一个 T，请问你 S 中是否存在一个子串，包含 T 中所有字符且不包含其他字符
// 排列 顺序不重要，长度一致
func CheckInclusion(s1 string, s2 string) bool {
	// 滑动窗口左右指针
	left, right := 0, 0
	needs := make(map[string]int)
	for index := range s1 {
		c := string(s1[index])
		needs[c]++
	}
	windows := make(map[string]int)
	// 滑动窗口中满足 needs 的字符个数，当 valid == len(needs) 是，说明全部满足
	var valid int
	// 退出时，结束条件是 right == len(s2)
	for right < len(s2) {
		char1 := string(s2[right])
		windows[char1]++
		if val1, ok1 := needs[char1]; ok1 {
			if val2 := windows[char1]; val1 == val2 {
				valid++
			}
		}
		right++
		// 由于是问 s2 是否包含 s1 的排列，排列是定长的，所以可以维护一个定长窗口
		for right-left >= len(s1) {
			if valid == len(needs) {
				return true
			}
			char2 := string(s2[left])
			if val1, ok1 := needs[char2]; ok1 {
				if val2 := windows[char2]; val2 == val1 {
					valid--
				}
			}
			windows[char2]--
			left++
		}
	}
	return false
}
