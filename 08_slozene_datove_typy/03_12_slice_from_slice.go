// Řezy v programovacím jazyku Go
//
// - vytvoření řezu z jiného řezu
// - zjištění, jak se modifikace prvku v jednom řezu
//   projeví na druhém řezu

package main

import "fmt"

func main() {
	// pole
	var a [10]int

	// řez vytvořený z pole (řez polem)
	slice1 := a[4:9]

	// řez vytvořený z jiného řezu (řez řezem)
	slice2 := slice1[3:]

	fmt.Printf("Pole:            %v\n", a)
	fmt.Printf("Délka pole:      %d\n\n", len(a))

	fmt.Printf("Řez 1:           %v\n", slice1)
	fmt.Printf("Délka řezu 1:    %d\n", len(slice1))
	fmt.Printf("Kapacita řezu 1: %d\n\n", cap(slice1))

	fmt.Printf("Řez 2:           %v\n", slice2)
	fmt.Printf("Délka řezu 2:    %d\n", len(slice2))
	fmt.Printf("Kapacita řezu 2: %d\n\n", cap(slice2))

	// modifikace prvků přes druhý řez
	slice2[0] = 99
	slice2[1] = 99

	fmt.Printf("Pole:            %v\n", a)
	fmt.Printf("Řez 1:           %v\n", slice1)
	fmt.Printf("Řez 2:           %v\n", slice2)
}
