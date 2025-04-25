// Reflexe v programovacím jazyku Go
//
// - zjištění typu hodnoty uložené do proměnné
//   typu any (prázdné rozhraní) s využitím
//   funkce reflect.ValueOf

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

	// tisk typu hodnoty
	fmt.Println("type is: ", typ)
}

// uživatelský datový typ
type user struct {
	name    string
	surname string
}

func main() {
	// proměnná je typu "any", ovšem obsahuje hodnotu typu "user"
	var nil1 interface{} = user{name: "foo", surname: "bar"}
	test_get_type(nil1)
}
