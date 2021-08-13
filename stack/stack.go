package stack

import (
	"errors"
	"sync"
)

type Stack interface {
	Pop() (interface{}, error)
	Push(elementList ...interface{})
	GetLength() int
}

type stack struct {
	elements []interface{}
	lock     sync.RWMutex
}

func New() Stack {
	return &stack{
		elements: make([]interface{}, 0),
	}
}

func (s *stack) Pop() (interface{}, error) {
	len := len(s.elements)
	if len == 0 {
		return 0, errors.New("The stack is empty!")
	}

	s.lock.Lock()
	res := s.elements[len-1]
	s.elements = s.elements[:len-1]
	s.lock.Unlock()

	return res, nil
}

func (s *stack) Push(elementList ...interface{}) {
	s.lock.Lock()
	s.elements = append(s.elements, elementList...)
	s.lock.Unlock()
}

func (s *stack) GetLength() int {
	return len(s.elements)
}
