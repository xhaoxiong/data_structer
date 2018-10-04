package queue

import (
	"data_structer/arraryList"
	"github.com/spf13/cast"
)



type ArraryQueue struct {
	arraryList.ArraryList
}

func InitNoParam() *ArraryQueue {
	return &ArraryQueue{arraryList.ArraryList{Slice: make([]interface{}, 0)}}
}

func (this *ArraryQueue) Init(cap int) *ArraryQueue {
	return &ArraryQueue{arraryList.ArraryList{Slice: make([]interface{}, 0, cap)}}
}

func (this *ArraryQueue) EnQueue(elem interface{}) {
	this.ArraryList.AddLast(&this.Slice, elem)
}

func (this *ArraryQueue) DeQueue() {
	this.Slice = this.ArraryList.Remove(this.Slice, 0)
}

func (this *ArraryQueue) GetFront() (elem interface{}) {
	elem = this.ArraryList.GetFirst(this.Slice)
	return
}

func (this *ArraryQueue) GetSize() (int) {
	return len(this.Slice)
}

func (this *ArraryQueue) IsEmpty() (bool) {
	if len(this.Slice) == 0 {
		return true
	}

	return false
}

func (this *ArraryQueue) ToString() string {
	s := "front ["
	for index, sl := range this.Slice {

		if index == len(this.Slice)-1 {
			s += cast.ToString(sl) + "] tail"
		} else {
			s += cast.ToString(sl) + ","
		}
	}
	return s
}
