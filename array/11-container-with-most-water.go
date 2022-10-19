package array

import "math"

/* 题目
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i])。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器。
*/

/* 实例
输入：[1,8,6,2,5,4,8,3,7]
输出：49

输入：height = [1,1]
输出：1
*/

// MaxAreaLoop 简单粗暴的算法，挨个获取面积，并进行比较，返回最大值
// 在leetcode上提交会报运行超时。leetcode测试用例太变态了，一下给了好几个页面的参数，@~@
func MaxAreaLoop(height []int) int {
	maxArea := 0
	if len(height) == 0 {
		return maxArea
	}

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			curMaxArea := int(math.Min(float64(height[i]), float64(height[j]))) * (j - i)
			if curMaxArea > maxArea {
				maxArea = curMaxArea
			}
		}
	}
	return maxArea

}

// MaxAreaCode 使用力扣上边的题解
// 这一题是对撞指针的思路。首尾分别 2 个指针，每次向内移动最短板以后都分别判断长宽的乘积是否最大。
// 在每个状态下，无论长板或短板哪个向中间收窄一格，都会导致水槽 底边宽度 -1 变短：
//
//	若向内 移动短板 ，水槽的短板 min(h[i], h[j])min(h[i],h[j]) 可能变大，因此下个水槽的面积 可能增大 。
//	若向内 移动长板 ，水槽的短板 min(h[i], h[j])min(h[i],h[j])​ 不变或变小，因此下个水槽的面积 一定变小 。
//	因此，初始化双指针分列水槽左右两端，循环每轮将短板向内移动一格，并更新面积最大值，直到两指针相遇时跳出；即可获得最大面积。
//
//算法流程：
//	初始化： 双指针 i , j 分列水槽左右两端；
//	循环收窄： 直至双指针相遇时跳出；
//	更新面积最大值 maxArea ；
//	选定两板高度中的短板，向中间收窄一格；
//	返回值： 返回面积最大值 maxArea 即可；
func MaxAreaCode(height []int) int {
	maxArea := 0
	if len(height) == 0 {
		return maxArea
	}
	startIndex, endIndex := 0, len(height)-1
	for {
		if startIndex == endIndex {
			return maxArea
		}

		curMaxArea := 0
		if height[startIndex] < height[endIndex] {
			curMaxArea = height[startIndex] * (endIndex - startIndex)
			startIndex++
		} else {
			curMaxArea = height[endIndex] * (endIndex - startIndex)
			endIndex--
		}

		if curMaxArea > maxArea {
			maxArea = curMaxArea
		}

	}

}
