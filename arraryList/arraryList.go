package arraryList


type ArraryList struct {
	Slice []interface{}
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
	ss := make([]interface{}, 0, cap(s))
	if index == 0 {
		ss = this.RemoveFirst(s)
		return ss
	}
	ss = append(s[:index], s[index+1:]...)
	return ss
}

func (this *ArraryList) RemoveFirst(s []interface{}) ([]interface{}) {
	ss := make([]interface{}, 0, cap(s))

	ss = s[1:len(s)]
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
