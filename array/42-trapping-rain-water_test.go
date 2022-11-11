package array

import (
	"testing"
)

var testCases42 = []struct {
	Case   []int
	Target int
	Name   string
}{
	{
		Case:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
		Target: 6,
		Name:   "Case1",
	},
	{
		Case:   []int{4, 2, 0, 3, 2, 5},
		Target: 9,
		Name:   "Case2",
	},
}

func TestTrapDoublePointForce(t *testing.T) {
	for _, testCase := range testCases42 {
		got := TrapDoublePointForce(testCase.Case)
		if got != testCase.Target {
			t.Errorf("case 【%s】 expected %d, but got %d\n", testCase.Name, testCase.Target, got)
		}
	}
}

func TestTrapDoublePointForceOffice(t *testing.T) {
	for _, testCase := range testCases42 {
		got := TrapDoublePointForceOffice(testCase.Case)
		if got != testCase.Target {
			t.Errorf("case 【%s】 expected %d, but got %d\n", testCase.Name, testCase.Target, got)
		}
	}
}
