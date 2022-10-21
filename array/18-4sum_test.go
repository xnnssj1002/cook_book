package array

import (
	"fmt"
	"testing"
)

func TestFourSumTwoPoint(t *testing.T) {
	res := FourSumTwoPoint([]int{-2, -1, -1, 1, 1, 2, 2}, 0)
	fmt.Println(res)

	res1 := FourSumTwoPoint([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Println(res1)

	res2 := FourSumTwoPoint([]int{2, 2, 2, 2}, 8)
	fmt.Println(res2)
}

func TestFourSumTwoPointOptimal(t *testing.T) {
	res := FourSumTwoPointOptimal([]int{-2, -1, -1, 1, 1, 2, 2}, 0)
	fmt.Println(res)

	res1 := FourSumTwoPointOptimal([]int{1, 0, -1, 0, -2, 2}, 0)
	fmt.Println(res1)

	res2 := FourSumTwoPointOptimal([]int{2, 2, 2, 2}, 8)
	fmt.Println(res2)
}
