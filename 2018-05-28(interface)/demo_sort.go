package main

import (
	"sort"
	"fmt"
)

//使用sort包对任意类型元素的集合进行排序
type Person struct {
	name string
	age  int
}

func (p *Person)String()string{
	return fmt.Sprintf("{%s,%d}",p.name,p.age)
}

type PersonSlice []*Person

func (slice PersonSlice) Less(i, j int) bool {
	if slice[i].age < slice[j].age {
		return true
	} else if slice[i].age > slice[j].age {
		return false
	} else {
		return slice[i].name < slice[j].name
	}
}
func (slice PersonSlice) Swap(i,j int){
	slice[i],slice[j] = slice[j],slice[i]
}

func (slice PersonSlice) Len()int  {
	return len(slice)
}

func main()  {
	p1 :=Person{"aaa",30}
	p2 :=Person{"bbb",29}
	p3 :=Person{"ccc",31}
	p4 :=Person{"ddd",29}
	p5 :=Person{"eee",35}

	s1:=make([]*Person,0,10)
	s1 = append(s1,&p1,&p2,&p3,&p4,&p5)
	//s1 := PersonSlice([]*Person{&p1,&p2,&p3,&p4,&p5})
	s2:=PersonSlice(s1)
	sort.Sort(s2)

	fmt.Println(s2)



}
