package array

import (
	"fmt"
	"testing"
)

var testCases46 = []struct {
	Case      []int
	TargetLen int
	Name      string
}{
	{
		Case:      []int{1, 2, 3},
		TargetLen: 6,
		Name:      "case1",
	},
	{
		Case:      []int{0, 1},
		TargetLen: 2,
		Name:      "case2",
	},
	{
		Case:      []int{1},
		TargetLen: 1,
		Name:      "case3",
	},
	{
		Case:      []int{},
		TargetLen: 0,
		Name:      "case4",
	},
}

func TestPermuteForce(t *testing.T) {
	for _, testCase := range testCases46 {
		got := Permute(testCase.Case)
		fmt.Printf("case 【%s】 res is %v\n", testCase.Name, got)
		if len(got) != testCase.TargetLen {
			t.Errorf("case 【%s】 expected len %d, but got len %d\n", testCase.Name, testCase.TargetLen, len(got))
		}
	}
}
