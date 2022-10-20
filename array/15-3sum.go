package array

import (
	"fmt"
	"sort"
)

/* 题目
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。
请你返回所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
*/

/* 实例
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
	nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
	nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
	nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
	不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
	注意，输出的顺序和三元组的顺序并不重要。

输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。

输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。
*/

// ThreeSum 使用三个指针依次往后移动。难点在于去重。该方法无法去重
func ThreeSum(nums []int) [][]int {
	resSli := make([][]int, 0)
	if nums == nil || len(nums) < 3 {
		return resSli
	}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					resSli = append(resSli, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return resSli
}

// ThreeSumCode 使用 code 中的题解
// 排序 + 双指针
//本题的难点在于如何去除重复解。
//
//算法流程：
//1、特判，对于数组长度 n，如果切片为 nil 或者切片长度小于 3，返回 []。
//2、对切片进行排序。
//3、遍历排序后切片：
//	若 nums[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果。
//	对于重复元素：跳过，避免出现重复解
//	令左指针 L=i+1，右指针 R=n-1，当 L<R 时，执行循环：
//		当 nums[i]+nums[L]+nums[R]==0时，判断左界和右界是否和下一位置重复，去除重复解，添加到新切片中。并同时将 L,R 移到下一位置，寻找新的解
//		若和大于 0，说明 nums[R] 太大，R 左移
//		若和小于 0，说明 nums[L] 太小，L 右移
func ThreeSumCode(nums []int) [][]int {
	resSli := make([][]int, 0)
	// 特判
	if nums == nil || len(nums) < 3 {
		return resSli
	}
	// 排序
	fmt.Println("原切片", nums)
	sort.Sort(mySlice(nums))
	fmt.Println("排序后的切片", nums)

	// 遍历排序后切片
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		// 索引i对应元素去重。循环到的i如果不是第一个，与前一个判断，如果与前一个相等，直接执行下一次循环
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 定义左右指针
		L := i + 1
		R := len(nums) - 1
		for L < R {
			curNums := nums[i] + nums[L] + nums[R]
			if curNums > 0 { // 需要右指针左移，来使总值减小
				R--
			} else if curNums < 0 { // 需要左指针右移，来使总值增大
				L++
			} else {
				resSli = append(resSli, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] { // 左指针判重
					L++
				}
				for L < R && nums[R] == nums[R-1] { // 右指针判重
					R--
				}
				L++
				R--
			}
		}
	}

	return resSli
}

type mySlice []int

func (s mySlice) Len() int           { return len(s) }
func (s mySlice) Less(i, j int) bool { return s[i] < s[j] }
func (s mySlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
