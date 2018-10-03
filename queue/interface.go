package queue

type Queue interface {
	EnQueue(interface{})
	DeQueue()interface{}
}