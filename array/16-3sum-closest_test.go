package array

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	res := ThreeSumClosest([]int{-1, 2, 1, -4}, 1)
	if res != 2 {
		t.Fatalf("expected %d, but %d got", 2, res)
	}
	res1 := ThreeSumClosest([]int{0, 0, 0}, 1)
	if res1 != 0 {
		t.Fatalf("expected %d, but %d got", 0, res1)
	}
}

func TestThreeSumClosestBy15Question(t *testing.T) {
	res := ThreeSumClosestBy15Question([]int{-1, 2, 1, -4}, 1)
	if res != 2 {
		t.Fatalf("expected %d, but %d got", 2, res)
	}
	res1 := ThreeSumClosestBy15Question([]int{0, 0, 0}, 1)
	if res1 != 0 {
		t.Fatalf("expected %d, but %d got", 0, res1)
	}
}
