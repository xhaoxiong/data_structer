/**
*@Author: haoxiongxiao
*@Date: 2018/10/4
*@Description: CREATE GO FILE linkList
*/
package linkList

import (
	"testing"
	"fmt"
)

func TestLinkList_Add(t *testing.T) {
	LinkList := &LinkList{dummyHead: &node{nil, nil}}
	for i := 0; i < 10; i++ {
		LinkList.AddFirst(i)
	}
	fmt.Println(LinkList.ToString())

}

func TestSliceToList(t *testing.T) {
	var s []interface{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	LinkList := &LinkList{dummyHead: &node{nil, nil}}
	LinkList = SliceToList(s)

	fmt.Println(LinkList.ToString())
}

func TestLinkList_RemoveRepeat(t *testing.T) {
	LinkList := &LinkList{dummyHead: &node{nil, nil}}
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			LinkList.AddFirst(i)
		} else {
			LinkList.AddFirst(10)
		}
	}

	fmt.Println(LinkList.ToString())
	LinkList.RemoveRepeat(10)
	fmt.Println(LinkList.ToString())
}

func TestLinkList_NormalRemove(t *testing.T) {
	LinkList := &LinkList{dummyHead: &node{nil, nil}}
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			LinkList.AddFirst(i)
		} else {
			LinkList.AddFirst(10)
		}
	}

	fmt.Println(LinkList.ToString())
	LinkList.NormalRemove(10)
	fmt.Println(LinkList.ToString())
}
