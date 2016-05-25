package main

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

func main() {
	var t1 T1
	var t2 T2

	t1 = T1{Name:"aaa", Value:10}
	t2 = T2{Name:"bbb", Value:20}

	fmt.Println("begin", t1, t2)


	//t1.Inter.Set("ccc", 30)
	inter(&t2)


	fmt.Println("end", t1, t2)
}

func inter(v Iner) {
/*
	switch a := v.(type) {
		case T1:
			a.Set("T1", 30)
			fmt.Println("Type T1 ", a)
		case T2:
			a.Set("T2", 30)
			fmt.Println("Type T2 ", a)
	}
*/
	//v.Set("ccc", 30)

	v.Set("ccc", 30)
	return
}