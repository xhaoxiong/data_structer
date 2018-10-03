package stack

type stack interface {
	Push(interface{})
	Pop()
	Peek()  interface{}
	GetSize() int
	IsEmpty() bool
	ToString() string
}