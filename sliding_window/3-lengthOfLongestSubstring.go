package sliding_window

import "strings"

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

// 滑动窗口（不定长度）
// 初始时，left, right := 0, 0，再定义一个map来存储每个字节的个数
// 每次向右移动right指针，然后将right指针对应byte在map中的个数+1
// 如果right指针对应byte个数大于1，就需要依次将left指针右移，直到right指针对应个数<=1
//
// 完整的思路：
// 用滑动窗口 window 来记录不重复的字符个数， window 为哈希表类型。
// 1. 设定两个指针： left、 right，分别指向滑动窗口的左右边界，保证窗口中没有重复字符。
// 2. 一开始， left、 right 都指向 0。
// 3. 向右移动 right，将最右侧字符 s[right] 加入当前窗口 window 中，记录该字符个数。
// 4. 如果该窗口中该字符的个数多于 1 个，即 window[s[right]]>1，
//    则不断右移 left，缩小滑动窗口长度，并更新窗口中对应字符的个数，直到 window[s[right]]≤1。
// 5.维护更新无重复字符的最长子串长度。然后继续右移 right，直到 right≥len(nums) 结束。
// 6.输出无重复字符的最长子串长度。
func lengthOfLongestSubstring1(s string) int {
	left, right, maxSubLen := 0, 0, 0
	m := make(map[byte]int)

	for right < len(s) {
		// 当前right指针对应的byte
		curRightByte := s[right]
		// 将当前right指针对应byte在map中的个数增加1
		m[curRightByte]++

		// 判断是否需要右移left指针
		for m[curRightByte] > 1 {
			// 将当前left指针对应byte在map中的个数减1
			curLeftByte := s[left]
			m[curLeftByte]--
			// 将left右移一位，进行下次内循环
			left++
		}

		// 判断最大字串个数
		if right-left+1 > maxSubLen {
			maxSubLen = right - left + 1
		}

		// 将right右移一位，进行下次外循环
		right++
	}

	return maxSubLen
}

func lengthOfLongestSubstring2(s string) int {
	m := make(map[byte]int)
	var r = 0
	var n = len(s)
	var ans = 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}

		for r < n && m[s[r]] == 0 {
			m[s[r]]++
			r++
		}

		if r-i > ans {
			ans = r - i
		}
	}
	return ans
}

func lengthOfLongestSubstring3(s string) int {
	// 滑动窗口，只增大不减小
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start
}
