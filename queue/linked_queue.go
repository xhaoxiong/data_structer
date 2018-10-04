package queue

import (
	"github.com/spf13/cast"
)

func main() {
}

type LinkedQueue struct {
	*Node
	Size int
}

type Node struct {
	head *Node
	tail *Node
	elem interface{}
	next *Node
}

func Init() *LinkedQueue {
	return &LinkedQueue{&Node{}, 0}
}

func (this *LinkedQueue) GetSize() int {
	return this.Size
}

func (this *LinkedQueue) IsEmpty() bool {
	return this.Size == 0
}

func (this *LinkedQueue) EnQueue(elem interface{}) {
	if this.tail == nil {
		this.tail = &Node{elem: elem}
		this.head = this.tail
	} else {
		this.tail.next = &Node{elem: elem}
		this.tail = this.tail.next

	}
	this.Size++
}

func (this *LinkedQueue) DeQueue() interface{} {
	if this.IsEmpty() {
		panic("out of range")
	}
	retNode := this.head
	this.head = this.head.next
	retNode.next = nil
	if this.head == nil {
		this.tail = nil
	}
	this.Size--
	return retNode.elem
}

func (this *LinkedQueue) GetFront() interface{} {
	if this.IsEmpty() {
		panic("out of range ")
	}
	return this.head.elem
}

func (this *LinkedQueue) ToString() string {
	s := "queue front "
	cur := this.head
	for i := 0; i < this.Size; i++ {
		s += cast.ToString(cur.elem) + "<-"
		cur = cur.next
	}
	s += "tail"

	return s
}
