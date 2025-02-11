// Pole v programovacím jazyku Go
//
// - dvourozměrná pole - matice

package main

import "fmt"

func main() {
	// deklarace matice
	var matrix [4][3]float32

	// tisk původního obsahu matice
	fmt.Printf("Matrix:    %v\n", matrix)

	// modifikace prvků matice
	for j := range matrix {
		for i := range matrix[j] {
			matrix[j][i] = 1.0 * float32(i+1) * float32(j+1)
		}
	}

	// tisk nového obsahu matice
	fmt.Printf("Matrix:    %v\n", matrix)
}
