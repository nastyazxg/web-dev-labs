package main

import "errors"

type QueueStruct struct {
	elements []any
}

func CreateQueue() *QueueStruct {
	return &QueueStruct{
		elements: make([]any, 0),
	}
}

func (q *QueueStruct) Add(v any) {
	q.elements = append(q.elements, v)
}

func (q *QueueStruct) Remove() (any, error) {
	if len(q.elements) == 0 {
		return nil, errors.New("в очереди нет элементов")
	}
	val := q.elements[0]
	q.elements = q.elements[1:]
	return val, nil
}

func (q *QueueStruct) First() (any, error) {
	if len(q.elements) == 0 {
		return nil, errors.New("в очереди нет элементов")
	}
	return q.elements[0], nil
}

func (q *QueueStruct) Empty() bool {
	return len(q.elements) == 0
}
