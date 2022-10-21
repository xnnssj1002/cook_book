package array

import (
	"math"
	"sort"
)

/* 题目
给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
返回这三个数的和。
假定每组输入只存在恰好一个解。
*/

/* 实例
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。

输入：nums = [0,0,0], target = 1
输出：0
*/

// ThreeSumClosest 首先想到的思路
// 1、特判。当切片为 nil，或者切片长度小于3，直接返回 0
// 2、定义一个变量 optimalNum，表示最优解
// 3、定一个变量 diffNum，表示最优解 optimalNum 与目标值 target 差值的绝对值。
// 4、定义三个指针，分别代表三个数值的索引值
// 5、分别对三个指针对应的元素进行求和 sum，与目标值 target 做减法，即sum-target，获取到绝对值A，与diffNum比较。
// 		如果 A < diffNum，则更新最优解 optimalNum = sum，并更新 diffNum = a
//		其他场景，继续下次循环
func ThreeSumClosest(nums []int, target int) int {
	// 1、特判
	if len(nums) < 3 {
		return 0
	}
	// 2、定义最优解
	optimalNum := 0
	// 3、定义差值。即最优解 optimalNum 与 目标值 target 差值的绝对值
	diffNum := 0

	// 4、三层循环，定义三个指针，分别代表三个数值的索引值
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				// 优化点：如果sum与target，直接返回即可
				if sum == target {
					return sum
				}

				A := int(math.Abs(float64(sum - target)))
				// 如果是切片前三个元素，则直接初始化最优解与差值
				if i == 0 && j == 1 && k == 2 {
					optimalNum = sum
					diffNum = A
					continue
				}
				// 5、差值的比较
				if A < diffNum {
					optimalNum = sum
					diffNum = A
				}
			}
		}
	}
	return optimalNum
}

// ThreeSumClosestBy15Question 借助第15题的排序加上双指针来实现
func ThreeSumClosestBy15Question(nums []int, target int) int {
	// 特判
	if len(nums) < 3 {
		return 0
	}
	// 求绝对值函数
	abs := func(x int) int {
		if x < 0 {
			return -1 * x
		}
		return x
	}
	// 初始化返回值变量。默认前三个是最优解
	optimalNum := nums[0] + nums[1] + nums[2]
	// 差值绝对值
	diffNum := abs(target - optimalNum)
	// 排序
	sort.Ints(nums)
	// 双指针判断 left, right
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { // 第一个数字判重处理
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for right > left {
			nowNum := nums[i] + nums[left] + nums[right]
			nowDiffNum := abs(target - nowNum)
			if nowNum == target { // 如果刚好相等，直接返回即可
				return nowNum
			} else if nowNum < target { // 需要左指针右移，让下一个sum增大
				if right > left && nums[left] == nums[left+1] { // 左指针判重
					left++
				}
				left++
			} else { // 右指针左移，让下一个sum减小
				if right > left && nums[right] == nums[right-1] { // 右指针判重
					right--
				}
				right--
			}
			if nowDiffNum < diffNum {
				optimalNum = nowNum
				diffNum = nowDiffNum
			}
		}
	}

	return optimalNum
}
