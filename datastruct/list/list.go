package list

import "constraints"

type List[T constraints.Ordered] interface {
	Front() T
	Add(index int, value T)
	AddFront(value T)
	AddLast(value T)
	Delete(index int)
	Get(index int) T
	Len() int
}
