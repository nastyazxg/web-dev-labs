package main

import "errors"

type Element struct {
	Val  any
	Next *Element
}

type List struct {
	First *Element
	Last  *Element
	Count int
}

func NewList() *List {
	return &List{
		First: nil,
		Last:  nil,
		Count: 0,
	}
}

func (l *List) Append(v any) {
	newElement := &Element{Val: v, Next: nil}

	if l.First == nil {
		l.First = newElement
		l.Last = newElement
	} else {
		l.Last.Next = newElement
		l.Last = newElement
	}
	l.Count++
}

func (l *List) ItemAt(pos int) (any, error) {
	if pos < 0 || pos >= l.Count {
		return nil, errors.New("индекс за пределами списка")
	}

	curr := l.First
	for i := 0; i < pos; i++ {
		curr = curr.Next
	}
	return curr.Val, nil
}

func (l *List) DeleteAt(pos int) error {
	if pos < 0 || pos >= l.Count {
		return errors.New("индекс за пределами списка")
	}

	if pos == 0 {
		l.First = l.First.Next
		if l.First == nil {
			l.Last = nil
		}
	} else {
		curr := l.First
		for i := 0; i < pos-1; i++ {
			curr = curr.Next
		}
		curr.Next = curr.Next.Next
		if curr.Next == nil {
			l.Last = curr
		}
	}
	l.Count--
	return nil
}

func (l *List) GetAll() []any {
	res := make([]any, l.Count)
	curr := l.First
	for i := 0; i < l.Count; i++ {
		res[i] = curr.Val
		curr = curr.Next
	}
	return res
}
