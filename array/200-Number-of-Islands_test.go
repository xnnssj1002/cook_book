package array

import (
	"fmt"
	"testing"
)

var grid = [][]byte{
	[]byte{'1', '1', '0', '0', '0'},
	[]byte{'1', '1', '0', '0', '0'},
	[]byte{'0', '0', '1', '0', '0'},
	[]byte{'0', '0', '0', '1', '1'},
}

func TestNumIslandsByDfs(t *testing.T) {
	res := NumIslandsByDfs(grid)
	fmt.Println(res)
}

func TestNumIslandsByDfsSelf(t *testing.T) {
	res := NumIslandsByDfsSelf(grid)
	fmt.Println(res)
}

func TestNumIslandsBfs(t *testing.T) {
	res := NumIslandsBfs(grid)
	fmt.Println(res)
}

func TestNumIslandsByBfsSelf(t *testing.T) {
	res := NumIslandsByBfsSelf(grid)
	fmt.Println(res)
}
