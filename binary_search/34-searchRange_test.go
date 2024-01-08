package binary_search

import "testing"

var searchRangeSources = []struct {
	name   string
	nums   []int
	target int
	hope   []int
}{
	{
		name:   "test0",
		nums:   []int{5, 7, 7, 8, 8, 10},
		target: 8,
		hope:   []int{3, 4},
	},
	{
		name:   "test1",
		nums:   []int{5, 7, 7, 8, 8, 10},
		target: 6,
		hope:   []int{-1, -1},
	},
	{
		name:   "test2",
		nums:   []int{},
		target: 0,
		hope:   []int{-1, -1},
	},
}

func TestSearchRange(t *testing.T) {

	for _, source := range searchRangeSources {
		r := searchRange(source.nums, source.target)
		if r == nil || r[0] != source.hope[0] || r[1] != source.hope[1] {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}

func TestSearchRangeMid(t *testing.T) {
	for _, source := range searchRangeSources {
		r := searchRangeMid(source.nums, source.target)
		if r == nil || r[0] != source.hope[0] || r[1] != source.hope[1] {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}
