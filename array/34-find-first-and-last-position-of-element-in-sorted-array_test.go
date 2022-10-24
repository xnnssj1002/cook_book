package array

import (
	"fmt"
	"testing"
)

func TestSearchRange(t *testing.T) {
	res := SearchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println(res)
}
