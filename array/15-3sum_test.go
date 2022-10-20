package array

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	res := ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(res)
}

func TestThreeSumCode(t *testing.T) {
	res := ThreeSumCode([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(res)
}
