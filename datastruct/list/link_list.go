package list

import (
	"constraints"
	"errors"
)

type LinkList[T constraints.Ordered] struct {
	head *node[T]
	len  int
}

func (item *LinkList[T]) Front() T {
	if item.len == 0 {
		panic(errors.New("out of range"))
	}
	return item.head.data
}

func (item *LinkList[T]) Add(index int, value T) {
	if index < 0 || index > item.len {
		panic(errors.New("out of range"))
	}
	if index == 0 {
		item.AddFront(value)
		return
	}
	pre := item.head
	for i := 0; i < index; i++ {
		pre = pre.next
	}
	pre.next = &node[T]{
		data: value,
		next: pre.next,
	}
	item.len++
}

func (item *LinkList[T]) AddFront(value T) {
	item.head = &node[T]{
		data: value,
		next: item.head,
	}
}

func (item *LinkList[T]) AddLast(value T) {
	item.Add(item.len, value)
}

func (item *LinkList[T]) Delete(index int) {
	//TODO implement me
	panic("implement me")
}

func (item *LinkList[T]) Get(index int) T {
	//TODO implement me
	panic("implement me")
}

func (item *LinkList[T]) Len() int {
	return item.len
}

type node[T constraints.Ordered] struct {
	data T
	next *node[T]
}

func NewLinkList[T constraints.Ordered]() *LinkList[T] {
	return *LinkList[T]{}
}
