package array

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	res := SearchInsertCode1([]int{1, 3}, 0)
	fmt.Println(res)
	res1 := SearchInsertCod2([]int{1, 3, 5, 6}, 5)
	fmt.Println(res1)
	res2 := SearchInsert([]int{1, 3, 5, 6}, 7)
	fmt.Println(res2)
}
