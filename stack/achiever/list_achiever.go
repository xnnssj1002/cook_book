package achiever

// 通过链表实现的链式栈

// node 节点
type node struct {
	prev, next *node
	value      string
}

// ListStack 链表链式栈
type ListStack struct {
	root *node
	len  int
}

// NewListStack NewListStack
func NewListStack() *ListStack {
	n := &node{}
	n.next = n
	n.prev = n
	return &ListStack{root: n}
}

// Push 入栈
func (s *ListStack) Push(item string) {
	n := &node{value: item}
	s.root.prev.next = n
	n.prev = s.root.prev
	n.next = s.root
	s.root.prev = n
	s.len++
}

// Pop 出栈
func (s *ListStack) Pop() string {
	item := s.root.prev
	if item == s.root {
		return ""
	}

	s.root.prev = item.prev
	item.prev.next = s.root
	// 避免内存泄漏
	item.prev = nil
	item.next = nil
	s.len--
	return item.value
}
