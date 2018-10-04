package stack

import "data_structer/linkList"

type LinkedStack struct {
	*linkList.LinkList
}

func InitLinkedStack() *LinkedStack {
	return &LinkedStack{linkList.Init()}
}

func (this *LinkedStack) Pop() {
	this.RemoveFirst()
}

func (this *LinkedStack) Push(elem interface{}) {
	this.AddFirst(elem)
}

func (this *LinkedStack) Peek() interface{} {
	return this.GetFirst()
}

func (this *LinkedStack) GetSize() int {
	return this.LinkList.GetSize()
}

func (this *LinkedStack) IsEmpty() bool {
	return this.LinkList.IsEmpty()

}

func (this *LinkedStack) ToString() string {
	s := "stack top "
	s += this.LinkList.ToString()

	return s
}
