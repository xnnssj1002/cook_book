package array

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	res := RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	fmt.Println(res)
}
