/**
*@Author: haoxiongxiao
*@Date: 2018/10/5
*@Description: CREATE GO FILE queue
*/
package queue

import (
	"testing"

	"fmt"
)

func TestArraryQueue_EnQueue(t *testing.T) {
	arraryQueue := InitNoParam()
	for i := 0; i < 10; i++ {
		arraryQueue.EnQueue(i)
	}
	fmt.Println(arraryQueue.ToString())
	arraryQueue.DeQueue()
	fmt.Println(arraryQueue.ToString())

}

func TestLinkedQueue_EnQueue(t *testing.T) {
	listQueue := Init()
	for i := 0; i < 10; i++ {
		listQueue.EnQueue(i)
	}
	fmt.Println(listQueue.ToString())
	listQueue.DeQueue()
	fmt.Println(listQueue.ToString())

}
