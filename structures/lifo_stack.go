package structures

import "errors"

type LifoStack interface {
	Pop() (interface{}, error)
	Push(element interface{})
	GetLength() int
}

type stack struct {
	elements []interface{}
}

func NewLifoStack() LifoStack {
	return &stack{
		elements: make([]interface{}, 0, 0),
	}
}

func (s *stack) Pop() (interface{}, error) {
	len := len(s.elements)
	if len == 0 {
		return 0, errors.New("The stack is empty!")
	}

	res := s.elements[len-1]
	s.elements = s.elements[:len-1]
	return res, nil
}

func (s *stack) Push(element interface{}) {
	s.elements = append(s.elements, element)
}

func (s *stack) GetLength() int {
	return len(s.elements)
}
