package two_points

import "testing"

var intersectionSources = []struct {
	name  string
	nums1 []int
	nums2 []int
	hope  []int
}{
	{
		name:  "test0",
		nums1: []int{1, 2, 2, 1},
		nums2: []int{2, 2},
		hope:  []int{2},
	},
	{
		name:  "test1",
		nums1: []int{4, 9, 5},
		nums2: []int{9, 4, 9, 8, 4},
		hope:  []int{9, 4},
	},
}

func TestIntersection1(t *testing.T) {

	for _, source := range intersectionSources {
		r := intersection1(source.nums1, source.nums2)
		if !towSliceEquation(source.hope, r) {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}

}

func TestIntersection2(t *testing.T) {

	for _, source := range intersectionSources {
		r := intersection2(source.nums1, source.nums2)
		if !towSliceEquation(source.hope, r) {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}

}

func TestIntersection2Optimize(t *testing.T) {

	for _, source := range intersectionSources {
		r := intersection2Optimize(source.nums1, source.nums2)
		if !towSliceEquation(source.hope, r) {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}

}

// 判断两个切片是否相等
func towSliceEquation(num1 []int, num2 []int) bool {
	if len(num1) != len(num2) {
		return false
	}
	for i := 0; i < len(num1); i++ {
		for k := 0; k < len(num2); k++ {
			if k == len(num2)-1 && num2[k] != num1[i] {
				return false
			}
			if num2[k] == num1[i] {
				break
			}
		}
	}
	return true
}
