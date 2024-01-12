package sliding_window

import "testing"

var numOfSubArraysSources = []struct {
	name      string
	arr       []int
	k         int
	threshold int
	hope      int
}{
	{
		name:      "test0",
		arr:       []int{2, 2, 2, 2, 5, 5, 5, 8},
		k:         3,
		threshold: 4,
		hope:      3,
	},
	{
		name:      "test1",
		arr:       []int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3},
		k:         3,
		threshold: 5,
		hope:      6,
	},
}

func TestIntersection1(t *testing.T) {

	for _, source := range numOfSubArraysSources {
		r := numOfSubArrays1(source.arr, source.k, source.threshold)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)

		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}

}

func TestIntersection2(t *testing.T) {

	for _, source := range numOfSubArraysSources {
		r := numOfSubArrays2(source.arr, source.k, source.threshold)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)

		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}

func TestIntersection3(t *testing.T) {

	for _, source := range numOfSubArraysSources {
		r := numOfSubArrays3(source.arr, source.k, source.threshold)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)

		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}
