package stack

type Stack[T any] struct {
}

func (s Stack[T]) Push(v T) {

}

func (s Stack[T]) Pop() T {
	panic("stack is empty")
}

func (s Stack[T]) isEmpty() bool {

}