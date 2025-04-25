// Reflexe v programovacím jazyku Go
//
// - zjištění typu pro různé hodnoty

package main

import (
	"fmt"
	"reflect"
)

func test_get_type(x any) {
	value := reflect.ValueOf(x)
	typ := value.Type()
	fmt.Println("type is: ", typ)
}

func main() {
	a := 42
	test_get_type(a)

	b := '*'
	test_get_type(b)

	c := 3.14
	test_get_type(c)

	d := 3 + 2i
	test_get_type(d)

	e := true
	test_get_type(e)

	f := "foobar"
	test_get_type(f)

	g := []int{1, 2, 3}
	test_get_type(g)

	h := [...]int{1, 2, 3}
	test_get_type(h)
}
