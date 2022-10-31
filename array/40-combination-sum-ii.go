package array

import "sort"

/*
解释语句: if cur > begin && candidates[cur-1] == candidates[cur] 是如何避免重复的。

这个避免重复当思想是在是太重要了。
这个方法最重要的作用是，可以让同一层级，不出现相同的元素。即
                  1
                 / \
                2   2  这种情况不会发生 但是却允许了不同层级之间的重复即：
               /     \
              5       5
                例2
                  1
                 /
                2      这种情况确是允许的
               /
              2

为何会有这种神奇的效果呢？
首先 cur-1 == cur 是用于判定当前元素是否和之前元素相同的语句。这个语句就能砍掉例1。
可是问题来了，如果把所有当前与之前一个元素相同的都砍掉，那么例二的情况也会消失。
因为当第二个2出现的时候，他就和前一个2相同了。

那么如何保留例2呢？
那么就用cur > begin 来避免这种情况，你发现例1中的两个2是处在同一个层级上的，
例2的两个2是处在不同层级上的。
在一个for循环中，所有被遍历到的数都是属于一个层级的。
我们要让一个层级中，必须出现且只出现一个2，那么就放过第一个出现重复的2，但不放过后面出现的2。
第一个出现的2的特点就是 cur == begin. 第二个出现的2 特点是 cur > begin.
*/

// 40、组合总和 II
// 这一题是第 39 题的加强版，第 39 题中元素可以重复利用(重复元素可无限次使用)，这一题中元素只能有限次数的利用，因为存在重复元素，并且每个元素只能用一次(重复元素只能使用有限次)

/* 题目
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用 一次 。
注意：解集不能包含重复的组合。
*/

/* 实例：
示例 1:
输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
	[1,1,6],
	[1,2,5],
	[1,7],
	[2,6]
]

实例 2：
输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
	[1,2,2],
	[5]
]
*/

func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var temp []int
	var dfs func(int, int)
	dfs = func(target, index int) {
		if target == 0 {
			res = append(res, append([]int(nil), temp...))
			return
		}
		// 前面代码没多啥变化

		// 考虑的就是比如1可以和多个2组成一组，但不可以分别和多个2组成一组，因为属于重复
		for i := index; i < len(candidates); i++ { // 通过for循环来往下递归
			// 剪枝：如果当前数和上一个数相同，上一个数是组过队了，所以没必要在组一次
			if i-1 >= index && candidates[i] == candidates[i-1] {
				continue
			}
			if target-candidates[i] >= 0 {
				temp = append(temp, candidates[i])
				dfs(target-candidates[i], i+1)
				temp = temp[:len(temp)-1]
			} else { // 剪枝：对于排序过的数组，当前数大于target后面就不用比了
				return
			}
		}
	}
	dfs(target, 0)
	return res
}
