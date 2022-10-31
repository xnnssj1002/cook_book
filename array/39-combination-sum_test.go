package array

import (
	"fmt"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	res := CombinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println(res)
}

func TestCombinationSumOffice(t *testing.T) {
	res := CombinationSumOffice([]int{2, 3, 6, 7}, 7)
	fmt.Println(res)
}

func TestCombinationSumNetVersion(t *testing.T) {
	res := CombinationSumNetVersion([]int{2, 3, 6, 7}, 7)
	fmt.Println(res)
}
