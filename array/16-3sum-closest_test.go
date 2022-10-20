package array

import (
	"fmt"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	res := ThreeSumClosest([]int{-1, 2, 1, -4}, 1)
	fmt.Println(res)
}
