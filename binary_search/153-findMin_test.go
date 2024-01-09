package binary_search

import "testing"

var findMinSources = []struct {
	name string
	nums []int
	hope int
}{
	{
		name: "test0",
		nums: []int{3, 4, 5, 1, 2},
		hope: 1,
	},
	{
		name: "test1",
		nums: []int{4, 5, 6, 7, 0, 1, 2},
		hope: 0,
	},
	{
		name: "test2",
		nums: []int{11, 13, 15, 17},
		hope: 11,
	},
	{
		name: "test3",
		nums: []int{2, 3, 4, 5, 1},
		hope: 1,
	},
}

func TestFindMin(t *testing.T) {

	for _, source := range findMinSources {
		r := findMin(source.nums)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}
