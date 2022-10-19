package array

import "testing"

func TestMaxAreaLoop(t *testing.T) {
	res := MaxAreaLoop([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if res != 49 {
		t.Fatalf("expected %d, but got %d\n", 49, res)
	}
}

func TestMaxAreaCode(t *testing.T) {
	res := MaxAreaCode([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if res != 49 {
		t.Fatalf("expected %d, but got %d\n", 49, res)
	}
}
