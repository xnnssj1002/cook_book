package binary_search

import "testing"

var findMinDifficultSources = []struct {
	name string
	nums []int
	hope int
}{
	{
		name: "test0",
		nums: []int{1, 3, 5},
		hope: 1,
	},
	{
		name: "test1",
		nums: []int{2, 2, 2, 0, 1},
		hope: 0,
	},
	{
		name: "test2",
		nums: []int{1, 3, 3},
		hope: 1,
	},
}

func TestFindMinDifficult(t *testing.T) {

	for _, source := range findMinDifficultSources {
		r := findMinDifficult(source.nums)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}

}
