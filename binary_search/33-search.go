package binary_search

/*
33. 搜索旋转排序数组

整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

示例 2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

示例 3：
输入：nums = [1], target = 0
输出：-1

提示：
1 <= nums.length <= 5000
-104 <= nums[i] <= 104
nums 中的每个值都 独一无二
题目数据保证 nums 在预先未知的某个下标上进行了旋转
-104 <= target <= 104
*/

// 与普通的二分查找一样，根据中间元素 mid ，将序列划分为 [low,mid-1]、mid和 [mid+1,high] 。
//
// 如果target == nums[mid] ，代表找到了该元素，直接返回即可；
// 否则，比较target 大于还是小于 nums[mid] ，决定下一个查找的子序列
// 与普通的二分不同的是，在比较target 和 nums[mid] 谁大谁小之前，先比较 nums[low] 和 nums[mid] 的关系，原理如下：
//
// 第一类序列：【2 3 4 5 6 7 1】，也就是 nums[low] < nums[mid] ，此例中为2<5。这种情况下，可以保证序列的前半部分是有序的。
// 如果 nums[low] <= target < nums[mid] ，则去前半部分找
// 否则，去后半部分找
//
// 第二类序列：【6 7 1 2 3 4 5】，也就是 nums[low] > nums[mid] ，此例中为6>2。这种情况下，可以保证序列的后半部分是有序的。
// 如果 nums[mid] <= target < nums[end] ，则去后半部分找
// 否则，去前半部分找

// 根据比较 nums[low] 和 nums[mid] 的大小，得出哪一个子序列是有序的，因为只有有序，才可以进行下一步的大小比较，如判断 nums[low] <= target < nums[mid] 是否成立，
// 如果不能确定具体的序列是否有序，就开始比较的话，很可能发生错误。如序列【3 4 5 6 7 1 2】，就永远找不到一个大于7小于2的元素。
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		middle := low + (high-low)>>1
		if nums[middle] == target {
			return middle
		}
		// 根据 middle 位置元素与 low 位置元素比较，判断 middle 落在在左侧较大的有序序列，还是落在右侧较小的有序序列
		if nums[low] <= nums[middle] { // middle 处于左侧有序序列，即[low, middle]是有序的
			// 判断 target 是在 middle 的左边还是右边，从而调整左右边界 low 和 high
			if nums[low] <= target && target < nums[middle] { // 目标处于 middle 的左侧，且必须大于 low 所在位置元素。即nums[low] <= target < num[middle]
				high = middle - 1
			} else {
				low = middle + 1
			}

		} else { // [middle, high]是有序的
			if nums[middle] < target && target <= nums[high] {
				low = middle + 1
			} else {
				high = middle - 1
			}
		}

	}
	return -1
}
