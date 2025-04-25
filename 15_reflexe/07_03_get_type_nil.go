// Reflexe v programovacím jazyku Go
//
// - zjištění typu "nulového ukazatele"

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
	// ukazatel na celé číslo, který obsahuje nulovou hodnotu
	var nil1 *int = nil
	test_get_type(nil1)
}
