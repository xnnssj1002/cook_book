package two_points

import "testing"

var removeDuplicatesSources = []struct {
	name string
	nums []int
	hope int
}{
	{
		name: "test0",
		nums: []int{1, 1, 2},
		hope: 2,
	},
	{
		name: "test1",
		nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		hope: 5,
	},
}

func TestRemoveDuplicates(t *testing.T) {

	for _, source := range removeDuplicatesSources {
		r := removeDuplicates(source.nums)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}

}
