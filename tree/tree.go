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

type Node struct {
	Val   int
	Left  *Node
	Right *Node
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
