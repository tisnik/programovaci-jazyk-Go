// Reflexe v programovacím jazyku Go
//
// - zjištění typu pro různé ukazatele

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

type user struct {
	name    string
	surname string
}

func main() {
	var ptr1 *int = nil
	test_get_type(ptr1)

	var ptr2 *rune = nil
	test_get_type(ptr2)

	var ptr3 *float32 = nil
	test_get_type(ptr3)

	var ptr4 *complex64 = nil
	test_get_type(ptr4)

	var ptr5 *bool = nil
	test_get_type(ptr5)

	var ptr6 *[]int = nil
	test_get_type(ptr6)

	a := [...]int{1, 2, 3}
	ptr7 := &a
	test_get_type(ptr7)
}
