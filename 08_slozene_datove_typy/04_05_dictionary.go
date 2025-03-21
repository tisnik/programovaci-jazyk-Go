// Mapy v programovacím jazyku Go
//
// - klasický slovník: klíče i hodnoty jsou řetězci

package main

import "fmt"

func main() {
	m1 := make(map[string]string)
	fmt.Println(m1)

	m1["zero"] = "nula"
	m1["one"] = "jedna"
	m1["two"] = "dva"
	m1["three"] = "tri"
	m1["four"] = "ctyri"
	m1["five"] = "pet"
	m1["six"] = "sest"

	fmt.Println(m1)
}
