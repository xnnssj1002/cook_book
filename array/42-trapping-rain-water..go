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

// TrapDoublePointForceOffice 官方解法 按列求水的高度，最后累加
// 本题是想求针对每个 i，找到它左边最大值 leftMax，右边的最大值 rightMax，然后 min(leftMax，rightMax) 为能够接到水的高度。left 和 right 指针是两边往中间移动的游标指针。
// 最傻地解题思路是针对每个下标 i，往左循环找到第一个最大值，往右循环找到第一个最大值，然后把这两个最大值取出最小者，即为当前雨水的高度。这样做时间复杂度高，浪费了很多循环。
// i 在从左往右的过程中，是可以动态维护最大值的。右边的最大值用右边的游标指针来维护。从左往右扫一遍下标，和，从两边往中间遍历一遍下标，是相同的结果，每个下标都遍历了一次。
// 每个 i 的宽度固定为 1，所以每个“坑”只需要求出高度，即当前这个“坑”能积攒的雨水。最后依次将每个“坑”中的雨水相加既是能接到的雨水数。
//
// 总结：从两边的两个指针 left、right 依次往中间碰撞，在碰撞过程中依次维护左右方向最大的高度maxLeft、maxRight。然后取出两个指针对应元素最小的指针，求当前指针位置竖向所能存水的多少
// 如果 left 所在元素最小，则需要将left进行++, 在++之前需要判断 height[left] 与 maxLeft 的大小关系。【left与right哪个不动，哪个就是已经循环的元素中最大的】
// 		1、若 height[left] < maxLeft，需要求left位置水在竖向的面积，即maxLeft - height[left]，因为宽度为1。
// 		2、若 height[left] >= maxLeft,
func TrapDoublePointForceOffice(height []int) int {
	res, left, right, maxLeft, maxRight := 0, 0, len(height)-1, 0, 0
	for left <= right { // 碰撞到且之前，都进行走向彼此
		if height[left] < height[right] {
			if height[left] < maxLeft {
				res += maxLeft - height[left]
			} else {
				maxLeft = height[left]
			}
			left++
		} else {
			if height[right] < maxRight {
				res += maxRight - height[right]
			} else {
				maxRight = height[right]
			}
			right--
		}

	}
	return res
}
