package two_points

import "sort"

/*
349. 两个数组的交集
给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。

示例 1：
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]

示例 2：
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
解释：[4,9] 也是可通过的
*/

// 使用两个map进行去重，然后在遍历最短的map
func intersection1(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})

	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}

	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}

	var intersection []int
	for v := range set1 {
		if _, ok := set2[v]; ok {
			intersection = append(intersection, v)
		}
	}

	return intersection
}

// 先排序，再使用双指针
// 注意去重
func intersection2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	left1, left2 := 0, 0

	var intersection []int
	for left1 < len(nums1) && left2 < len(nums2) {
		if nums1[left1] == nums2[left2] {
			// 在插入元素时，进行去重处理。
			// 在插入元素之前，判断已经存在元素切片的最后一个元素，是否与当前元素相等。若相等，不添加；不想等，进行添加
			if len(intersection) == 0 || intersection[len(intersection)-1] != nums1[left1] {
				intersection = append(intersection, nums1[left1])
			}
			left1++
			left2++

			// 谁小，谁右移
		} else if nums1[left1] < nums2[left2] {
			left1++
		} else {
			left2++
		}

	}
	return intersection
}

// 对分离双指针进行优化。将双指针定义为局部变量
func intersection2Optimize(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var intersection []int
	for left1, left2 := 0, 0; left1 < len(nums1) && left2 < len(nums2); {
		if nums1[left1] == nums2[left2] {
			// 在插入元素时，进行去重处理。
			// 在插入元素之前，判断已经存在元素切片的最后一个元素，是否与当前元素相等。若相等，不添加；不想等，进行添加
			if len(intersection) == 0 || intersection[len(intersection)-1] != nums1[left1] {
				intersection = append(intersection, nums1[left1])
			}
			left1++
			left2++

			// 谁小，谁右移
		} else if nums1[left1] < nums2[left2] {
			left1++
		} else {
			left2++
		}

	}
	return intersection
}
