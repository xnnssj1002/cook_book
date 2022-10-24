package array

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	mySli := []int{4, 5, 2, 6, 3, 1}
	NextPermutation(mySli)
	fmt.Println(mySli)
}
func TestNextPermutationCode(t *testing.T) {
	mySli := []int{4, 5, 2, 6, 3, 1}
	NextPermutationCode(mySli)
	fmt.Println(mySli)
}
