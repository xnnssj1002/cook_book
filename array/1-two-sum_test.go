package array

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		Name     string
		Nums     []int
		Target   int
		Expected []int
	}{
		{
			Name:     "exists-one",
			Nums:     []int{2, 7, 11, 15},
			Target:   9,
			Expected: []int{0, 1},
		},
		{
			Name:     "exists-two",
			Nums:     []int{2, 6, 5, 3},
			Target:   8,
			Expected: []int{0, 1},
		},
		{
			Name:     "non-exists",
			Nums:     []int{2, 7, 11, 15},
			Target:   8,
			Expected: nil,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			//IndexSli := TwoSumSelf(c.Nums, c.Target)
			IndexSli := TwoSumBook(c.Nums, c.Target)
			fmt.Println(c.Name, IndexSli)

			if c.Expected == nil && IndexSli != nil {
				t.Fatalf("expected %v, but %v got", c.Expected, IndexSli)
			}

			if c.Expected != nil && IndexSli == nil {
				t.Fatalf("expected %v, but %v got", c.Expected, IndexSli)
			}

			expectedIndexSum := 0
			for i := 0; i < len(c.Expected); i++ {
				expectedIndexSum += c.Expected[i]
			}
			returnIndexSum := 0
			for i := 0; i < len(IndexSli); i++ {
				returnIndexSum += IndexSli[i]
			}
			if expectedIndexSum != returnIndexSum {
				t.Fatalf("expected %v, but %v got", c.Expected, IndexSli)
			}

		})
	}
}
