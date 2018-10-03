package main

import (
	"github.com/pkg/errors"
	"fmt"
	"github.com/spf13/cast"
)

type Link struct {
	node
	dummyHead *node //虚拟头节点
	size      int
}

type node struct {
	Elem interface{}
	Next *node
}

func main() {
	link := &Link{dummyHead: &node{nil, nil}}
	for i := 0; i < 10; i++ {
		link.AddFirst(i)
	}
	link.ToString()
	link.Set(10, 1)
	link.ToString()
	link.RemoveFirst()
	link.ToString()
	link.RemoveLast()
	link.ToString()
}

func (this *Link) Remove(index int) (elem interface{}) {
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

func (this *Link) RemoveFirst() {
	this.Remove(0)
}

func (this *Link) RemoveLast() {
	this.Remove(this.size - 1)
}

func (this *Link) IsEmpty() (bool) {
	return this.size == 0
}

func (this *Link) Add(elem interface{}, index int) {
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

func (this *Link) AddFirst(elem interface{}) {
	this.Add(elem, 0)
}

func (this *Link) AddLast(elem interface{}) {
	this.Add(elem, this.size)
}

func (this *Link) Get(index int) (elem interface{}) {
	if index < 0 || index >= this.size {
		panic("out of range ")
	}

	cur := this.dummyHead.Next

	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	return cur.Elem
}

func (this *Link) GetFirst() (elem interface{}) {
	return this.Get(0)
}

func (this *Link) GetLast() (elem interface{}) {
	return this.Get(this.size - 1)
}

func (this *Link) Set(elem interface{}, index int) {
	if index < 0 || index >= this.size {
		panic("out of range ")
	}

	cur := this.dummyHead.Next

	for i := 0; i < index; i++ {
		cur = cur.Next
	}

	cur.Elem = elem
}

func (this *Link) Contains(elem interface{}) (bool) {

	cur := this.dummyHead.Next
	for i := 0; i < this.size; i++ {
		cur = cur.Next
	}

	if cur.Elem == elem {
		return true
	}
	return false
}

func (this *Link) ToString() {
	cur := this.dummyHead.Next
	s := cast.ToString(cur.Elem) + "->"
	for cur.Next != nil {
		cur = cur.Next
		s += cast.ToString(cur.Elem) + "->"

	}
	s += "nil"
	fmt.Println(s)
}
