package main

import (
	"fmt"
	"reflect"
)

type Node2 struct {
	val  int
	next *Node2
}

func main() {
	arr := []interface{}{1, "123", 2.222, &Node2{val: 1}}
	for _, i := range arr {
		a(i)
	}
}

func a(aa interface{}) {
	fmt.Println(reflect.TypeOf(aa))
}
