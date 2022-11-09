package tree

/*
	使用 栈+标记法，实现前中后遍历统一写法
*/

/*
标记法的核心思想为：
1、使用颜色标记结点的状态，新结点为白色，已访问的结点为灰色。
2、出栈时，如果遇到的结点为白色，则将其标记为灰色，然后将其右子结点、自身、左子结点依次入栈。
3、出栈时，如果遇到的结点为灰色，则将结点的值输出。
*/

/* 以 中序遍历 为例，解释一下入栈顺序
解释一下为什么需要 “右、中、左” 依次入栈

我们有一棵二叉树：

               中
              /  \
             左   右
前序遍历：中，左，右
中序遍历：左，中，右
后序遍历：左，右，中

以 中序遍历 为例：
栈是一种 先进后出的结构，出栈顺序为左，中，右
那么入栈顺序必须调整为倒序，也就是右，中，左

同理，如果是前序遍历，入栈顺序为 右，左，中；
后序遍历，入栈顺序中，右，左
*/

type color int

const (
	WHITE color = 0
	GRAY  color = 1
)

type signNode struct {
	node  *Node
	color color
}

// ------------- 前序遍历 ------------- //
// 根 -----> 左子树 -----> 右子树

// PreorderDfsTraversalByStackSign 深度优先-前序遍历 之 栈+标记
func (n *Node) PreorderDfsTraversalByStackSign() []int {
	resSli := make([]int, 0)

	stack := make([]signNode, 0)
	stack = append(stack, signNode{node: n, color: WHITE})

	for len(stack) > 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if temp.node == nil {
			continue
		}
		if temp.color == WHITE {
			stack = append(stack, signNode{node: temp.node.Right, color: WHITE})
			stack = append(stack, signNode{node: temp.node.Left, color: WHITE})
			stack = append(stack, signNode{node: temp.node, color: GRAY})
		}
		if temp.color == GRAY {
			resSli = append(resSli, temp.node.Val)
		}
	}

	return resSli
}

// ------------- 中序遍历 ------------- //
// 左子树 -----> 根 ------> 右子树

// MiddleOrderDfsTraversalByStackSign 深度优先-中序遍历 之 栈+标记
func (n *Node) MiddleOrderDfsTraversalByStackSign() []int {
	resSli := make([]int, 0)

	stack := make([]signNode, 0)
	stack = append(stack, signNode{node: n, color: WHITE})

	for len(stack) > 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if temp.node == nil {
			continue
		}
		if temp.color == WHITE {
			stack = append(stack, signNode{node: temp.node.Right, color: WHITE})
			stack = append(stack, signNode{node: temp.node, color: GRAY})
			stack = append(stack, signNode{node: temp.node.Left, color: WHITE})
		}
		if temp.color == GRAY {
			resSli = append(resSli, temp.node.Val)
		}
	}

	return resSli
}

// ------------- 后序遍历 ------------- //
// 左子树 -----> 右子树 -----> 根

// PostorderDfsTraversalByStackSign 深度优先-后序遍历 之 栈+标记
func (n *Node) PostorderDfsTraversalByStackSign() []int {
	resSli := make([]int, 0)

	stack := make([]signNode, 0)
	stack = append(stack, signNode{node: n, color: WHITE})

	for len(stack) > 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if temp.node == nil {
			continue
		}
		if temp.color == WHITE {
			stack = append(stack, signNode{node: temp.node, color: GRAY})
			stack = append(stack, signNode{node: temp.node.Right, color: WHITE})
			stack = append(stack, signNode{node: temp.node.Left, color: WHITE})
		}
		if temp.color == GRAY {
			resSli = append(resSli, temp.node.Val)
		}
	}

	return resSli
}
