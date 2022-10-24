package array

/*
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
*/

/*
示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

示例2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]

示例 3：
输入：nums = [], target = 0
输出：[-1,-1]
*/

func SearchRange(nums []int, target int) []int {
	returnSli := []int{-1, -1}
	if len(nums) == 0 || (len(nums) == 1 && nums[0] != target) {
		return returnSli
	}

	low, high := 0, len(nums)-1
	for low <= high {
		middle := low + (high-low)>>1
		if target == nums[middle] {
			for i := middle; i < len(nums) && nums[i] == target; i++ { // 往后查找
				returnSli[1] = i
			}
			for i := middle; i >= 0 && nums[i] == target; i-- { // 往前查找
				returnSli[0] = i
			}
			break
		} else if target < nums[middle] {
			high = middle - 1
		} else {
			low = middle + 1
		}

	}

	return returnSli
}
