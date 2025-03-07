// Reflexe v programovacím jazyku Go
//
//   - zjištění typu hodnoty uložené do proměnné x
//     s využitím funkce reflect.ValueOf

package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 42

	// získání informací o hodnotě v čase běhu aplikace
	value := reflect.ValueOf(x)

	// zjištění typu hodnoty
	typ := value.Type()
	fmt.Println("type is: ", typ)
}
