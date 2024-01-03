package radix_tree

import "strings"

type radixNode struct {
	// 当前节点的相对路径
	path string
	// 完整路径
	fullPath string
	// 每个 indice（标记） 字符对应一个孩子节点的 path 首字母
	indices string
	// 后继节点
	children []*radixNode
	// 是否有路径以当前节点为终点
	end bool
	// 记录有多少路径途径当前节点
	passCnt int
}

type Radix struct {
	root *radixNode
}

func NewRadix() *Radix {
	return &Radix{
		root: &radixNode{},
	}
}

// StartWith 判断是否以某种字符作为前缀的核心逻辑
// • 我们通过调用根节点 root 的 search 方法，检索出可能包含 prefix 为前缀的节点 node
// • 如果对应节点存在，并且其全路径 fullPath 确实以 prefix 为前缀，则前缀匹配成功
func (r *Radix) StartWith(prefix string) bool {
	node := r.root.search(prefix)
	return node != nil && strings.HasPrefix(node.fullPath, prefix)
}

// PassCnt 统计的核心逻辑
// • 我们通过调用根节点 root 的 search 方法，检索出可能包含以 prefix 为前缀的节点 node
// • 如果对应节点存在并且其全路径 fullPath 确实以 prefix 为前缀，则返回该节点 passCnt 计数器的值
func (r *Radix) PassCnt(prefix string) int {
	node := r.root.search(prefix)
	if node == nil || !strings.HasPrefix(node.fullPath, prefix) {
		return 0
	}
	return node.passCnt
}
