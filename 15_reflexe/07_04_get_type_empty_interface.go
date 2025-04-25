// Reflexe v programovacím jazyku Go
//
// - pokus o zjištění hodnoty "nil" uložené
//   v proměnné typu any (prázdné rozhraní)

package main

import (
	"fmt"
	"reflect"
)

func test_get_type(x any) {
	// získání informací o hodnotě v čase běhu aplikace
	value := reflect.ValueOf(x)

	// zjištění typu hodnoty
	typ := value.Type()

	// tisk typu
	fmt.Println("type is: ", typ)
}

func main() {
	// proměnná typu "any", která je inicializovaná na nulovou hodnotu
	var nil1 interface{} = nil
	test_get_type(nil1)
}
