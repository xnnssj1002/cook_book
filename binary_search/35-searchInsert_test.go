package binary_search

import "testing"

func TestSearchInsert(t *testing.T) {
	sources := []struct {
		name   string
		nums   []int
		target int
		hope   int
	}{
		{
			name:   "test0",
			nums:   []int{1, 3, 5, 6},
			target: 5,
			hope:   2,
		},
		{
			name:   "test1",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			hope:   1,
		},
		{
			name:   "test2",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			hope:   4,
		},
	}
	for _, source := range sources {
		r := searchInsert(source.nums, source.target)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %d, got %d\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}
