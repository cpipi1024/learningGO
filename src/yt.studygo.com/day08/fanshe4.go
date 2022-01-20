package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  int
}

func main() {
	var p = person{}
	v := reflect.ValueOf(&p)
	fmt.Printf("%T \n", v)
	fmt.Printf("%T \n", v.Elem())
}
