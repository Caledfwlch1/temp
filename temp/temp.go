package temp

import (
	"fmt"
)

type Iner interface {
	Set(n string, i int)
}

type T1 struct {
	Name string
	Value int
}

type T2 struct {
	Name string
	Value int
}

var (
	_ Iner = (*T1)(nil)
	_ Iner = (*T2)(nil)
)


func (p *T1)Set(n string, i int){
	p.Name = n
	p.Value = i
}

func (p *T2)Set(n string, i int){
	p.Name = n
	p.Value = i
}

func Tmp() {
	var t1 T1
	var t2 T2

	t1 = T1{Name:"aaa", Value:10}
	t2 = T2{Name:"bbb", Value:20}

	fmt.Println("begin", t1, t2)

	inter(&t2)


	fmt.Println("end", t1, t2)
}

func inter(v Iner) {
	v.Set("ccc", 30)
	return
}