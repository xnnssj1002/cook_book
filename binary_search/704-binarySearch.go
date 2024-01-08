package binary_search

// 704. 二分查找

/*
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target.
写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
*/

/*
示例 1:
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4

示例 2:
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
*/

// binarySearch 直接法
func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// binarySearchRemoveLeft 排除法，排除掉左侧部分
func binarySearchRemoveLeft(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high { // 循环退出条件是low=high
		mid := low + (high-low)>>1
		if nums[mid] < target { // 排除mid左侧的部分
			low = mid + 1
		} else {
			high = mid
		}
	}

	if nums[low] == target {
		return low
	} else {
		return -1
	}
}

// binarySearchRemoveRight 排除法，排除掉右侧部分
// 这里取中间值为mid := low + (high-low+1)>>1的原因如下：
// 为了避免陷入死循环，当区分被划分为 [low, mid - 1] 与 [mid, high]两部分时，mid 取值要向上取整。即 mid = low + (high - low + 1) / 2。
// 因为如果当区间中只剩下两个元素时（此时 high = low + 1），一旦进入 low = mid 分支，区间就不会再缩小了，下一次循环的查找区间还是 [low, high]，就陷入了死循环。
// 比如左边界 low = 5，右边界 high = 6，此时查找区间为 [5, 6]，mid = 5 + (6 - 5) / 2 = 5，如果进入 low = mid 分支，那么下次查找区间仍为 [5, 6]，区间不再缩小，陷入死循环。
// 这种情况下，mid 应该向上取整，mid = 5 + (6 - 5 + 1) / 2 = 6，如果进入 low = mid 分支，则下次查找区间为 [6, 6]。
//
// 关于排除法的边界设置，可以记忆为：可以记为：
// left = mid + 1、right = mid 和 mid = left + (right - left) / 2 一定是配对出现的。即目标值在中间值右侧(大的一侧)，中间值向下取整
// right = mid - 1、left = mid 和 mid = left + (right - left + 1) / 2 一定是配对出现的。即目标值在中间值左侧(小的一侧)，中间值向上取整
func binarySearchRemoveRight(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high { // 循环退出条件是low=high
		mid := low + (high-low+1)>>1
		if target < nums[mid] { // 排除mid左侧的部分
			high = mid - 1
		} else {
			low = mid
		}
	}
	if nums[low] == target {
		return low
	} else {
		return -1
	}
}

// 直接法与排除法的适用范围
// 直接法：因为判断语句是 left <= right，有时候要考虑返回是 left 还是 right。循环体内有 3 个分支，并且一定有一个分支用于退出循环或者直接返回。
// 这种思路适合解决简单题目。即要查找的元素性质简单，数组中都是非重复元素，且 ==、>、< 的情况非常好写的时候。
//
// 排除法：更加符合二分查找算法的减治思想。每次排除目标元素一定不存在的区间，达到减少问题规模的效果。然后在可能存在的区间内继续查找目标元素。
// 这种思路适合解决复杂题目。比如查找一个数组里可能不存在的元素，找边界问题，可以使用这种思路。
