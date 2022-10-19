package array

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArraysSelf(t *testing.T) {
	jiRes := FindMedianSortedArraysMerge([]int{1, 5}, []int{2, 3, 4})
	fmt.Println(jiRes)
	ouRes := FindMedianSortedArraysMerge([]int{1, 5}, []int{2, 3, 4, 6})
	fmt.Println(ouRes)
}

func TestFindMedianSortedArraysLoseOne(t *testing.T) {
	jiRes := FindMedianSortedArraysLoseOne([]int{1, 5}, []int{2, 3, 4})
	fmt.Println(jiRes)
	ouRes := FindMedianSortedArraysLoseOne([]int{1, 5}, []int{2, 3, 4, 6})
	fmt.Println(ouRes)
}

func TestFindMedianSortedArraysLoseHalf(t *testing.T) {
	// [4]
	// [1,2,3,5,6]
	leetCodeRes := FindMedianSortedArraysLoseHalf([]int{2, 3, 4, 5, 6}, []int{1})
	fmt.Println(leetCodeRes)
	jiRes := FindMedianSortedArraysLoseHalf([]int{1, 5}, []int{2, 3, 4})
	fmt.Println(jiRes)
	ouRes := FindMedianSortedArraysLoseHalf([]int{1, 5}, []int{2, 3, 4, 6})
	fmt.Println(ouRes)
}

func TestFindTwoSortedArraysLoseKNums(t *testing.T) {
	res := FindTwoSortedArraysKNums([]int{1, 5}, []int{2, 3, 4}, 2)
	fmt.Println(res)
}
