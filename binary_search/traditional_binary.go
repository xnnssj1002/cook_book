package binary_search

// ************** 传统的二分写法 ~ 直接法 **************
// 二分搜索的经典写法。需要注意的三点：
// 	1、循环退出条件，注意是 low <= high，而不是 low < high。
// 	2、mid 的取值，mid := low + (high-low)»1
// 	3、low 和 high 的更新。low = mid + 1，high = mid - 1。

// 从一个数组中查找元素是否存在。如果存在返回该元素对应的下标；如果不存在返回-1
func binarySearchMatrix(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high { // 1、循环退出条件，注意是 low <= high，而不是 low < high。
		middle := low + (high-low)>>1 // 2、middle 的取值，middle := low + (high-low)>>1
		if nums[middle] == target {
			return middle

			// 3、low 和 high 的更新，low = mid + 1，high = mid - 1
		} else if nums[middle] > target {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}

	return -1
}

// ************** 传统的二分写法 ~ 排除法1 目标值在中间值的右侧 **************
func binarySearchMatrix1(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high { // 循环退出条件是low=high
		mid := low + (high-low)>>1
		if target > nums[mid] {
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

// ************** 传统的二分写法 ~ 排除法2 目标值在中间值的左侧 **************
// 这里取中间值为mid := low + (high-low+1)>>1的原因如下：
// 为了避免陷入死循环，当区分被划分为 [low, mid - 1] 与 [mid, high]两部分时，mid 取值要向上取整。即 mid = low + (high - low + 1) / 2。
// 因为如果当区间中只剩下两个元素时（此时 high = low + 1），一旦进入 low = mid 分支，区间就不会再缩小了，下一次循环的查找区间还是 [low, high]，就陷入了死循环。
// 比如左边界 low = 5，右边界 high = 6，此时查找区间为 [5, 6]，mid = 5 + (6 - 5) / 2 = 5，如果进入 low = mid 分支，那么下次查找区间仍为 [5, 6]，区间不再缩小，陷入死循环。
// 这种情况下，mid 应该向上取整，mid = 5 + (6 - 5 + 1) / 2 = 6，如果进入 low = mid 分支，则下次查找区间为 [6, 6]。
//
// 关于排除法的边界设置，可以记忆为：可以记为：
// low = mid + 1、high = mid 和 mid = low + (high - low) / 2 一定是配对出现的。即 high=mid时，求中间值需要取左边的数(数组个数为偶数时)
// high = mid - 1、low = mid 和 mid = low + (high - low + 1) / 2 一定是配对出现的。即low=mid时，求中间值时需要取右边的数(数组个数为偶数时)
func binarySearchMatrix2(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low < high { // 循环退出条件是low-high
		mid := low + (high-low+1)>>1 // 注意这里已经是向上取整，与排除法1的不同之处
		if target < nums[mid] {
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

// ************** 二分搜索的变种写法 **************
// 	1、查找第一个与 target 相等的元素，时间复杂度 O(logn)
//  2、查找最后一个与 target 相等的元素，时间复杂度 O(logn)
// 	3、查找第一个大于等于 target 的元素，时间复杂度 O(logn)
// 	4、查找最后一个小于等于 target 的元素，时间复杂度 O(logn)

// 查找第一个与 target 相等的元素，时间复杂度 O(logn)
func searchFirstEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		middle := low + (high-low)>>1
		if nums[middle] == target {
			// 找到第一个与 target 相等的元素
			if middle == 0 || nums[middle] != nums[middle-1] {
				return middle
			}
			high = middle - 1
		} else if nums[middle] > target {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -1
}

// 查找最后一个与 target 相等的元素，时间复杂度 O(logn)
func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		middle := low + (high-low)>>1
		if nums[middle] == target {
			// 找到最后一个与 target 相等的元素
			if middle == len(nums)-1 || nums[middle] != nums[middle+1] {
				return middle
			}
			low = middle + 1
		} else if nums[middle] > target {
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -1
}

// 查找第一个大于等于 target 的元素，时间复杂度 O(logn)
func searchFirstGreaterElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		middle := low + (high-low)>>1
		if nums[middle] >= target {
			// 找到第一个大于等于 target 的元素
			if middle == 0 || nums[middle-1] < target {
				return middle
			}
			high = middle - 1
		} else {
			low = middle + 1
		}
	}
	return -1
}

// 查找最后一个小于等于 target 的元素，时间复杂度 O(logn)
func searchLastLessElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		middle := low + (high-low)>>1
		if nums[middle] <= target {
			if middle == len(nums)-1 || nums[middle+1] > target {
				return middle
			}
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return -1
}
