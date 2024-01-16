package sort

// insertionSort 插入排序
// 「插入排序 insertion sort」的工作原理：在未排序区间选择一个基准元素，将该元素与其左侧已排序区间的元素逐一比较大小，并将该元素插入到正确的位置。
// 可以发现，插入排序与我们现实生活中手动整理一副扑克牌的过程非常相似。
//
// 流程如下：
// 1、初始状态下，数组的第 1 个元素已完成排序。
// 2、选取数组的第 2 个元素作为 base ，将其插入到正确位置后，数组的前 2 个元素已排序。
// 3、选取第 3 个元素作为 base ，将其插入到正确位置后，数组的前 3 个元素已排序。
// 4、以此类推，在最后一轮中，选取最后一个元素作为 base ，将其插入到正确位置后，所有元素均已排序。
func insertionSort(nums []int) []int {
	// 外循环：未排序区间为 [i+1, len(nums)+1]
	for i := 1; i < len(nums); i++ {
		base := nums[i]
		j := i - 1
		// 内循环：将 base 插入到已排序部分的正确位置
		for j >= 0 && nums[j] > base {
			nums[j+1] = nums[j] // 将 nums[j] 向右移动一位
			j--
		}
		nums[j+1] = base // 将 base 赋值到正确位置
	}
	return nums
}
