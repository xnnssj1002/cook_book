package sliding_window

import "testing"

var minSubArrayLenSources = []struct {
	name   string
	nums   []int
	target int
	hope   int
}{
	{
		name:   "test0",
		nums:   []int{2, 3, 1, 2, 4, 3},
		target: 7,
		hope:   2,
	},
	{
		name:   "test1",
		nums:   []int{1, 4, 4},
		target: 4,
		hope:   1,
	},
	{
		name:   "test2",
		nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
		target: 11,
		hope:   0,
	},
}

func TestMinSubArrayLen(t *testing.T) {

	for _, source := range minSubArrayLenSources {
		r := minSubArrayLen1(source.target, source.nums)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)

		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}

}

func TestMinSubArrayLen1Optimize(t *testing.T) {

	for _, source := range minSubArrayLenSources {
		r := minSubArrayLenStruct1Optimize1(source.target, source.nums)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)

		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}

}

func TestMinSubArrayLen2Optimize(t *testing.T) {

	for _, source := range minSubArrayLenSources {
		r := minSubArrayLenStruct1Optimize2(source.target, source.nums)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)

		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}

}
