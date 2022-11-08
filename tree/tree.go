package tree

import "container/list"

/*
遍历二叉树有两种模式：
1：深度优先
2：广度优先

深度优先有三种方法：
1、前序遍历：前序遍历的顺序是  根 -----> 左子树 -----> 右子树
2、中序遍历：中序遍历的顺序是 左子树 -----> 根 ------> 右子树
3、后序遍历：后序遍历的顺序是 左子树 -----> 右子树 -----> 根

广度优先即优先按层遍历，又叫层序遍历，顺序为：根(0 level) -----> 从左到右(1 level) -----> 从左到右(2 level)
*/

/*
专业术语之------迭代，循环，遍历，递归的区别？
　　loop、iterate、traversal和recursion这几个词是计算机技术书中经常会出现的几个词汇。众所周知，这几个词分别翻译 为：循环、迭代、遍历和递归。乍一看，这几个词好像都与重复（repeat）有关，但有的又好像不完全是重复的意思。那么这几个词到底各是什么含义，有什么区别和联系呢？下面就试着解释一下。
　　1，循环（loop），指的是在满足条件的情况下，重复执行同一段代码。比如，while语句。   循环则技能对应集合，列表，数组等，也能对执行代码进行操作。
　　2，迭代（iterate），指的是按照某种顺序逐个访问列表中的每一项。比如，for语句。  　　迭代只能对应集合，列表，数组等。不能对执行代码进行迭代。
　　3，遍历（traversal），指的是按照一定的规则访问树形结构中的每个节点，而且每个节点都只访问一次。   遍历同迭代一样，也不能对执行代码进行遍历。
　　4，递归（recursion），指的是一个函数不断调用自身的行为。比如，以编程方式输出著名的斐波纳契数列
　　　　(1)，通俗解释：递归就像往存钱罐里存钱，先往里边塞钱，2块，5块，10块这样的塞，叫入栈。取钱的时候，后塞进去的先取出来，这叫出栈。具体多少钱，要全部出栈才知道。
　　　　(2)，递归分类：线性递归和尾递归。
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

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

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

// ------------- 广度优先遍历 ------------- //

// BfsTraversal 广度优先遍历
func (n *Node) BfsTraversal() []int {
	resSli := make([]int, 0)
	// 创建一个队列 queue
	queue := list.New()
	queue.PushBack(n)
	for queue.Len() > 0 {
		temp := queue.Front()
		queue.Remove(temp)
		node := temp.Value.(*Node)
		resSli = append(resSli, node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return resSli
}
