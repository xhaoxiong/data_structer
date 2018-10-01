package main

import (
	"important_test/data_structer/arraryList"
	"fmt"
	"github.com/spf13/cast"
)

type Stack struct {
	arraryList.ArraryList
}

func main() {

	stack := InitNoParam()
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}

	stack.Pop()
	fmt.Println(stack.ToString())

}

func InitNoParam() *Stack {
	return &Stack{arraryList.ArraryList{Slice: make([]interface{}, 0)}}
}

func Init(cap int) *Stack {
	return &Stack{arraryList.ArraryList{Slice: make([]interface{}, 0, cap)}}
}

func (this *Stack) Push(elem interface{}) {
	this.ArraryList.AddLast(&this.Slice, elem)
}

func (this *Stack) Pop() {
	this.Slice = this.Remove(this.Slice, len(this.Slice)-1)
}

func (this *Stack) Peek() (interface{}) {
	return this.GetLast(this.Slice)
}

func (this *Stack) GetSize() (int) {
	return len(this.Slice)
}

func (this *Stack) IsEmpty() (bool) {
	if len(this.Slice) == 0 {
		return true
	}
	return false
}

func (this *Stack) ToString() string {
	s := "["
	for index, sl := range this.Slice {

		if index == len(this.Slice)-1 {
			s += cast.ToString(sl)+"] top"
		} else {
			s += cast.ToString(sl) + ","
		}
	}
	return s
}

func (this *Stack) GetCapacity() (interface{}) {
	return cap(this.Slice)
}
