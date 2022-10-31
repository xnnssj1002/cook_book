package bfs

import (
	"fmt"
	"testing"
)

var root = Node{
	Val: 2,
	Left: &Node{
		Val: 1,
		Left: &Node{
			Val: 3,
			Left: &Node{
				Val: 8,
			},
			Right: &Node{
				Val: 80,
			},
		},
		Right: &Node{
			Val: 4,
			Left: &Node{
				Val: 90,
				Left: &Node{
					Val: 77,
				},
				Right: &Node{
					Val: 89,
				},
			},
			Right: &Node{
				Val: 33,
				Left: &Node{
					Val: 41,
				},
				Right: &Node{
					Val: 46,
				},
			},
		},
	},
	Right: &Node{
		Val: 10,
		Left: &Node{
			Val: 31,
			Left: &Node{
				Val: 82,
			},
		},
		Right: &Node{
			Val: 43,
			Left: &Node{
				Val: 904,
				Left: &Node{
					Val: 775,
				},
				Right: &Node{
					Val: 896,
				},
			},
			Right: &Node{
				Val: 337,
				Left: &Node{
					Val: 418,
				},
				Right: &Node{
					Val: 469,
				},
			},
		},
	},
}

func TestLoopTraversal(t *testing.T) {
	res := LoopTraversal(&root)
	fmt.Println(res)
}
