package array

import (
	"fmt"
	"sort"
)

/* 题目 31. 下一个排列
整数数组的一个 排列  就是将其所有成员以序列或线性顺序排列。

例如，arr = [1,2,3] ，以下这些都可以视作 arr 的排列：[1,2,3]、[1,3,2]、[3,1,2]、[2,3,1] 。
整数数组的 下一个排列 是指其整数的下一个字典序更大的排列。
更正式地，如果数组的所有排列根据其字典顺序从小到大排列在一个容器中，那么数组的 下一个排列 就是在这个有序容器中排在它后面的那个排列。
如果不存在下一个更大的排列，那么这个数组必须重排为字典序最小的排列（即，其元素按升序排列）。

例如，arr = [1,2,3] 的下一个排列是 [1,3,2] 。
类似地，arr = [2,3,1] 的下一个排列是 [3,1,2] 。
而 arr = [3,2,1] 的下一个排列是 [1,2,3] ，因为 [3,2,1] 不存在一个字典序更大的排列。
给你一个整数数组 nums ，找出 nums 的下一个排列。

必须 原地 修改，只允许使用额外常数空间。
*/

/* 示例
输入：nums = [1,2,3]
输出：[1,3,2]

输入：nums = [3,2,1]
输出：[1,2,3]

输入：nums = [1,1,5]
输出：[1,5,1]

*/

/* 思路与解法
以排列 [4,5,2,6,3,1] 为例：
我们能找到的符合条件的一对「较小数」与「较大数」的组合为 2 与 3，满足「较小数」尽量靠右，而「较大数」尽可能小。
将 2 与 3 交换
当我们完成交换后排列变为 [4,5,3,6,2,1]，此时我们可以重排「较小数」右边的序列，即原来 2 的位置，现在 3 位置后的序列。
最终序列变为 [4,5,3,1,2,6]。
*/

func NextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	// 1、寻找符合条件的一对「较小数」与「较大数」
	minIndex, maxIndex := 0, 0
	maxNum := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if maxNum > nums[i] {
			minIndex = i
			minDiff := maxNum - nums[i]
			for j := i + 1; j < n; j++ {
				if nums[j] > nums[minIndex] && nums[j]-nums[minIndex] <= minDiff {
					minDiff = nums[j] - nums[minIndex]
					maxIndex = j
				}
			}
			break
		} else {
			maxNum = nums[i]
		}
	}
	fmt.Println("min index=", minIndex, "min=", nums[minIndex], "max index=", maxIndex, "max=", nums[maxIndex])

	// 2、交换
	nums[minIndex], nums[maxIndex] = nums[maxIndex], nums[minIndex]
	fmt.Println("交换后的切片为", nums)

	// 3、重排「较小数」右边的序列
	sort.Ints(nums[minIndex+1:])

	// 4、如果没有满足条件的一对「较小数」与「较大数」，直接将切片进行正序排列
	if minIndex == 0 && maxIndex == 0 {
		sort.Ints(nums)
	}

}

// NextPermutationCode 网站源代码
func NextPermutationCode(nums []int) {
	n := len(nums)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}
