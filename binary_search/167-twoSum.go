package binary_search

/*
167. 两数之和 II - 输入有序数组 ~ 二分查找

给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。
如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。
以长度为2的整数数组 [index1, index2] 的形式返回这两个整数的下标 index1 和 index2。

你可以假设每个输入 只对应唯一的答案 ，而且你 不可以 重复使用相同的元素。
你所设计的解决方案必须只使用常量级的额外空间。

示例 1：
输入：numbers = [2,7,11,15], target = 9
输出：[1,2]
解释：2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。

示例 2：
输入：numbers = [2,3,4], target = 6
输出：[1,3]
解释：2 与 4 之和等于目标数 6 。因此 index1 = 1, index2 = 3 。返回 [1, 3] 。

示例 3：
输入：numbers = [-1,0], target = -1
输出：[1,2]
解释：-1 与 0 之和等于目标数 -1 。因此 index1 = 1, index2 = 2 。返回 [1, 2] 。

提示：
仅存在一个有效答案
*/

// twoSumSort 二分查找解法
// 可以首先固定第一个数，然后寻找第二个数，第二个数等于目标值减去第一个数的差。
// 利用数组的有序性质，可以通过二分查找的方法寻找第二个数。
// 为了避免重复寻找，在寻找第二个数时，只在第一个数的右侧寻找
func twoSumSort(numbers []int, target int) []int {

	for i := 0; i < len(numbers); i++ {
		// 使用二分查找，需要查找新目标值
		newTarget := target - numbers[i]

		// 二分查找的逻辑
		low, high := i+1, len(numbers)-1
		for low <= high {
			mid := low + (high-low)>>1
			if numbers[mid] == newTarget {
				return []int{i + 1, mid + 1}
			}
			if newTarget < numbers[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}

	}

	return []int{-1, -1}
}
