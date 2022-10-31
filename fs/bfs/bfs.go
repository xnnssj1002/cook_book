package bfs

import "container/list"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// LoopTraversal 使用 队列 数据结构，完成广度优先遍历
// 一个节点弹出队列的同时，要分别将该节点对应的左节点、右节点 先后入队列
// 从队列头部再弹出一个元素，重复上边的操作
func LoopTraversal(root *Node) []int {
	resSli := make([]int, 0)
	if root == nil {
		return resSli
	}
	// 创建一个队列
	queue := list.New()
	// 将根节点加入队列
	queue.PushBack(root)
	for queue.Len() > 0 {
		// 从队列头部弹出一个元素，并获取元素中的 element
		temp := queue.Remove(queue.Front()).(*Node)
		resSli = append(resSli, temp.Val)
		// 将左节点入队列
		if temp.Left != nil {
			queue.PushBack(temp.Left)
		}
		// 将右节点入队列
		if temp.Right != nil {
			queue.PushBack(temp.Right)
		}
	}
	return resSli
}
