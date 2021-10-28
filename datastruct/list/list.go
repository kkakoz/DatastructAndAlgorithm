package list


type List interface {
	Front() (int, error)
	Push(value ...int)
	PushFont(value int)
	Delete(index int) error
	Get(index int) (int, error)
	Len() int

}
