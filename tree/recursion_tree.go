package tree

/*
	使用 递归 的方法完成前中后序遍历二叉树
*/

/*
每次写递归，需要确定三要素，可以保证写出正确地递归算法！
1、确定递归函数的参数和返回值：
确定哪些参数是递归的过程中需要处理的，那么就在递归函数里加上这个参数， 并且还要明确每次递归的返回值是什么进而确定递归函数的返回类型。
2、确定终止条件：
写完了递归算法, 运行的时候，经常会遇到栈溢出的错误，就是没写终止条件或者终止条件写的不对，操作系统也是用一个栈的结构来保存每一层递归的信息，如果递归没有终止，操作系统的内存栈必然就会溢出。
3、确定单层递归的逻辑：
确定每一层递归需要处理的信息。在这里也就会重复调用自己来实现递归的过程。
*/

// ------------- 前序遍历 ------------- //
// 根 -----> 左子树 -----> 右子树

// PreorderDfsTraversalByRecursion 深度优先-前序遍历 之 递归实现
func (n *Node) PreorderDfsTraversalByRecursion() []int {
	resSli := make([]int, 0)

	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		// 递归顺序：先处理根结点，再处理左子树，最后处理右子树
		resSli = append(resSli, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(n)

	return resSli
}

// ------------- 中序遍历 ------------- //
// 左子树 -----> 根 ------> 右子树

// MiddleOrderDfsTraversalByRecursion 深度优先-中序遍历 之 递归实现
func (n *Node) MiddleOrderDfsTraversalByRecursion() []int {
	resSli := make([]int, 0)

	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		// 递归顺序：先处理左子树，再处理根结点，最后处理右子树
		dfs(node.Left)
		resSli = append(resSli, node.Val)
		dfs(node.Right)
	}
	dfs(n)

	return resSli
}

// ------------- 后序遍历 ------------- //
// 左子树 -----> 右子树 -----> 根

// PostorderDfsTraversalByRecursion 深度优先-后序遍历 之 递归实现
func (n *Node) PostorderDfsTraversalByRecursion() []int {
	resSli := make([]int, 0)

	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		// 递归顺序：先处理左子树，再处理右子树，最后处理根结点
		dfs(node.Left)
		dfs(node.Right)
		resSli = append(resSli, node.Val)
	}
	dfs(n)

	return resSli
}
