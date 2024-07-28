package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := interface{}(11)
	b := interface{}(11)
	fmt.Println(a == b, reflect.TypeOf(a).Name(), reflect.TypeOf(b).Name())
}
