package stack

import "errors"

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{
		items: []int{},
	}
}

func (s *Stack) Push(value int) {
	s.items = append(s.items, value)
}

func(s *Stack) Peek() (int, error) {
	if len(s.items)<1 {
		return 0,errors.New("the stack is empty")
	}
	return s.items[len(s.items)-1], nil
}


// func(s *Stack) Pop() (int, error) {
// 	if len(s.items)<1 {
// 		return 0,errors.New("the stack is empty")
// 	}
// 	itemsCopy := make([]int, len(s.items)-1)
// 	copy(itemsCopy, s.items)
// 	s.items=itemsCopy
// 	return itemsCopy[len(itemsCopy)-1], nil
// }



//or 

func (s *Stack) Pop() (int, error) {
	if len(s.items) < 1 {
		return 0, errors.New("the stack is empty")
	}

	lastIndex := len(s.items) - 1
	value := s.items[lastIndex]

	s.items = s.items[:lastIndex]

	return value, nil
}