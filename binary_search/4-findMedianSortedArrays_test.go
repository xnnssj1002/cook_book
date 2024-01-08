package binary_search

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	sources := []struct {
		name  string
		nums1 []int
		nums2 []int
		hope  float64
	}{
		{
			name:  "test0",
			nums1: []int{1, 3},
			nums2: []int{2},
			hope:  2,
		},
		{
			name:  "test1",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			hope:  2.5,
		},
	}
	for _, source := range sources {
		r := findMedianSortedArrays1(source.nums1, source.nums2)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %f, got %f\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}
