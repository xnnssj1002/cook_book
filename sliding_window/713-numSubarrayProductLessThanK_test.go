package sliding_window

import "testing"

var numSubarrayProductLessThanKSources = []struct {
	name   string
	nums   []int
	target int
	hope   int
}{
	{
		name:   "test0",
		nums:   []int{10, 5, 2, 6},
		target: 100,
		hope:   8,
	},
	{
		name:   "test1",
		nums:   []int{1, 2, 3},
		target: 0,
		hope:   0,
	},
}

func TestNumSubarrayProductLessThanK(t *testing.T) {

	for _, source := range numSubarrayProductLessThanKSources {
		r := numSubarrayProductLessThanK(source.nums, source.target)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)

		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}

}
