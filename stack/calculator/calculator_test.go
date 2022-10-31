package calculator

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	res := Calculate("1+2*3-1/1")
	fmt.Println(res)
}
