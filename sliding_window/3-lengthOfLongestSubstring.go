package sliding_window

/*
3. 无重复字符的最长子串

给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0

	bytes := make(map[byte]int)
	maxSubLen := 0

	for right < len(s) {
		if num, ok := bytes[s[right]]; ok {
			bytes[s[right]] = num + 1

		} else {
			bytes[s[right]] = 1
		}

		for true {
			rightNum := bytes[s[right]]
			if rightNum <= 1 {
				break
			}
			leftNum := bytes[s[left]]
			bytes[s[left]] = leftNum - 1
			left++
		}

		curLen := right - left + 1
		if curLen > maxSubLen {
			maxSubLen = curLen
		}

		right++
	}

	return maxSubLen
}
