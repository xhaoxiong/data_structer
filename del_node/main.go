package main

import (
	"data_structer/linkList"
	"fmt"
)

func main() {

	var s = []interface{}{25, 25, 25, 25, 27, 28, 29, 30}

	link := linkList.SliceToList(s)
	fmt.Println(link.ToString())

	link.NormalRemove(25)
	fmt.Println(link.ToString())
}


//leetcode题解
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	} else {
		return head
	}
}
