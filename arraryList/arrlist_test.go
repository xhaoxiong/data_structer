/**
*@Author: haoxiongxiao
*@Date: 2018/10/5
*@Description: CREATE GO FILE arraryList
*/
package arraryList

import (
	"testing"
	"fmt"
)

func TestArraryList(t *testing.T) {
	//s := make([]interface{}, 10)
	s := make([]interface{}, 0, 10)
	for i := 0; i < cap(s); i++ {
		s = append(s, i)
	}
	fmt.Println("the origin s and cap:", s, cap(s))
	/**
		slice s的增删查改
     */
	arraryList := &ArraryList{}
	s = arraryList.Add(s, 100, 6)
	arraryList.AddLast(&s, 101)
	fmt.Println("add s and cap:", s, cap(s))

	s = arraryList.Edit(s, 100, 5)
	fmt.Println("edit s and cap:", s, cap(s))

	s = arraryList.Remove(s, 5)
	fmt.Println("delete s and cap:", s, cap(s))

	value := arraryList.Get(s, 5)
	fmt.Printf("get s[%d]'s value:%v\n", 5, value)

	s1 := s

	//使用字面量赋值
	s = []interface{}{10, 11, 12, 13, 14, 15, 16}

	//赋值之后不影响之前的s
	fmt.Println("after reset s, s1 and cap:", s1, cap(s1))

	s2 := s1[1:]
	s3 := s[1:]
	fmt.Println(" cap(s2) and cap(s3)", cap(s2), cap(s3))
}