package main

import "errors"

type StackStruct struct {
	elements []any
}

func CreateStack() *StackStruct {
	return &StackStruct{
		elements: make([]any, 0),
	}
}

func (s *StackStruct) Add(v any) {
	s.elements = append(s.elements, v)
}

func (s *StackStruct) Remove() (any, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("в стеке нет элементов")
	}
	val := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return val, nil
}

func (s *StackStruct) Top() (any, error) {
	if len(s.elements) == 0 {
		return nil, errors.New("в стеке нет элементов")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *StackStruct) Empty() bool {
	return len(s.elements) == 0
}
