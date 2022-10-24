package array

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	res := Search([]int{1, 3}, 0)
	fmt.Println(res)

	res1 := Search([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	fmt.Println(res1)
}

func TestSearchBinary(t *testing.T) {
	res0 := SearchBinary([]int{3, 1}, 1)
	fmt.Println(res0)

	res := SearchBinary([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	fmt.Println(res)

	res1 := SearchBinary([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	fmt.Println(res1)
}
