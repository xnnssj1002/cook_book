package sort

import "testing"

var sortSources = []struct {
	name string
	nums []int
	hope []int
}{
	{
		name: "test0",
		nums: []int{2, 1, 5, 3, 4, 7, 6, 9, 8},
		hope: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		name: "test1",
		nums: []int{3, 6, 2},
		hope: []int{2, 3, 6},
	},
}

// 冒泡排序
func TestBubbleSort(t *testing.T) {
	for _, source := range sortSources {
		r := bubbleSort(source.nums)
		if !towSliceStrictEquation(source.hope, r) {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}

// 选择排序
func TestSelectionSort(t *testing.T) {
	for _, source := range sortSources {
		r := selectionSort(source.nums)
		if !towSliceStrictEquation(source.hope, r) {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}

// 插入排序
func TestInsertionSort(t *testing.T) {
	for _, source := range sortSources {
		r := insertionSort(source.nums)
		if !towSliceStrictEquation(source.hope, r) {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}

// 快速排序
func TestQuickSort(t *testing.T) {
	for _, source := range sortSources {
		r := quickSort(source.nums)
		if !towSliceStrictEquation(source.hope, r) {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)
		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}
}

// 判断两个切片是否相等
func towSliceStrictEquation(num1 []int, num2 []int) bool {
	if len(num1) != len(num2) {
		return false
	}
	for i := 0; i < len(num1); i++ {
		if num1[i] != num2[i] {
			return false
		}
	}
	return true
}
