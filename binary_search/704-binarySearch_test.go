package binary_search

import "testing"

var binarySearchSource = []struct {
	name   string
	nums   []int
	target int
	hope   int
}{
	{
		name:   "test0",
		nums:   []int{0, 1, 2, 4, 5, 6, 7},
		target: 5,
		hope:   4,
	},
	{
		name:   "test1",
		nums:   []int{0, 1, 2, 4, 5, 6, 7},
		target: 7,
		hope:   6,
	},
	{
		name:   "test2",
		nums:   []int{1},
		target: 0,
		hope:   -1,
	},
}

func TestBinarySearch(t *testing.T) {
	for _, source := range binarySearchSource {
		r := binarySearch(source.nums, source.target)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %d, got %d\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}

func TestBinarySearchRemoveLeft(t *testing.T) {
	for _, source := range binarySearchSource {
		r := binarySearchRemoveLeft(source.nums, source.target)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %d, got %d\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}

func TestBinarySearchRemoveRight(t *testing.T) {
	for _, source := range binarySearchSource {
		r := binarySearchRemoveRight(source.nums, source.target)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %d, got %d\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}
