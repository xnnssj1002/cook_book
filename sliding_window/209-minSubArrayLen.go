package sliding_window

/*
209. 长度最小的子数组
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其总和大于等于 target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。
如果不存在符合条件的子数组，返回 0 。

示例 1：
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。

示例 2：
输入：target = 4, nums = [1,4,4]
输出：1

示例 3：
输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0
*/

// 不定长滑动窗口
// 1、left, right := 0, 0
// 2、先循环right，使right一直往右移动，直到窗口中的元素总和>=target。判断窗口元素个数是否比以前的窗口数量少
// 3、再循环left，使left一直往右移动，直到窗口中的元素<target
// 4、重复2、3，直到right达到len(nums)-1
func minSubArrayLen1(target int, nums []int) int {
	left, right, max, ans := 0, 0, 0, 0

	for right < len(nums) {
		max += nums[right]
		if max >= target {

			for max >= target {
				if ans == 0 || right-left+1 < ans {
					ans = right - left + 1
				}

				max -= nums[left]
				left++
			}

			right++
		} else {
			right++
		}

	}

	return ans
}

// 对方法 minSubArrayLen1 代码结构的优化1
func minSubArrayLenStruct1Optimize1(target int, nums []int) int {
	left, right, max, ans := 0, 0, 0, 0

	for right < len(nums) {
		max += nums[right]

		for max >= target {
			if ans == 0 || right-left+1 < ans {
				ans = right - left + 1
				// 这里再加一个优化点。若一个数字都大于target，则直接返回即可
				if ans == 1 {
					return 1
				}
			}

			max -= nums[left]
			left++
		}

		right++

	}

	return ans
}

// 对方法 minSubArrayLen1 代码结构的优化2
// 通过for代码块自动控制步幅的增长
func minSubArrayLenStruct1Optimize2(target int, nums []int) int {
	left, max, ans := 0, 0, 0

	for right, num := range nums {
		max += num

		for ; max >= target; left++ {
			if ans == 0 || right-left+1 < ans {
				ans = right - left + 1
				// 这里再加一个优化点。若一个数字都大于target，则直接返回即可
				if ans == 1 {
					return 1
				}
			}
			max -= nums[left]
		}

	}

	return ans
}
