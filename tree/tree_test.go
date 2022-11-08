package tree

import (
	"fmt"
	"testing"
)

/*
				  2
			/           \
		   1            17
		/    \         /   \
	   3	  4       31    43
	  / \    /  \    /     /  \
	 8  80  90  33  82    91  35

*/

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
			},
			Right: &Node{
				Val: 33,
			},
		},
	},
	Right: &Node{
		Val: 17,
		Left: &Node{
			Val: 31,
			Left: &Node{
				Val: 82,
			},
		},
		Right: &Node{
			Val: 43,
			Left: &Node{
				Val: 91,
			},
			Right: &Node{
				Val: 35,
			},
		},
	},
}

// 递归深度优先-前序遍历
func TestNode_PreorderDfsTraversalByRecursion(t *testing.T) {
	res := root.PreorderDfsTraversalByRecursion()
	fmt.Println("前序遍历-递归", res)
}

// 非递归(栈实现)深度优先 - 前序遍历
func TestNode_PreorderDfsTraversalByStack(t *testing.T) {
	res := root.PreorderDfsTraversalByIterateInStack()
	fmt.Println("前序遍历-栈", res)
}

// 递归深度优先 - 中序遍历
func TestNode_MiddleOrderDfsTraversalByRecursion(t *testing.T) {
	res := root.MiddleOrderDfsTraversalByRecursion()
	fmt.Println("中序遍历-递归", res)
}

// 非递归(栈实现)深度优先 - 中序遍历
func TestNode_MiddleOrderDfsTraversalByIterateInStack(t *testing.T) {
	res := root.MiddleOrderDfsTraversalByIterateInStack()
	fmt.Println("中序遍历-栈", res)
}

// 递归深度优先 - 后序遍历
func TestNode_PostorderDfsTraversalByRecursion(t *testing.T) {
	res := root.PostorderDfsTraversalByRecursion()
	fmt.Println("后序遍历-递归", res)
}

// 反转前序遍历 - 后序遍历
func TestNode_PostorderDfsReversePreorder(t *testing.T) {
	res := root.PostorderDfsReversePreorder()
	fmt.Println("后序遍历-反转前序遍历", res)
}

// 非递归(栈实现)深度优先 - 后序遍历
func TestNode_PostorderDfsTraversalByIterateInStack(t *testing.T) {
	res := root.PostorderDfsTraversalByIterateInStack()
	fmt.Println("后序遍历-栈", res)
}

// 广度优先遍历
func TestNode_BfsTraversal(t *testing.T) {
	res := root.BfsTraversal()
	fmt.Println("广度优先", res)
}
