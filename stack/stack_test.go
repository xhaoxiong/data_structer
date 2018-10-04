/**
*@Author: haoxiongxiao
*@Date: 2018/10/5
*@Description: CREATE GO FILE stack
*/
package stack

import (
	"testing"
	"fmt"
)

func TestLinkedStack(t *testing.T) {
	linkStack := InitLinkedStack()

	for i := 0; i < 10; i++ {
		linkStack.Push(i)
	}

	fmt.Println(linkStack.ToString())
	linkStack.Pop()
	fmt.Println(linkStack.ToString())
}

func TestArraryStack(t *testing.T) {
	arraryStack := InitNoParam()
	for i := 0; i < 10; i++ {
		arraryStack.Push(i)
	}
	fmt.Println(arraryStack.ToString())
	arraryStack.Pop()
	fmt.Println(arraryStack.ToString())
}
