package binary_search

// 4. 寻找两个正序数组的中位数
/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

示例 2：
输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
*/

// findMedianSortedArrays1 使用双指针的方法来进行
// 1 3 4 5 7
// 2 3 6 7 9
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	// 为了求中位数，且兼容总数是奇偶，需要先找到中间的两个数，然后求这两个数的平均值即可
	// 假设两个数组的总长度为L。对偶数、奇数分别求中间的两个数
	// 偶数：midLeft = L/2；midRight = L/2+1。例如：总数为10，midLeft=10/2=5；midRight=10/2+1=6
	// 奇数：midLeft = L/2+1；midRight = L/2+1。例如：总数为5，midLeft=5/2+1=3；midRight=5/2+1=3
	// 注意：为了兼容偶数与奇数，可以实现一个统一计算中间两个数的逻辑。即：midLeft = (L+1)/2；midRight = (L+2)/2。例子如下：
	// 偶数：总数为10，midLeft=(10+1)/2=5；midRight=(10+2)/2=6
	// 奇数：总数为5， midLeft=(5+1)/2=3； midRight=(5+2)/2=3

	L := len(nums1) + len(nums2)

	//maxIndexLeft := (L+1)/2 - 1
	//maxIndexLeftNum := 0.0

	//maxIndexRight := (L+2)/2 - 1
	//maxIndexRightNum := 0.0

	for i := 0; i < L/2+1; i++ {

	}

	return 0
}

//  1,3  0
//  2    0
//  问题是寻找中位数
//  中位数，特性，有一半比他大，有一半比他小
//  将中位数 切分为 左边的都比右边的小，只要切到正常位置即可，因为是有序的
//  偶数， 左边 m+n /2 ， 中位数就是 m+n/2  m+n /2 +1
//  基数， 左边 m+n /2   中位数是 m+n/2+1
//  寻找中位线 m+n / 2
//  取中位值，那就是 取中位线左边的最大值和右边的最大值，
//  由于中位数靠左属性, 1的时候会有问题  [], [1]
//  时钟存在panic 数组下标的问题
//  且数组
//  0 d 1 d 2 d

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	s, e := 0, m

	for s <= e {
		i := s + (e-s)/2
		j := (m+n)/2 - i
		// i-1 i
		// j-1 j
		if i-1 >= 0 && j < n && nums1[i-1] > nums2[j] { // i-1 < j
			e = i - 1
		} else if j-1 >= 0 && i < m && nums1[i] < nums2[j-1] { // j-1 > i
			s = i + 1
		} else { // 符合要求
			maxLeft := 0
			// -1 0
			// j-1 j
			if i == 0 {
				if j != 0 {
					maxLeft = nums2[j-1]
				}
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}

			minRight := 0
			// m-1 m -1
			// j-1 j j+1
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums1[i], nums2[j])
			}

			if (m+n)%2 == 1 {
				return float64(minRight)
			} else {
				return float64(maxLeft+minRight) / 2
			}
		}
	}
	return float64(0)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}