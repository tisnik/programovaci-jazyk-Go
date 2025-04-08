// Řezy v programovacím jazyku Go
//
// - vytvoření řezu
// - průchod všemi prvky řezu

package main

import "fmt"

func main() {
	// řez
	slice := []rune{'a', 'b', 'c', 'd', 'e'}

	// průchod prvky řezu
	for _, item := range slice {
		fmt.Printf("%c\n", item)
	}
}
