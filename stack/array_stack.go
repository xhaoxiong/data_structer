package stack

import (
	"data_structer/arraryList"
	"github.com/spf13/cast"
)

type ArraryStack struct {
	arraryList.ArraryList
}

func InitNoParam() *ArraryStack {
	return &ArraryStack{arraryList.ArraryList{Slice: make([]interface{}, 0)}}
}

func Init(cap int) *ArraryStack {
	return &ArraryStack{arraryList.ArraryList{Slice: make([]interface{}, 0, cap)}}
}

func (this *ArraryStack) Push(elem interface{}) {
	this.ArraryList.AddLast(&this.Slice, elem)
}

func (this *ArraryStack) Pop() {
	this.Slice = this.Remove(this.Slice, len(this.Slice)-1)
}

func (this *ArraryStack) Peek() (interface{}) {
	return this.GetLast(this.Slice)
}

func (this *ArraryStack) GetSize() (int) {
	return len(this.Slice)
}

func (this *ArraryStack) IsEmpty() (bool) {
	if len(this.Slice) == 0 {
		return true
	}
	return false
}

func (this *ArraryStack) ToString() string {
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

func (this *ArraryStack) GetCapacity() (interface{}) {
	return cap(this.Slice)
}
