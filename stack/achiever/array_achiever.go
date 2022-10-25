package achiever

// 通过数组实现的顺序栈
// 非并发安全

// ArrayStack 数组顺序栈
type ArrayStack struct {
	items   []string
	current int
}

// NewArrayStack 创建一个数组顺序栈
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		items:   make([]string, 10),
		current: 0,
	}
}

// Push 入栈
func (s *ArrayStack) Push(item string) {
	s.current++
	// 判断底层 slice 是否满了，如果满了就 append
	if s.current == len(s.items) {
		s.items = append(s.items, item)
		return
	}
	s.items[s.current] = item
}

// Pop 出栈
func (s *ArrayStack) Pop() string {
	if s.current == 0 {
		return ""
	}
	item := s.items[s.current]
	s.current--
	return item
}
