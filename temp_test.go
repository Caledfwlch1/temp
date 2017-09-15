package main

import (
	"testing"
)

func TestDisk(t *testing.T) {

	t1 := T1{Name: "aaa", Value: 10}
	//t2 := T1{Name:"aaa", Value:10}
	t2 := T1{Name: "ccc", Value: 30}
	inter(&t1)

	if t1.Name == t2.Name && t1.Value == t2.Value {
		t.Log("Test OK")
	} else {
		t.Error("Test Fail")
	}
	return
}
