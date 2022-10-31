package array

import (
	"sort"
)

/* 39. 组合总和
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和 等于目标数 target 的 所有不同组合 ，
并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为 target 的不同组合数少于 150 个。
*/

/*
示例 1：
输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。

示例 2：
输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]

示例 3：
输入: candidates = [2], target = 1
输出: []
*/

// CombinationSum Cook book 书中代码
func CombinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	var c []int
	var res [][]int
	sort.Ints(candidates)
	findCombinationSum(candidates, target, 0, c, &res)
	return res
}

func findCombinationSum(nums []int, target, index int, c []int, res *[][]int) {
	if target <= 0 {
		if target == 0 {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
		return
	}
	for i := index; i < len(nums); i++ {
		if nums[i] > target { // 这里可以剪枝优化
			break
		}
		c = append(c, nums[i])
		findCombinationSum(nums, target-nums[i], i, c, res) // 注意这里迭代的时候 index 依旧不变，因为一个元素可以取多次
		c = c[:len(c)-1]
	}
}

func CombinationSumNetVersion(candidates []int, target int) [][]int {
	sort.Ints(candidates) // 排序一下数组，方便剪枝操作
	var res [][]int
	var temp []int

	var dfs func(int, int)
	dfs = func(target, index int) { // 递归方法，根据目标值和数组下标递归
		// 成功案例结果 刚好target = 0,所以该组合成立
		if target == 0 {
			res = append(res, append([]int(nil), temp...)) // 不可以直接temp不然结果错乱是同一个切片
			return
		}
		// 退出递归条件:到数组末尾
		if index == len(candidates) {
			return
		}
		// 剪枝:如果数组中有重复数，重复的那个直接跳过不考虑,因为数组中的数可以重复使用，重复的数没有意义
		if index > 0 && candidates[index] == candidates[index-1] {
			return
		}
		// 开始递归:使用本身或者使用数组下一个数
		if target-candidates[index] >= 0 { // 大于等于0说明可能组成一种想要的结果
			temp = append(temp, candidates[index]) // 将该结果保存到临时数组里
			dfs(target-candidates[index], index)   // 使用本数index递归，但target是他们的差，因为本数已经是temp的一个解
			temp = temp[:len(temp)-1]              // 剔除该数，恢复到原状，换另外一种情况尝试
		} else { // 剪枝：由于数组排序过，如果当前index的数大于了target，那数组后面的数也同样大于，无需再考虑！
			return
		}

		// 本身数字 尝试完后，测试数组下一个数
		dfs(target, index+1)
	}

	// 从0下标开始递归
	dfs(target, 0)
	return res
}

// CombinationSumOffice 力扣官网解法
func CombinationSumOffice(candidates []int, target int) (ans [][]int) {
	var comb []int
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}
