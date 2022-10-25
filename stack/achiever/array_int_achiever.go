package achiever

// 通过数组实现的顺序栈 元素为 int 型
// 非并发安全

// ArrayIntStack 数组顺序栈
type ArrayIntStack struct {
	items   []int
	current int
}

// NewArrayIntStack 创建一个数组顺序栈
func NewArrayIntStack() *ArrayIntStack {
	return &ArrayIntStack{
		items:   make([]int, 10),
		current: 0,
	}
}

// Push 入栈
func (s *ArrayIntStack) Push(item int) {
	s.current++
	// 判断底层 slice 是否满了，如果满了就 append
	if s.current == len(s.items) {
		s.items = append(s.items, item)
		return
	}
	s.items[s.current] = item
}

// Pop 出栈
func (s *ArrayIntStack) Pop() int {
	if s.current == 0 {
		return 0
	}
	item := s.items[s.current]
	s.current--
	return item
}
