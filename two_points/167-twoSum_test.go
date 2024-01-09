package two_points

import "testing"

var twoSumSources = []struct {
	name   string
	nums   []int
	target int
	hope   []int
}{
	{
		name:   "test0",
		nums:   []int{2, 7, 11, 15},
		target: 9,
		hope:   []int{1, 2},
	},
	{
		name:   "test1",
		nums:   []int{2, 3, 4},
		target: 6,
		hope:   []int{1, 3},
	},
	{
		name:   "test2",
		nums:   []int{-1, 0},
		target: -1,
		hope:   []int{1, 2},
	},
	{
		name:   "test3",
		nums:   []int{3, 6},
		target: 8,
		hope:   []int{-1, -1},
	},
}

func TestTwoSumSort(t *testing.T) {

	for _, source := range twoSumSources {
		r := twoSumSort(source.nums, source.target)
		if r == nil || r[0] != source.hope[0] || r[1] != source.hope[1] {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}
