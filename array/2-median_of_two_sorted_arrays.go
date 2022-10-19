package array

import (
	"fmt"
	"strconv"
)

// 题目：
/*
	给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
	请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
	你可以假设 nums1 和 nums2 不会同时为空。
*/

// 实例1：
/*
	nums1 = [1, 3]
	nums2 = [2]

	The median is 2.0
*/

// 实例2：
/*
	nums1 = [1, 2]
	nums2 = [3, 4]

	The median is (2 + 3)/2 = 2.5
*/

/* 两个数保留两位小数
func main() {
	var a int = 3
	var b int = 7
	fmt.Println(a / b)  // 结果为0，不符合预期
	num0, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", a/b), 64)                 // 保留2位小数
	fmt.Println(num0)   // 结果为0，不符合预期
	num1, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(a)/float64(b)), 64) // 保留2位小数
	fmt.Println(num1)   // 0.43
	num2, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", float64(a)/float64(b)), 64) // 保留5位小数
	fmt.Println(num2)   // 0.42857
}
*/

/*判断一个数字是奇数还是偶数，通过 按位& 1运算，当 按位& 1后，为 0 代表该数为偶数，为 1 代表该数为奇数
按位&，两个数的同一位都为1时，结果才为1，换言之，只要有一个0，结果就是0
与上一个1，该1表示最低位是，其他高位全部位0，与另一个数与时，结果是其他高位全部位0。
如果这个数是偶数，最低位必须是0，与1时，最低位还是0，因为0&1 = 0，随意按位与1时，结果为0表示偶数。
对于奇数同样的道理，奇数最低位是1，与上一个1，最低位是1，即1&1=1，结果为1表示奇数
*/

// FindMedianSortedArraysMerge 自己的思路
// 现将两个切片排好序后，形成一个新切片，然后对新切片取中位数
func FindMedianSortedArraysMerge(nums1 []int, nums2 []int) float64 {
	newNums := make([]int, 0)
	nums1Index := 0
	nums2Index := 0
	for {
		if nums1Index >= len(nums1) || nums2Index >= len(nums2) {
			break
		}
		if nums1[nums1Index] < nums2[nums2Index] {
			newNums = append(newNums, nums1[nums1Index])
			nums1Index++
		} else {
			newNums = append(newNums, nums2[nums2Index])
			nums2Index++
		}
	}
	if nums1Index < len(nums1) {
		newNums = append(newNums, nums1[nums1Index:]...)
	}
	if nums2Index < len(nums2) {
		newNums = append(newNums, nums2[nums2Index:]...)
	}
	fmt.Println("合并后的切片为:", newNums)

	if len(newNums)%2 == 1 { // 奇数
		middleIndex := len(newNums) / 2
		float, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(newNums[middleIndex])), 64)
		return float

	} else { // 偶数
		middleIndex := len(newNums) / 2
		res := float64(newNums[middleIndex]+newNums[middleIndex-1]) / 2
		float, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", res), 64)
		return float
	}

}

// FindMedianSortedArraysLoseOne 首先知道了两个数组的总元素个数，找到中位数所需的两个元素位置（奇数按照两个相同来寻找）。
// 借助合并两个数组的方法，从两个数组开头依次遍历，一个一个丢弃，直到找到所需的中位数
func FindMedianSortedArraysLoseOne(nums1 []int, nums2 []int) float64 {
	// 获取两个数组总元素个数
	sum := len(nums1) + len(nums2)

	// 根据 奇数还是偶数 得到计算中位数所需两个元素的索引值
	startIndex, endIndex := 0, 0
	if sum&1 == 0 { // 偶数
		startIndex = sum/2 - 1
		endIndex = sum / 2
	} else {
		startIndex, endIndex = sum/2, sum/2
	}
	fmt.Printf("总数组 开始索引位置:%d, 结束索引位置:%d\n", startIndex, endIndex)

	// 计算中位数，所需的两个元素值
	startNum, endNum := 0, 0

	// 已经循环到的总元素数组的索引值
	arriveIndex := 0

	// 分别记录循环到两个数组的索引值
	nums1Index, nums2Index := 0, 0
	for {
		if arriveIndex > startIndex && arriveIndex > endIndex {
			break
		}

		curNum := 0
		if nums1Index > len(nums1)-1 { // 超出 num1 范围，取 num2 的元素
			curNum = nums2[nums2Index]
			nums2Index++
		} else if nums2Index > len(nums2)-1 { // 超出 num2 范围，取 num1 的元素
			curNum = nums1[nums1Index]
			nums1Index++
		} else if nums1[nums1Index] < nums2[nums2Index] { // num1 num2 都没有超出范围，取元素最小的
			curNum = nums1[nums1Index]
			nums1Index++
		} else {
			curNum = nums2[nums2Index]
			nums2Index++
		}

		if arriveIndex == startIndex {
			startNum = curNum
		}
		if arriveIndex == endIndex {
			endNum = curNum
		}

		arriveIndex++
	}

	res := float64(startNum+endNum) / 2
	float, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", res), 64)
	return float
}

