package tree

/*
	借助 栈(非递归的方法) 完成前中后序遍历二叉树
	前中后三种写法不统一
*/

// ------------- 前序遍历 ------------- //
// 根 -----> 左子树 -----> 右子树

// PreorderDfsTraversalByIterateInStack 深度优先-前序遍历 之 非递归实现，借助 stack 栈使用迭代实现
func (n *Node) PreorderDfsTraversalByIterateInStack() []int {
	resSli := make([]int, 0)

	stack := make([]*Node, 0)
	stack = append(stack, n)

	for len(stack) > 0 {
		temp := stack[len(stack)-1]
		resSli = append(resSli, temp.Val)
		// 弹出最后一个元素
		if len(stack) <= 1 {
			stack = []*Node{}
		} else {
			stack = stack[:len(stack)-1]
		}

		// 先放右节点，再放左节点
		if temp.Right != nil {
			stack = append(stack, temp.Right)
		}
		if temp.Left != nil {
			stack = append(stack, temp.Left)
		}
	}

	return resSli
}

// ------------- 中序遍历 ------------- //
// 左子树 -----> 根 ------> 右子树

// MiddleOrderDfsTraversalByIterateInStack 深度优先-中序遍历 之 非递归实现，借助 stack 栈使用迭代实现
// 需要一个当前结点。
// 1、将传入结点作为当前结点
// 2、如果当前结点不为nil的话，让当前结点入栈，同时将当前结点的左结点作为当前结点，进行循环
// 3、如果当前结点为nil的话，从栈中弹出一个元素，将弹出元素的值放入返回切片中，同时将当前结点的右结点作为当前结点，进行循环。
func (n *Node) MiddleOrderDfsTraversalByIterateInStack() []int {
	resSli := make([]int, 0)
	stack := make([]*Node, 0)
	root := n
	for root != nil || len(stack) > 0 {
		// 只要左结点不为nil，就一直将当前结点入栈
		for root != nil {
			stack = append(stack, root)
			// 设置 左结点 为当前结点
			root = root.Left
		}
		// 从栈中弹出一个元素
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 将弹出的元素放入返回切片中
		resSli = append(resSli, root.Val)
		// 设置 右结点 为当前结点
		root = root.Right
	}
	return resSli
}

// ------------- 后序遍历 ------------- //
// 左子树 -----> 右子树 -----> 根

// PostorderDfsReversePreorder 深度优先-后序遍历 使用【前序遍历逻辑】结果反转 获取 后序遍历结果
// 前序：中->左->右；相同逻辑中->右->左；反转左->右->中
func (n *Node) PostorderDfsReversePreorder() []int {
	resSli := make([]int, 0)
	stack := make([]*Node, 0)
	stack = append(stack, n)
	for len(stack) > 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		resSli = append(resSli, temp.Val)
		if temp.Left != nil {
			stack = append(stack, temp.Left)
		}
		if temp.Right != nil {
			stack = append(stack, temp.Right)
		}
	}
	// 反转resSli
	i, j := 0, len(resSli)-1
	for i <= j {
		resSli[i], resSli[j] = resSli[j], resSli[i]
		i++
		j--
	}
	return resSli
}

// PostorderDfsTraversalByIterateInStack 深度优先-后序遍历 之 非递归实现，借助 stack 栈使用迭代实现
func (n *Node) PostorderDfsTraversalByIterateInStack() []int {
	resSli := make([]int, 0)
	stack := make([]*Node, 0)
	var prev *Node
	root := n
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			resSli = append(resSli, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return resSli
}
