package list

import "errors"

type LinkList struct {
	head *node
	len  int
}

func (item *LinkList) Front() (int, error) {
	if item.head == nil {
		return 0, errors.New("not found")
	}
	return item.head.value, nil
}

func (item *LinkList) Get(index int) (int, error) {
	start := item.head.next
	for i := 0; i < index; i++ {
		if start != nil {
			start = start.next
		} else {
			return 0, errors.New("out of range")
		}
	}
	return start.value, nil
}

func (item *LinkList) Push(values ...int) {
	if len(values) == 0 {
		return
	}
	lastNode := item.head
	for {
		if lastNode.next == nil {
			break
		}
		lastNode = lastNode.next
	}
	item.push(lastNode, values...)
}

func (item *LinkList) push(n *node, values ...int) {
	for _, v := range values {
		n.next = &node{
			value: v,
			next:  nil,
		}
		n = n.next
		item.len++
	}
}

func (item *LinkList) PushFont(value int) {
	node := &node{
		value: value,
		next:  item.head,
	}
	item.head.next = node
}

func (item *LinkList) Delete(index int) error {
	cur := item.head
	var before *node
	for i := 0; i < index; i++ {
		if cur.next == nil {
			return errors.New("out of range")
		}
		before = cur
		cur = cur.next
	}
	if before != nil {
		before.next = cur.next
	} else {
		item.head = item.head.next
	}
	return nil
}

func (item *LinkList) Len() int {
	return item.len
}

func (item *LinkList) Foreach(f func(v int) int) {
	cur := item.head
	for {
		if cur.next == nil {
			break
		}
		cur.next.value = f(cur.next.value)
	}
}

func NewLinkList() List {
	return &LinkList{
		head: &node{},
	}
}

type node struct {
	value int
	next  *node
}