// FindMedianSortedArraysLoseHalf leetcode 上的解题思路，一半一半的丢弃
// 该场景就是，求两个有序数组第 k 小的数的一种特殊情况，先实现该场景，例如方法： FindTwoSortedArraysKNums
// 如果两个数组的总元素个数为偶数，比如：6，中位数就是 第3小的数 与 第4小的数 的平均数
// 如果两个数组的总元素个数为技术，比如：7，中位数就是 第4小的数。如果不区分奇数还是偶数边界的话，可以求两次 第4小的数 进行求平均，这样就和偶数一致了
// 由于数列是有序的，完全可以一半儿一半儿的排除。假设我们要找第 k 小数，我们可以每次循环排除掉 k/2 个数。
func FindMedianSortedArraysLoseHalf(nums1 []int, nums2 []int) float64 {
	// 计算中位数所需的两个位置
	startPosition := (len(nums1) + len(nums2) + 1) / 2
	endPosition := (len(nums1) + len(nums2) + 2) / 2

	// 计算中位数所需的两个元素
	startNum := FindTwoSortedArraysKNums(nums1, nums2, startPosition)
	endNum := FindTwoSortedArraysKNums(nums1, nums2, endPosition)

	res := float64(startNum+endNum) / 2
	float, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", res), 64)
	return float
}

// FindTwoSortedArraysKNums 求两个有序数组第 k 小的数
func FindTwoSortedArraysKNums(nums1 []int, nums2 []int, target int) int {
	if target > len(nums1)+len(nums2) || target <= 0 {
		return 0
	}

	// 如果有一个数组为空，直接从另一个数组中减去 target，然后继续
	if len(nums1) == 0 {
		return nums2[target-1]
	}
	if len(nums2) == 0 {
		return nums1[target-1]
	}

	// 接下来两个数组都不为空
	if target == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}

	middlePosition := target / 2
	middleTargetIndex := middlePosition - 1

	if middleTargetIndex >= len(nums1) { // 一半的位置超过了nums1的范围
		if nums1[len(nums1)-1] >= nums2[middleTargetIndex] { // 减去nums2
			return FindTwoSortedArraysKNums(nums1, nums2[middlePosition:], target-middlePosition)
		} else {
			return FindTwoSortedArraysKNums([]int{}, nums2, target-len(nums1))
		}
	}
	if middleTargetIndex >= len(nums2) {
		if nums2[len(nums2)-1] >= nums1[middleTargetIndex] { // 减去nums1
			return FindTwoSortedArraysKNums(nums1[middlePosition:], nums2, target-middlePosition)
		} else {
			return FindTwoSortedArraysKNums(nums1, []int{}, target-len(nums2))
		}
	}

	if nums1[middleTargetIndex] < nums2[middleTargetIndex] {
		return FindTwoSortedArraysKNums(nums1[middlePosition:], nums2, target-middlePosition)
	} else {
		return FindTwoSortedArraysKNums(nums1, nums2[middlePosition:], target-middlePosition)
	}

}
