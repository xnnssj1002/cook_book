package array

import (
	"sort"
)

/* 题目：41. 缺失的第一个正数
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
*/

/* 示例
示例 1：
输入：nums = [1,2,0]
输出：3

示例 2：
输入：nums = [3,4,-1,1]
输出：2

示例 3：
输入：nums = [7,8,9,11,12]
输出：1
*/

// FirstMissingPositiveSort 先将数组排序，在逐个遍历。对不同的场景进行判断处理
func FirstMissingPositiveSort(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 排好序的切片
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		// 第一个元素大于1时，直接返回1
		if nums[0] > 1 {
			return 1
		}

		// 当前元素大于0，并且比上一个元素的差值大于1
		if i > 0 && nums[i] > 0 && nums[i]-nums[i-1] > 1 {
			// 当前元素大于1， 上一个元素小于0，直接返回1
			if nums[i] > 1 && nums[i-1] < 0 {
				return 1
			}
			// 当前元素以及上一个元素都大于0，且之间差值大于1
			if nums[i-1] >= 0 {
				return nums[i-1] + 1
			}
		}
	}

	if nums[len(nums)-1] < 0 {
		return 1
	}
	return nums[len(nums)-1] + 1
}

// FirstMissingPositiveSameHash 官方题解哈希表，其实是模拟哈希表的方法，即打标记的方法
// 实际上，对于一个长度为 N 的数组，其中没有出现的最小正整数只能在 [1,N+1] 中。这是因为如果 [1,N] 都出现了，那么答案是 N+1，否则答案是 [1,N] 中没有出现的最小正整数。
// 这样一来，我们将所有在 [1,N] 范围内的数放入哈希表，也可以得到最终的答案。而给定的数组恰好长度为 N，这让我们有了一种将数组设计成哈希表的思路：
//
// 我们对数组进行遍历，对于遍历到的数 x，如果它在 [1,N] 的范围内，那么就将数组中的第 x−1 个位置（注意：数组下标从 0 开始）打上「标记」。
// 在遍历结束之后，如果所有的位置都被打上了标记，那么答案是 N+1，否则答案是最小的没有打上标记的位置加 1。
//
// 那么如何设计这个「标记」呢？由于数组中的数没有任何限制，因此这并不是一件容易的事情。
// 但我们可以继续利用上面的提到的性质：由于我们只在意 [1,N] 中的数，因此我们可以先对数组进行遍历，把不在 [1,N] 范围内的数修改成任意一个大于 N 的数（例如 N+1）。
// 这样一来，数组中的所有数就都是正数了，因此我们就可以将「标记」表示为「负号」。
func FirstMissingPositiveSameHash(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 1、我们将数组中所有小于等于 0 的数修改为 N+1；
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	var abs func(int) int
	abs = func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}

	// 2、遍历数组中的每一个数 x，它可能已经被打了标记，因此原本对应的数为 |x|，其中 || 为绝对值符号。如果 ∣x∣∈[1,N]，那么我们给数组中的第 ∣x∣−1 个位置的数添加一个负号。注意如果它已经有负号，不需要重复添加；
	for i := 0; i < len(nums); i++ {
		//if abs(nums[i]) >= 1 && abs(nums[i]) <= n {
		//	if nums[abs(nums[i])-1] > 0 {
		//		nums[abs(nums[i])-1] = -nums[abs(nums[i])-1]
		//	}
		//}

		num := abs(nums[i])
		if num <= n {
			nums[num-1] = -abs(nums[num-1])
		}
	}

	// 3、再次遍历整个数组，如果数组中的每一个数都是负数，那么答案是 N+1，否则答案是第一个正数的位置加 1。
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1
}

// FirstMissingPositiveExchange 官方置换法
// 总：N为切片长度。如果数组中包含 x∈[1,N]，那么恢复后，数组的第 x−1 个元素为 x。
// 在恢复后，数组应当有 [1, 2, ..., N] 的形式，但其中有若干个位置上的数是错误的，每一个错误的位置就代表了一个缺失的正数。
// 以示例二 [3, 4, -1, 1] 为例，恢复后的数组应当为 [1, -1, 3, 4]，我们就可以知道缺失的数为 2。
//
// 那么如何将数组进行恢复呢？
// 可以对数组进行一次遍历，对于遍历到的数 x=nums[i]，如果 x∈[1,N]，我们就知道 x 应当出现在数组中的 x−1 的位置，因此交换 nums[i] 和 nums[x−1]，这样 x 就出现在了正确的位置。
// 在完成交换后，现在 nums[i]的值为原来nums[x-1]，即此次循环下标 i 的值发生了变化，新的 nums[i] 可能还在 [1,N] 的范围内，因此我们需要继续进行交换操作，直到 x !∈ [1,N]。
//
// 注意到上面的方法可能会陷入死循环。如果 nums[i] 恰好与 nums[x−1] 相等，那么就会无限交换下去。
// 此时在交换时需要再加一个限制条件，即当nums[i] != nums[x−1]时，才进行交换说明。
//
// 最后再进行一次遍历，当数据下标i的值 x 不等于 i+1 时，说明数字 i+1 就是缺失的正整数
// 如果最后遍历完还没有找到，则说明该数据时依次的，直接返回N+1即可
func FirstMissingPositiveExchange(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		// nums[i] ∈ [1,n]，且 nums[i] != nums[nums[i]-1] 时交换
		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
