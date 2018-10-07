/**
*@Author: haoxiongxiao
*@Date: 2018/10/6
*@Description: CREATE GO FILE bst
*/
package bst

type BST struct {
	root *node
	node
	Size int
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

func (this *BST) add(node *node, elem int) *node {
	if node == nil {
		this.Size++
		return &node{elem: elem}
	}

	if elem < node.elem {
		node.left = this.add(node.left, elem)
	} else if elem > node.elem {
		node.right = this.add(node.right, elem)
	}
	return node
}
