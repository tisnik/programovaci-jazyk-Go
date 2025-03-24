// Řetězce v programovacím jazyku Go
//
// - deklarace víceřádkového řetězce

package main

import "fmt"

func main() {
	// řetězec deklarovaný na více řádcích
	var s string = `Hello

world!` // backtick je umístěn zde

	// tisk řetězce
	fmt.Println(s)
}
