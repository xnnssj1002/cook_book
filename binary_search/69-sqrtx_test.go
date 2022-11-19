package binary_search

import (
	"fmt"
	"testing"
)

var testCases69 = []struct {
	Case   int
	Target int
	Name   string
}{
	{
		Case:   1,
		Target: 1,
		Name:   "case1",
	},
	{
		Case:   2,
		Target: 1,
		Name:   "case2",
	},
	{
		Case:   5,
		Target: 2,
		Name:   "case3",
	},
	{
		Case:   8,
		Target: 2,
		Name:   "case4",
	},
}

func TestPermuteUnique(t *testing.T) {
	for _, testCase := range testCases69 {
		got := MySqrt(testCase.Case)
		fmt.Printf("case 【%s】 res is %v\n", testCase.Name, got)
		if got != testCase.Target {
			t.Errorf("case 【%s】 expected %d, but got len %d\n", testCase.Name, testCase.Target, got)
		}
	}
}
