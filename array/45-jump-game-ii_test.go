package array

import "testing"

var testCases45 = []struct {
	Case   []int
	Target int
	Name   string
}{
	{
		Case:   []int{2, 3, 1, 1, 4},
		Target: 2,
		Name:   "case1",
	},
	{
		Case:   []int{2, 3, 0, 1, 4},
		Target: 2,
		Name:   "case2",
	},
	{
		Case:   []int{2, 1},
		Target: 1,
		Name:   "case3",
	},
	{
		Case:   []int{3, 2, 1},
		Target: 1,
		Name:   "case4",
	},
}

func TestJump(t *testing.T) {
	for _, testCase := range testCases45 {
		got := Jump(testCase.Case)
		if got != testCase.Target {
			t.Errorf("case 【%s】 expected %d, but got %d\n", testCase.Name, testCase.Target, got)
		}
	}
}

func TestJumpForwardMax(t *testing.T) {
	for _, testCase := range testCases45 {
		got := JumpForwardMax(testCase.Case)
		if got != testCase.Target {
			t.Errorf("case 【%s】 expected %d, but got %d\n", testCase.Name, testCase.Target, got)
		}
	}
}

func TestJumpForwardMaxOptimization(t *testing.T) {
	for _, testCase := range testCases45 {
		got := JumpForwardMaxOptimization(testCase.Case)
		if got != testCase.Target {
			t.Errorf("case 【%s】 expected %d, but got %d\n", testCase.Name, testCase.Target, got)
		}
	}
}
