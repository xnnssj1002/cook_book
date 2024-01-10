package two_points

import "testing"

var maxAreaSources = []struct {
	name string
	nums []int
	hope int
}{
	{
		name: "test0",
		nums: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		hope: 49,
	},
	{
		name: "test1",
		nums: []int{1, 1},
		hope: 1,
	},
}

func TestMaxArea(t *testing.T) {

	for _, source := range maxAreaSources {
		r := maxArea(source.nums)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}

}
