package two_points

/*
11. 盛最多水的容器
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。

示例 1：
输入：[1,8,6,2,5,4,8,3,7]
输出：49

示例 2：
输入：height = [1,1]
输出：1

*/

func maxArea(height []int) int {
	left, right, max := 0, len(height)-1, 0

	for left < right {
		curMax := 0

		if height[left] < height[right] {
			curMax = (right - left) * height[left]
			left++

		} else {
			curMax = (right - left) * height[right]
			right--
		}

		if curMax > max {
			max = curMax
		}
	}

	return max

}