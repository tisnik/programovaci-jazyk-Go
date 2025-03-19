// Řezy v programovacím jazyku Go
package main

import "fmt"

func main() {
	var a [10]int

	slice1 := a[4:9]
	slice2 := slice1[3:]

	fmt.Printf("Pole:            %v\n", a)
	fmt.Printf("Délka pole:      %d\n\n", len(a))

	fmt.Printf("Řez 1:           %v\n", slice1)
	fmt.Printf("Délka řezu 1:    %d\n", len(slice1))
	fmt.Printf("Kapacita řezu 1: %d\n\n", cap(slice1))

	fmt.Printf("Řez 2:           %v\n", slice2)
	fmt.Printf("Délka řezu 2:    %d\n", len(slice2))
	fmt.Printf("Kapacita řezu 2: %d\n\n", cap(slice2))

	slice2[0] = 99
	slice2[1] = 99

	fmt.Printf("Pole:            %v\n", a)
	fmt.Printf("Řez 1:           %v\n", slice1)
	fmt.Printf("Řez 2:           %v\n", slice2)
}
