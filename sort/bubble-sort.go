package sort

// bubbleSort 冒泡排序
// 「冒泡排序 bubble sort」的工作原理：通过连续地比较与交换相邻元素实现排序
//
// 流程如下：
// 1、首先，对 n 个元素执行“冒泡”，将数组的最大元素交换至正确位置。
// 2、接下来，对剩余 n-1 个元素执行“冒泡”，将第二大元素交换至正确位置。
// 3、以此类推，经过 n-1 轮“冒泡”后，前 n-1 大的元素都被交换至正确位置。
// 4、仅剩的一个元素必定是最小元素，无须排序，因此数组排序完成。
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {

		// 优化点，只要元素之间没有交换，证明元素已经是排好序的了
		trans := false

		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
				trans = true
			}
		}

		if !trans {
			break
		}
	}
	return nums
}
