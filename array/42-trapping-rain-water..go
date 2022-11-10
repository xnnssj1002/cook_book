package array

/* 42. 接雨水 https://leetcode.cn/problems/trapping-rain-water/
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
*/

/*
示例1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9
*/

// TrapDoublePointForce 双指针暴力解法
// 内层循环遇到比外层元素大时，内层停止
// 在内层停止之前，累加上 (height[j]-middleItem 高度差) * j到i的距离
func TrapDoublePointForce(height []int) int {
	res := 0
	for i := 0; i < len(height)-1; i++ {
		if height[i] == 0 {
			continue
		}
		middleItem := 0
		for j := i + 1; j < len(height); j++ {
			if height[j] >= height[i] {
				res = res + ((j - i - 1) * (height[i] - middleItem))
				break
			} else {
				if height[j] > middleItem {
					res = res + ((j - i - 1) * (height[j] - middleItem))
					middleItem = height[j]
				}
			}
		}
	}
	return res
}
