/**
*@Author: haoxiongxiao
*@Date: 2018/10/6
*@Description: CREATE GO FILE bst
*/
package bst

import (
	"fmt"
	"github.com/spf13/cast"
)

type BST struct {
	root *node
	node
	Size int
	Str  string
}

type node struct {
	left  *node
	right *node
	elem  int
}

func (this *BST) IsEmpty() bool {
	return this.Size == 0
}

func (this *BST) Add(elem int) {

	this.root = this.add(this.root, elem)

}

func (this *BST) add(Node *node, elem int) (n *node) {
	if Node == nil {
		this.Size++
		return &node{elem: elem}
	}

	if elem < Node.elem {
		Node.left = this.add(Node.left, elem)
	} else if elem > Node.elem {
		Node.right = this.add(Node.right, elem)
	}
	return Node
}

func (this *BST) Contains(elem int) bool {
	return this.contains(this.root, elem)
}

func (this *BST) contains(node *node, elem int) bool {
	if node == nil {
		return false
	}

	if node.elem == elem {
		return true
	} else if elem < node.elem {
		return this.contains(node.left, elem)
	} else {
		return this.contains(node.right, elem)
	}

}

func (this *BST) PreOrder() {

}

func (this *BST) preOrder(Node *node) {
	if Node == nil {
		return
	}

	fmt.Println(Node.elem)
	this.preOrder(Node.left)
	this.preOrder(Node.right)
}

func (this *BST) ToString() string {
	return this.GenerateBSTString(this.root, 0)
}

func (this *BST) GenerateBSTString(Node *node, depth int) string {
	if Node == nil {
		this.Str += this.GenerateDepthString(depth) + "nil\n"
		return this.Str
	}

	this.Str += this.GenerateDepthString(depth) + cast.ToString(Node.elem) + "\n"
	this.Str += this.GenerateBSTString(Node.left, depth+1)
	this.Str += this.GenerateBSTString(Node.right, depth+1)
	return this.Str
}

func (this *BST) GenerateDepthString(depth int) string {
	str := ""
	for i := 0; i < depth; i++ {
		str += "--"
	}
	return str
}
