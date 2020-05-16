package stack

// Stack hold a list of items in a lifo stack
type Stack struct {
	items []int
}

// Push an item onto the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop an item off the stack
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}
