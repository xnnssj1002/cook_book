package sliding_window

/*
1343. 大小为 k 且平均值大于等于阈值的子数组数目
给你一个整数数组 arr 和两个整数 k 和 threshold 。
请你返回长度为k且平均值【大于等于】threshold 的【子数组】数目。

示例 1：
输入：arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4
输出：3
解释：子数组 [2,5,5],[5,5,5] 和 [5,5,8] 的平均值分别为 4，5 和 6 。其他长度为 3 的子数组的平均值都小于 4 （threshold 的值)。

示例 2：
输入：arr = [11,13,17,23,29,31,7,5,2,3], k = 3, threshold = 5
输出：6
解释：前 6 个长度为 3 的子数组平均值都大于 5 。注意平均值不是整数。
*/

func numOfSubArrays1(arr []int, k int, threshold int) int {
	left, right := 0, 0

	res := 0
	sum := 0

	// 循环退出条件为 right == len(arr)。即右指针循环到数组的最右边节点
	for right < len(arr) {
		sum += arr[right]

		if right-left+1 == k {
			if sum >= k*threshold { // 此时窗口满足条件
				res += 1
			}
			sum -= arr[left]
			left++
			right++

		} else {
			right++
		}
	}

	return res
}

// 该方法是对numOfSubArrays1的改写
func numOfSubArrays2(arr []int, k int, threshold int) int {
	res := 0
	sum := 0
	// 循环退出条件为 right == len(arr)。即右指针循环到数组的最右边节点
	for left, right := 0, 0; right < len(arr); right++ {
		sum += arr[right]
		if right-left+1 == k {
			if sum >= k*threshold { // 此时窗口满足
				res += 1
			}
			sum -= arr[left]
			left++
		}
	}

	return res
}

// 使用一个指针i，判断sum + arr[i]-arr[i-k] 与 k*threshold 的大小
func numOfSubArrays3(arr []int, k int, threshold int) int {
	res := 0
	if len(arr) < k {
		return res
	}

	// 1、先将前k个数字进行求和
	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}

	// 2、判断最开始的k个数是否满足条件
	if sum >= k*threshold {
		res++
	}

	// 3、index从k开始循环，直到 len(arr)-1
	for index := k; index < len(arr); index++ {
		// 判断i后边k个元素的总和
		sum = sum + arr[index] - arr[index-k]
		if sum >= k*threshold {
			res++
		}
	}

	return res
}
