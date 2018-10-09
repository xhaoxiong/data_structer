/**
*@Author: haoxiongxiao
*@Date: 2018/10/6
*@Description: CREATE GO FILE bst
*/
package bst

import (
	"fmt"
	"github.com/spf13/cast"
	"data_structer/stack"
	"data_structer/queue"
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
	this.preOrder(this.root)
}

func (this *BST) InOrder() {
	this.inOrder(this.root)
}

func (this *BST) PostOrder() {
	this.postOrder(this.root)
}

func (this *BST) preOrder(Node *node) {
	if Node == nil {
		return
	}

	fmt.Println(Node.elem)
	this.preOrder(Node.left)
	this.preOrder(Node.right)
}

//非递归遍历
func (this *BST) NormalPerOrder() {
	Stack := stack.InitLinkedStack()

	Stack.Push(this.root)
	for !this.IsEmpty() {
		if cur, ok := Stack.Peek().(*node); ok {

			Stack.Pop()
			fmt.Println(cur.elem)

			if cur.right != nil {
				Stack.Push(cur.right)
			}

			if cur.left != nil {
				Stack.Push(cur.left)
			}
		} else {
			break
		}
	}

}

func (this *BST) inOrder(Node *node) {
	if Node == nil {
		return
	}

	this.inOrder(Node.left)
	fmt.Println(Node.elem)
	this.inOrder(Node.right)

}

func (this *BST) postOrder(Node *node) {
	if Node == nil {
		return
	}

	this.postOrder(Node.left)
	this.postOrder(Node.right)
	fmt.Println(Node.elem)
}

//层序遍历
func (this *BST) levelOrder() {
	Queue := queue.Init()
	Queue.EnQueue(this.root)

	for !Queue.IsEmpty() {
		cur := Queue.DeQueue().(*node)
		fmt.Println(cur.elem)

		if cur.left != nil {
			Queue.EnQueue(cur.left)
		}

		if cur.right != nil {
			Queue.EnQueue(cur.right)
		}

	}

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

//寻找二分搜索树的最小元素
func (this *BST) Minimum() (int) {
	if this.Size == 0 {
		panic("error")
	}
	return this.minimum(this.root).elem
}

func (this *BST) minimum(Node *node) (*node) {
	if Node.left == nil {
		return Node
	}

	return this.minimum(Node.left)

}

//寻找二分搜索树的最小元素
func (this *BST) Maximum() (int) {
	if this.Size == 0 {
		panic("error")
	}
	return this.maximum(this.root).elem
}

func (this *BST) maximum(Node *node) (*node) {
	if Node.right == nil {
		return Node
	}

	return this.maximum(Node.left)

}

//删除最小值
func (this *BST) RemoveMin() (int) {

	this.root = this.removeMin(this.root)
	return this.Minimum()
}

func (this *BST) removeMin(Node *node) (*node) {
	if Node.left == nil {
		rightNode := Node.right
		Node.right = nil
		this.Size--
		return rightNode
	}
	Node.left = this.removeMin(Node.left)
	return Node
}

//删除最大值
func (this *BST) RemoveMax() (int) {
	this.root = this.removeMax(this.root)
	return this.Maximum()
}

func (this *BST) removeMax(Node *node) (*node) {
	if Node.right == nil {
		leftNode := Node.left
		Node.left = nil
		this.Size--
		return leftNode
	}
	Node.right = this.removeMin(Node.right)
	return Node
}
