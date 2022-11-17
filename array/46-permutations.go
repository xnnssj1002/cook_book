package array

/* 46、全排列
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
*/

/* 示例
示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：
输入：nums = [1]
输出：[[1]]
*/

// Permute 递归分治回溯方法
func Permute(nums []int) [][]int {
	currentRes := make([][]int, 0)
	if len(nums) == 1 {
		return append(currentRes, nums)
	}

	for i := 0; i < len(nums); i++ {
		tempData := make([]int, 0)
		tempData = append(tempData, nums[:i]...)
		tempData = append(tempData, nums[i+1:]...)
		tempRes := Permute(tempData)
		for j := 0; j < len(tempRes); j++ {
			currentRes = append(currentRes, append([]int{nums[i]}, tempRes[j]...))
		}
	}

	return currentRes
}
