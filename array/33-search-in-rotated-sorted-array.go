package array

/* 题目
整数数组 nums 按升序排列，数组中的值 互不相同 。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，
使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1。
你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
*/

/*
示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

示例 2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

示例 3：
输入：nums = [1], target = 0
输出：-1
*/

func Search(nums []int, target int) int {
	// 特判
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	resIndex := -1
	if nums[0] > target { // 从后边遍历
		for i := len(nums) - 1; i >= 0 && nums[i] >= target; i-- {
			if i+1 < len(nums) && nums[i] > nums[i+1] { // 从后部分遍历，遍历到了前半部分，直接停止
				break
			}
			if nums[i] == target {
				resIndex = i
				break
			}
		}
	} else { // 从前边遍历
		for i := 0; i < len(nums) && nums[i] <= target; i++ {
			if i > 0 && nums[i] < nums[i-1] { // 从前部分遍历，遍历到了后半部分，直接停止
				break
			}
			if nums[i] == target {
				resIndex = i
				break
			}
		}
	}

	return resIndex
}

// SearchBinary 使用常规的二分查找
// 给出一个数组，数组中本来是从小到大排列的，并且数组中没有重复数字。但是现在把后面随机一段的序列放到数组前面，这样形成了前后两端有序的子序列，前面一段是数值比较大的序列，后面一段是数值偏小的序列
// 先根据 nums[middle] 与 nums[low] 的关系判断 middle 是在左段还是右段，
// 接下来再判断 target 是在 middle 的左边还是右边，从而来调整左右的边界 low 和 high。
func SearchBinary(nums []int, target int) int {
	// 特判
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	// 二分查找
	low, high := 0, len(nums)-1
	for low <= high {
		middle := low + (high-low)>>1
		if nums[middle] == target {
			return middle
		}
		// 根据 middle 位置元素与 low 位置元素比较，判断 middle 落在在左侧较大的有序序列，还是落在右侧较小的有序序列
		if nums[middle] >= nums[low] { // middle 处于左侧有序序列
			// 判断 target 是在 middle 的左边还是右边，从而调整左右边界 low 和 high
			if target < nums[middle] && target >= nums[low] { // 目标处于 middle 的左侧，且必须大于 low 所在位置元素
				high = middle - 1
			} else {
				low = middle + 1
			}

		} else {
			if target > nums[middle] && target <= nums[high] {
				low = middle + 1
			} else {
				high = middle - 1
			}
		}
	}

	return -1
}

/**
方法三：先根据 nums[mid] 与 nums[lo] 的关系判断 mid 是在左段还是右段，接下来再判断 target 是在 mid 的左边还是右边，从而来调整左右边界 lo 和 hi。
Java

public int search(int[] nums, int target) {
    int lo = 0, hi = nums.length - 1, mid = 0;
    while (lo <= hi) {
        mid = lo + (hi - lo) / 2;
        if (nums[mid] == target) {
            return mid;
        }
        // 先根据 nums[mid] 与 nums[lo] 的关系判断 mid 是在左段还是右段
        if (nums[mid] >= nums[lo]) {
            // 再判断 target 是在 mid 的左边还是右边，从而调整左右边界 lo 和 hi
            if (target >= nums[lo] && target < nums[mid]) {
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        } else {
            if (target > nums[mid] && target <= nums[hi]) {
                lo = mid + 1;
            } else {
                hi = mid - 1;
            }
        }
    }
    return -1;
}
*/
