package binary_search

/*
34. 在排序数组中查找元素的第一个和最后一个位置
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。

示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

示例 3：
输入：nums = [], target = 0
输出：[-1,-1]
*/

// searchRange 此问题可以转换为：
// 查找第一个与 target 相等的元素 searchFirstEqualElement
// 查找最后一个与 target 相等的元素 searchLastEqualElement
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}

	res[0] = searchFirstEqualElement(nums, target)
	if res[0] == -1 {
		return res
	}

	res[1] = searchLastEqualElement(nums, target)

	return res
}

// searchRangeMid 第一次找到target后。以target为基准向左右两边扩散
func searchRangeMid(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}

	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			// 以mid为中心，向右边扩散
			for i := mid; i < len(nums) && nums[i] == target; i++ {
				res[1] = i
			}
			// 以mid为中心，向左边扩散
			for i := mid; i >= 0 && nums[i] == target; i-- {
				res[0] = i
			}
			break
		}

		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return res
}
