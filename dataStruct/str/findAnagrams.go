package str

// 力扣第 438 题
// 找到字符串中所有字母异位词

// https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/

// https://labuladong.github.io/algo/1/12/#三找所有字母异位词

// 异位词，排列，字符的顺序不重要，长度一致
// 输入一个串 S，一个串 T，找到 S 中所有 T 的排列，返回它们的起始索引
func FindAnagrams(s string, p string) []int {
	// 滑动窗口左右指针
	left, right := 0, 0
	// 计数器
	needs := make(map[string]int)
	for index := range p {
		c := string(p[index])
		needs[c]++
	}
	windows := make(map[string]int)
	// 窗口中满足 needs 的字符个数，当 valid == len(needs) 时表示此处满足 needs 中所有字符个数要求
	var valid int
	result := make([]int, 0)
	// 退出条件 right == len(s)
	for right < len(s) {
		char1 := string(s[right])
		windows[char1]++
		if val1, ok1 := needs[char1]; ok1 {
			if val2 := windows[char1]; val2 == val1 {
				valid++
			}
		}
		right++
		// 排列，定长，相当于维护一个定长的窗口
		for right-left >= len(p) {
			if valid == len(needs) {
				result = append(result, left)
			}
			char2 := string(s[left])
			if val1, ok1 := needs[char2]; ok1 {
				if val2 := windows[char2]; val2 == val1 {
					valid--
				}
			}
			windows[char2]--
			left++
		}
	}
	return result
}
