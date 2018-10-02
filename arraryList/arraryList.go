package arraryList

import "fmt"

type ArraryList struct {
	Slice []interface{}
}

func test() {
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

func (this *ArraryList) Add(s []interface{}, elem interface{}, index int) ([]interface{}) {
	ss := make([]interface{}, cap(s)+1)
	copy(ss, s[:index])
	ss[index] = elem
	copy(ss[index+1:], s[index:])
	return ss
}

func (this *ArraryList) AddLast(s *[]interface{}, elem interface{}) {
	*s = append(*s, elem)
}

func (this *ArraryList) Edit(s []interface{}, elem interface{}, index int) ([]interface{}) {
	s[index] = elem
	return s
}

func (this *ArraryList) Remove(s []interface{}, index int) ([]interface{}) {
	if index == 0 {
		ss := this.RemoveFirst(s)
		return ss
	}
	ss := append(s[:index], s[index+1:]...)
	return ss
}

func (this *ArraryList) RemoveFirst(s []interface{}) ([]interface{}) {
	ss := s[1:len(s)]
	return ss
}

func (this *ArraryList) Get(s []interface{}, index int) (interface{}) {
	return s[index]
}

func (this *ArraryList) GetFirst(s []interface{}) (interface{}) {
	return this.Get(s, 0)
}

func (this *ArraryList) GetLast(s []interface{}) (interface{}) {
	return this.Get(s, len(s)-1)
}
