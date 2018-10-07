package linkList

import (
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type LinkList struct {
	node
	dummyHead *node //虚拟头节点
	size      int
}

type node struct {
	Elem interface{}
	Next *node
}

func main() {
	LinkList := &LinkList{dummyHead: &node{nil, nil}}
	for i := 0; i < 10; i++ {
		LinkList.AddFirst(i)
	}
	LinkList.ToString()
	LinkList.Set(10, 1)
	LinkList.ToString()
	LinkList.RemoveFirst()
	LinkList.ToString()
	LinkList.RemoveLast()
	LinkList.ToString()
}

func Init() *LinkList {
	return &LinkList{dummyHead: &node{nil, nil}}
}

func SliceToList(slice []interface{}) *LinkList {
	link := Init()
	for _, s := range slice {
		link.AddFirst(s)
	}
	return link
}

func (this *LinkList) RemoveRepeat(val interface{}) {
	this.RecursiveRemove(this.dummyHead, val)
}

//递归法Remove链表中重复的值
func (this *LinkList) RecursiveRemove(head *node, elem interface{}) *node {
	if head == nil {
		return nil
	}

	head.Next = this.RecursiveRemove(head.Next, elem)

	if head.Elem == elem {
		return head.Next
	} else {
		return head
	}
}

//正常去除节点
func (this *LinkList) NormalRemove(elem interface{}) {

	prev := this.dummyHead

	for prev.Next != nil {
		cur := prev.Next
		if cur.Elem != elem {
			prev = prev.Next
		}
		prev.Next = cur.Next

	}

	return
}

func (this *LinkList) Remove(index int) (elem interface{}) {
	if index < 0 || index > this.size {
		panic("out of range")
	}
	prev := this.dummyHead

	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	retNode := prev.Next
	prev.Next = retNode.Next
	retNode.Next = nil
	elem = retNode.Elem
	this.size--
	return
}

func (this *LinkList) RemoveFirst() {
	this.Remove(0)
}

func (this *LinkList) RemoveLast() {
	this.Remove(this.size - 1)
}

func (this *LinkList) IsEmpty() (bool) {
	return this.size == 0
}

//递归法给链表添加
func (this *LinkList) AddRepeat(elem interface{}) {
	this.dummyHead = this.RecursiveAdd(this.dummyHead, elem)
}

func (this *LinkList) RecursiveAdd(Node *node, elem interface{}) (n *node) {
	if Node == nil {

		return &node{Elem: elem}
	}

	Node.Next = this.RecursiveAdd(Node.Next, elem)
	return Node
}

func (this *LinkList) Add(elem interface{}, index int) {
	if index < 0 || index > this.size {
		panic(errors.New("out of range "))
	}

	prev := this.dummyHead
	//找到index处对应的前一个node节点
	for i := 0; i < index; i++ {
		prev = prev.Next
	}

	newNode := &node{Elem: elem}
	newNode.Next = prev.Next
	prev.Next = newNode
	this.size++

}

func (this *LinkList) AddFirst(elem interface{}) {
	this.Add(elem, 0)
}

func (this *LinkList) AddLast(elem interface{}) {
	this.Add(elem, this.size)
}

func (this *LinkList) Get(index int) (elem interface{}) {
	if index < 0 || index >= this.size {
		panic("out of range ")
	}

	cur := this.dummyHead.Next

	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	return cur.Elem
}

func (this *LinkList) GetFirst() (elem interface{}) {
	return this.Get(0)
}

func (this *LinkList) GetLast() (elem interface{}) {
	return this.Get(this.size - 1)
}

func (this *LinkList) Set(elem interface{}, index int) {
	if index < 0 || index >= this.size {
		panic("out of range ")
	}

	cur := this.dummyHead.Next

	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	cur.Elem = elem
}

func (this *LinkList) Contains(elem interface{}) (bool) {

	cur := this.dummyHead.Next
	for i := 0; i < this.size; i++ {
		cur = cur.Next
	}

	if cur.Elem == elem {
		return true
	}
	return false
}

func (this *LinkList) ToString() string {
	cur := this.dummyHead.Next
	s := cast.ToString(cur.Elem) + "->"
	for cur.Next != nil {
		cur = cur.Next
		s += cast.ToString(cur.Elem) + "->"

	}
	s += "nil"
	return s
}

func (this *LinkList) GetSize() int {
	return this.size
}
