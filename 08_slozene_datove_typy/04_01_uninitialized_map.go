// Mapy v programovacím jazyku Go
//
// - deklarace nulové mapy
// - pokus o přidání dvojice klíč-hodnota do mapy

package main

import "fmt"

func main() {
	var m1 map[string]int
	fmt.Println(m1)

	// pokus o přidání nového prvku do mapy
	m1["nula"] = 0
	fmt.Println(m1)
}
