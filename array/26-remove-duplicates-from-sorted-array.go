package array

import "fmt"

/* 题目
给你一个升序排列的数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次 ，返回删除后数组的新长度。元素的相对顺序应该保持一致 。
由于在某些语言中不能改变数组的长度，所以必须将结果放在数组nums的第一部分。更规范地说，如果在删除重复项之后有 k 个元素，那么 nums 的前 k 个元素应该保存最终结果。
将最终结果插入nums 的前 k 个位置后返回 k 。

不要使用额外的空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
判题标准:
系统会用下面的代码来测试你的题解:

int[] nums = [...]; // 输入数组
int[] expectedNums = [...]; // 长度正确的期望答案

int k = removeDuplicates(nums); // 调用

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
如果所有断言都通过，那么您的题解将被 通过。
*/

/* 示例
输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
*/

// 有序切片去重，返回去重后的元素长度

func RemoveDuplicates(nums []int) int {
	fmt.Println(nums)
	// 特判
	if len(nums) == 0 {
		return 0
	}
	start, end := 0, len(nums)
	for start < end {
		if start > 0 && nums[start] == nums[start-1] {
			nums = append(nums[:start-1], nums[start:]...)
			end--
		} else {
			start++
		}
	}
	fmt.Println(nums)
	return end
}

// RemoveDuplicatesFastSlowPoint 使用快慢指针
func RemoveDuplicatesFastSlowPoint(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	//fast, slow := 1, 1
	//for fast < len(nums) {
	//	if nums[fast] != nums[fast-1] {
	//		nums[slow] = nums[fast]
	//		slow++
	//	}
	//	fast++
	//}

	// 更简洁的写法
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
