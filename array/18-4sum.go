package array

import (
	"sort"
)

/* 题目
给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。
请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：

0 <= a, b, c, d < n
a、b、c 和 d 互不相同
nums[a] + nums[b] + nums[c] + nums[d] == target
你可以按 任意顺序 返回答案
*/

/*实例
示例 1：
输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

示例 2：
输入：nums = [2,2,2,2,2], target = 8
输出：[[2,2,2,2]]
*/

// FourSumTwoPoint 采用三个数字取和的方法
func FourSumTwoPoint(nums []int, target int) [][]int {
	// 特判
	if nums == nil || len(nums) < 4 {
		return [][]int{}
	}
	// 排序
	sort.Ints(nums)
	// 初始化返回值变量
	resSli := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ { // [-2 -1 -1 1 1 2 2]
		// i 判重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			// j 判重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// 初始化两个指针
			left := j + 1
			right := len(nums) - 1
			for right > left {
				curNum := nums[i] + nums[j] + nums[left] + nums[right]
				if curNum > target { // 当前总值偏大，右指针左移
					right--
				} else if curNum < target { // 当前总值偏小，左指针右移
					left++
				} else {
					resSli = append(resSli, []int{nums[i], nums[j], nums[left], nums[right]})
					// 左指针右移，并判重
					left++
					for right > left && nums[left] == nums[left-1] {
						left++
					}
					// 右指针左移，并判重
					right--
					for right > left && nums[right] == nums[right+1] {
						right--
					}
				}
			}
		}
	}

	return resSli
}

// FourSumTwoPointOptimal 双指针 优化版
func FourSumTwoPointOptimal(nums []int, target int) [][]int {
	// 特判
	if nums == nil || len(nums) < 4 {
		return [][]int{}
	}
	// 初始化返回值
	quadruplets := make([][]int, 0)
	// 排序
	sort.Ints(nums)
	// 切片元素总个数
	n := len(nums)
	// 优化点1：nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target
	// 如果前四个总和都大于 target，因为切片时排好序的，后边的元素更大，所以无需往下进行
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		// 优化点2：nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target
		// 如果当前循环的元素，与最后的最大的三个元素的总和还比 target 小，就没有必要进行了。因为排好序的，往前更小
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		// 同优化点1：nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			// 同优化点2 nums[i]+nums[j]+nums[n-2]+nums[n-1] < target
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return quadruplets
}
