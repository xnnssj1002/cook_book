package array

import (
	"fmt"
	"testing"
)

var testCases41 = []struct {
	Case   []int
	Target int
	Name   string
}{
	{
		Case:   []int{1, 2, 0},
		Target: 3,
		Name:   "Case1",
	},
	{
		Case:   []int{3, 4, -1, 1},
		Target: 2,
		Name:   "Case2",
	},
	{
		Case:   []int{7, 8, 9, 11, 12},
		Target: 1,
		Name:   "Case3",
	},
	{
		Case:   []int{},
		Target: 0,
		Name:   "Case4",
	},
	{
		Case:   []int{3, 4, -1, 2},
		Target: 1,
		Name:   "Case5",
	},
}

func TestFirstMissingPositiveSort(t *testing.T) {
	fmt.Println(testCases41)
	for _, testCase := range testCases41 {
		got := FirstMissingPositiveSort(testCase.Case)
		if got != testCase.Target {
			t.Errorf("case 【%s】 expected %d, but got %d\n", testCase.Name, testCase.Target, got)
		}
	}
}

func TestFirstMissingPositiveSameHash(t *testing.T) {
	// 该方法修改了测试用例，如果同时执行的话，会对后边的测试有影响
	fmt.Println(testCases41)
	for _, testCase := range testCases41 {
		got := FirstMissingPositiveSameHash(testCase.Case)
		if got != testCase.Target {
			t.Errorf("case 【%s】 expected %d, but got %d\n", testCase.Name, testCase.Target, got)
		}
	}
}

func TestFirstMissingPositiveExchange(t *testing.T) {
	fmt.Println(testCases41)
	for _, testCase := range testCases41 {
		fmt.Println(testCase)
		got := FirstMissingPositiveExchange(testCase.Case)
		fmt.Println(got)
		if got != testCase.Target {
			t.Errorf("case 【%s】 expected %d, but got %d\n", testCase.Name, testCase.Target, got)
		}
	}
}
