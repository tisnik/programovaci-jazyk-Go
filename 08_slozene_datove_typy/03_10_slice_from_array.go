// Řezy v programovacím jazyku Go
//
// - vytvoření řezů z pole

package main

import "fmt"

func main() {
	// pole
	var a1 [100]byte

	// pole
	var a2 [100]int32

	fmt.Printf("Délka pole 1:   %d\n", len(a1))
	fmt.Printf("Délka pole 2:   %d\n", len(a2))
	fmt.Println()

	// řez vytvořený z prvního pole
	var slice0 []byte = a1[:]
	fmt.Printf("Délka řezu 0:     %d\n", len(slice0))
	fmt.Printf("Kapacita řezu 0:  %d\n", cap(slice0))
	fmt.Println()

	// řez vytvořený z prvního pole
	var slice1 []byte = a1[10:20]
	fmt.Printf("Délka řezu 1:     %d\n", len(slice1))
	fmt.Printf("Kapacita řezu 1:  %d\n", cap(slice1))
	fmt.Println()

	// řez vytvořený z prvního pole
	var slice2 = a1[20:30]
	fmt.Printf("Délka řezu 2:     %d\n", len(slice2))
	fmt.Printf("Kapacita řezu 2:  %d\n", cap(slice2))
	fmt.Println()

	// řez vytvořený z prvního pole
	slice3 := a1[30:40]
	fmt.Printf("Délka řezu 3:     %d\n", len(slice3))
	fmt.Printf("Kapacita řezu 3:  %d\n", cap(slice3))
}
