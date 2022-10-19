package array

// 题目：
// 在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。

// 实例：
/*
	Given nums = [2, 7, 11, 15], target = 9,

	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1]
*/

// TwoSumSelf 自己的思路
// 两层循环，内层循环的开始索引位置，是外层循环所处索引位置加1
func TwoSumSelf(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// TwoSumBook 书中思路
// 这道题最优的做法时间复杂度是 O(n)。
// 顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字。
// 如果找到了，直接返回 2 个数字的下标即可。
// 如果找不到，就把当前循环这个数字存入 map 中，key为当前数字，value当前数字在切片中的索引，等待扫到“另一半”数字的时候，再取出来返回结果。
func TwoSumBook(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if n, ok := m[another]; ok {
			return []int{n, i}
		}
		m[nums[i]] = i
	}
	return nil
}
