package binary_search

import "testing"

func TestBinarySearchMatrix(t *testing.T) {
	target := []struct {
		name   string
		nums   []int
		target int
		hope   int
	}{
		{
			name:   "test0",
			nums:   []int{1, 3, 4, 5, 6},
			target: 1,
			hope:   0,
		},
		{
			name:   "test1",
			nums:   []int{1, 3, 4, 5, 6},
			target: 6,
			hope:   4,
		},
		{
			name:   "test2",
			nums:   []int{1, 3, 4, 5, 6},
			target: 3,
			hope:   1,
		},
		{
			name:   "test3",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			hope:   2,
		},
		{
			name:   "test4",
			nums:   []int{1, 3, 5, 6},
			target: 4,
			hope:   -1,
		},
	}
	for _, s := range target {
		r := binarySearchMatrix(s.nums, s.target)
		if r != s.hope {
			t.Errorf("【%s】hope %d, got %d\n", s.name, s.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", s.name)
		}
	}
}

func TestSearchFirstEqualElement(t *testing.T) {
	target := []struct {
		name   string
		nums   []int
		target int
		hope   int
	}{
		{
			name:   "test0",
			nums:   []int{1, 1, 1, 1, 1},
			target: 1,
			hope:   0,
		},
		{
			name:   "test1",
			nums:   []int{1, 1, 1, 1},
			target: 1,
			hope:   0,
		},
		{
			name:   "test2",
			nums:   []int{1, 3, 6, 6, 6},
			target: 6,
			hope:   2,
		},
		{
			name:   "test3",
			nums:   []int{1, 3, 6, 6},
			target: 6,
			hope:   2,
		},
		{
			name:   "test4",
			nums:   []int{1, 1, 5, 5},
			target: 4,
			hope:   -1,
		},
	}
	for _, s := range target {
		r := searchFirstEqualElement(s.nums, s.target)
		if r != s.hope {
			t.Errorf("【%s】hope %d, got %d\n", s.name, s.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", s.name)
		}
	}
}
