package str

// 力扣第 3 题
// 无重复字符的最长子串

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/

// https://labuladong.github.io/algo/1/12/#四最长无重复子串

func LengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	windows := make(map[string]int)
	length := 0
	for right < len(s) {
		char1 := string(s[right])
		windows[char1]++
		right++
		for windows[char1] > 1 {
			char2 := string(s[left])
			windows[char2]--
			left++
		}
		if right-left > length {
			length = right - left
		}
	}
	return length
}
