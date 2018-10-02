package queue

import (
	"important_test/data_structer/arraryList"
	"github.com/spf13/cast"
	"fmt"
)

type ArraryQueue struct {
	arraryList.ArraryList
}

func main() {
	queue := InitNoParam()
	for i := 0; i < 10; i++ {
		queue.EnQueue(i)
	}
	queue.ToString()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.DeQueue()
	queue.ToString()

}

func InitNoParam() *ArraryQueue {
	return &ArraryQueue{arraryList.ArraryList{Slice: make([]interface{}, 0)}}
}

func Init(cap int) *ArraryQueue {
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
	fmt.Println(this.Slice)
	for index, sl := range this.Slice {

		if index == len(this.Slice)-1 {
			s += cast.ToString(sl) + "] tail"
		} else {
			s += cast.ToString(sl) + ","
		}
	}
	return s
}
