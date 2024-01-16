package sort

// quickSort 快速排序
// 快速排序分为两步
// 一、哨兵划分：选取一个基准数，将数组中比基准数大的元素移动到基准数右侧，比他小的元素移动到基准数左侧。
// 1、从当前数组中找到一个基准数 pivot（这里以当前数组第 1 个元素作为基准数，即 pivot = nums[low]）。
// 2、使用指针 i 指向数组开始位置，指针 j 指向数组末尾位置。
// 3、从右向左移动指针 j，找到第 1 个小于基准值的元素。
// 4、从左向右移动指针 i，找到第 1 个大于基准数的元素。
// 5、交换指针 i、指针 j 指向的两个元素位置。
// 6、重复第 3∼5 步，直到指针 i 和指针 j 相遇时停止，最后将基准数放到两个子数组交界的位置上。
//
// 二、递归分解：完成哨兵划分之后，对划分好的左右子数组分别进行递归排序。
// 1、按照基准数的位置将数组拆分为左右两个子数组。
// 2、对每个子数组分别重复「哨兵划分」和「递归分解」，直到各个子数组只有 1 个元素，排序结束。
func quickSort(nums []int) []int {
	recursionQuickSort(nums, 0, len(nums)-1)
	return nums
}

// recursionQuickSort 递归快速排序
func recursionQuickSort(nums []int, left, right int) {
	// 子数组长度为 1 时终止递归
	if left >= right {
		return
	}

	// 哨兵划分
	pivot := partition(nums, left, right)

	// 递归左子数组、右子数组
	recursionQuickSort(nums, left, pivot-1)
	recursionQuickSort(nums, pivot+1, right)
}

// partition 哨兵划分
func partition(nums []int, left, right int) int {
	// 以 nums[left] 为基准数
	i, j := left, right

	for i < j {
		// 1、先移动右边的指针
		for i < j && nums[j] >= nums[left] {
			j-- // 从右向左找首个小于基准数的元素
		}
		// 2、在移动左边的指针
		for i < j && nums[i] <= nums[left] {
			i++ // 从左向右找首个大于基准数的元素
		}
		// 元素交换
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 将基准数交换至两子数组的分界线
	nums[i], nums[left] = nums[left], nums[i]
	return i // 返回基准数的索引
}
