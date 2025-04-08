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
	for index, item := range slice {
		fmt.Printf("%d %c\n", index, item)
	}
}
