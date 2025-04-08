// Pole v programovacím jazyku Go
//
// - průchod všemi prvky pole

package main

import "fmt"

func main() {
	a1 := [...]int32{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}

	// výpis pole prvek po prvku
	for _, item := range a1 {
		fmt.Printf("%d\n", item)
	}
}
