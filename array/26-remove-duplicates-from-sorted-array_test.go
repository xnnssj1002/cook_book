package array

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	res1 := RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(res1)
}

func TestRemoveDuplicatesFastSlowPoint(t *testing.T) {
	res1 := RemoveDuplicatesFastSlowPoint([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(res1)
}
