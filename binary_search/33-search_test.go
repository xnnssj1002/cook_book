package binary_search

import "testing"

func TestSearch(t *testing.T) {
	sources := []struct {
		name   string
		nums   []int
		target int
		hope   int
	}{
		{
			name:   "test0",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			hope:   4,
		},
		{
			name:   "test1",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			hope:   -1,
		},
		{
			name:   "test2",
			nums:   []int{1},
			target: 0,
			hope:   -1,
		},
	}
	for _, source := range sources {
		r := search(source.nums, source.target)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %d, got %d\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}
