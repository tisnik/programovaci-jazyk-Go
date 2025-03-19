// Řezy v programovacím jazyku Go
package main

import "fmt"

func main() {
	var slice []int

	for i := 1; i < 11; i++ {
		slice = append(slice, i)
		fmt.Printf("Obsah řezu:    %v\n", slice)
		fmt.Printf("Délka řezu:    %d\n", len(slice))
		fmt.Printf("Kapacita řezu: %d\n\n", cap(slice))
	}
}
