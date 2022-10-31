package dfs

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// RecursionTraversal 深度优先遍历 - 递归实现
func RecursionTraversal(root *Node) (ResSli []int) {
	var preorder func(*Node)
	preorder = func(node *Node) {
		if node == nil {
			return
		}
		ResSli = append(ResSli, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

// LoopTraversal 深度优先遍历 - 非递归循环实现
// 		A
// 	   / \
//	  B   C
// 栈的初始元素为根节点，循环条件为栈的长度大于0
// 最简单的结构，开始将根节点推入栈中，在循环中，从栈中取出最后一个元素，在这里就是根节点，放入到返回切片中
// 然后将右节点入栈，接着将左节点入栈。之所以先右后左，是因为在循环体中要先处理左节点，即先弹出左节点的元素处理
// 如果左右节点还有子节点就继续这样的操作，即在入栈时判断，只有有子节点就按照先右后左入栈
func LoopTraversal(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	var stack []*Node
	stack = append(stack, root)
	for len(stack) > 0 {
		size := len(stack)
		temp := stack[size-1]
		if size <= 1 {
			stack = []*Node{}
		} else {
			stack = stack[:size-1]
		}
		res = append(res, temp.Val)
		if temp.Right != nil {
			stack = append(stack, temp.Right)
		}
		if temp.Left != nil {
			stack = append(stack, temp.Left)
		}
	}
	return res
}

func Depth(root *Node) []int {
	var resSli []int
	var stack []*Node
	// 先把根节点入栈
	stack = append(stack, root)
	for len(stack) > 0 {
		// 从栈中取出一个元素
		temp := stack[len(stack)-1]
		// 将弹出的元素值放入返回值中
		resSli = append(resSli, temp.Val)
		// 弹出元素后的栈
		stack = stack[:len(stack)-1] // 从原栈上去掉最后一个元素
		// 先处理右节点
		if temp.Right != nil {
			stack = append(stack, temp.Right)
		}
		if temp.Left != nil {
			stack = append(stack, temp.Left)
		}
	}
	return resSli
}
