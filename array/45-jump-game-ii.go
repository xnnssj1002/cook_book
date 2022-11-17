package array

/* 45. 跳跃游戏 II
给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置。
*/

/* 示例
示例 1:
输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

示例 2:
输入: nums = [2,3,0,1,4]
输出: 2
*/

// Jump 反向查找出发位置
// 1、目标是到达数组的最后一个位置，因此我们可以考虑最后一步跳跃前所在的位置，该位置通过跳跃能够到达最后一个位置。
// 2、如果有多个位置通过跳跃都能够到达最后一个位置，那么我们应该如何进行选择呢？
// 直观上来看，可以「贪心」地选择距离最后一个位置最远的那个位置，也就是对应下标最小的那个位置。因此，我们可以从左到右遍历数组，选择第一个满足要求的位置。
// 3、找到最后一步跳跃前所在的位置之后，我们继续贪心地寻找倒数第二步跳跃前所在的位置，以此类推，直到找到数组的开始位置。
func Jump(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return 0
	}
	position := length - 1 // 开始默认位置为最后，即索引位置为 length-1
	steps := 0
	for position > 0 {
		for i := 0; i < position; i++ { // 当前位置i的元素nums[i] 大于等于 最后位置与当前位置i的差值，证明【贪心】的找到最前的位置
			if nums[i] >= position-i {
				steps++
				position = i
			}
		}
	}
	return steps
}

// JumpForwardMax 正向查找可到达的最大位置
// 1、如果某一个作为 起跳点 的格子可以跳跃的距离是 3，那么表示后面 3 个格子都可以作为 起跳点。
//	11. 可以对每一个能作为 起跳点 的格子都尝试跳一次，把 能跳到最远的距离 不断更新。
// 2、如果从这个 起跳点 起跳叫做第 1 次 跳跃，那么从后面 3 个格子起跳 都 可以叫做第 2 次 跳跃。
// 3、所以，当一次 跳跃 结束时，从下一个格子开始，到现在 能跳到最远的距离，都 是下一次 跳跃 的 起跳点。
//	31. 对每一次 跳跃 用 for 循环来模拟。
//	32. 跳完一次之后，更新下一次 起跳点 的范围。
//	33. 在新的范围内跳，更新 能跳到最远的距离。
// 4、记录 跳跃 次数，如果跳到了终点，就得到了结果。
// 优化版请看 JumpForwardMaxOptimization
func JumpForwardMax(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return 0
	}
	steps := 0 // 步数
	start := 0 // 起跳点范围开始的位置
	end := 0   // 起跳点范围结束的位置

	for end < length-1 { // 只要结束的位置不到最后一个元素，就继续进行循环，即需要继续跳
		// 寻找下一个起点。下一个起点在 [ start+1, start+nums[start] ] 中其中一个
		// 取哪一个呢？结果为：所以的下一个起点遍历一遍，从所有的下一个起点起跳，那个起点跳的更远，就取哪个作为真正下一个起点
		maxNum := 0
		for i := start; i <= end; i++ {
			if i+nums[i] > maxNum {
				maxNum = i + nums[i]
			}
		}
		start = end + 1 // 下一次起跳点范围开始的格子
		end = maxNum    // 下一次起跳点范围结束的格子
		steps++
	}

	return steps
}

// JumpForwardMaxOptimization 正向查找可到达的最大位置 优化版
// 1、从上面代码观察发现，其实内层的 for 循环中，i 是从头跑到尾的。
// 2、只需要在一次 跳跃 完成时，更新下一次 能跳到最远的距离。
// 3、并以此刻作为时机来更新 跳跃 次数。
// 4、就可以在一次 for 循环中处理。
func JumpForwardMaxOptimization(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return 0
	}
	step := 0
	end := 0
	maxPosition := 0
	// 注意这里不需要循环最后一个元素
	// 因为i=0时就记录步数了，说明时从每次开始跳记录步数的。在倒数第二个记录的话，最后肯定能跳到结尾
	for i := 0; i < length-1; i++ {
		if i+nums[i] > maxPosition {
			maxPosition = i + nums[i]
		}
		if i == end { // 找到下一次跳跃的起点了
			end = maxPosition
			step++
		}
	}
	return step
}
