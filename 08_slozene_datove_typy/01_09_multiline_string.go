// Řetězce v programovacím jazyce Go
//
// - deklarace víceřádkového řetězce

package main

import "fmt"

func main() {
	// řetězec deklarovaný na více řádcích
	var s string = `Hello

world!
` // backtick je umístěn na samostatném řádku
	fmt.Println(s)
}
