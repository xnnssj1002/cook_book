package array

/* 46、全排列II
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
*/

/* 示例
示例 1：
输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]

示例 2：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/

func PermuteUnique(nums []int) [][]int {
	resSli := make([][]int, 0)
	if len(nums) == 1 {
		return append(resSli, nums)
	}

	// 判唯一的map
	alreadyMap := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		if _, exists := alreadyMap[nums[i]]; !exists {
			alreadyMap[nums[i]] = struct{}{}
			// 切割切片
			tempSli := make([]int, 0)
			tempSli = append(tempSli, nums[:i]...)
			tempSli = append(tempSli, nums[i+1:]...)
			tempRes := PermuteUnique(tempSli)
			for j := 0; j < len(tempRes); j++ {
				// 将本次循环的i拼接到然后结果中
				resSli = append(resSli, append([]int{nums[i]}, tempRes[j]...))
			}
		}
	}

	return resSli
}
