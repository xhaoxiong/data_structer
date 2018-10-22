/**
*@Author: haoxiongxiao
*@Date: 2018/10/6
*@Description: CREATE GO FILE bst
*/
package bst

import (
	"testing"
)

func TestBST_Add(t *testing.T) {
	bst := &BST{}
	var nums = []int{5, 3, 6, 8, 4, 2}
	for _, num := range nums {
		bst.Add(num)
	}

	bst.levelOrder()
	bst.NormalPreOrder()
	//fmt.Print(bst.ToString())
	//fmt.Println(bst.Contains(10))
	//fmt.Println(bst.Contains(5))
}

/*
5
--3
----2
------nil
------nil
----4
------nil
------nil
--6
----nil
----8
------nil
------nil
*/
