// Základy programovacího jazyka Go
//
// - automatické odvození typu proměnné
// - explicitní přetypování hodnot před přiřazením do proměnné

package main

import "fmt"

func main() {
	// deklarace proměnných s automatickým odvozením
	// jejich typu
	a := uint(1)
	b := byte(1)
	c := complex64(1)
	d := complex128(1)
	f := float32(3.1415927)

	// tisk typů i hodnot proměnných
	fmt.Printf("%T\t %v\n", a, a)
	fmt.Printf("%T\t %v\n", b, b)
	fmt.Printf("%T\t %v\n", c, c)
	fmt.Printf("%T\t %v\n", d, d)
	fmt.Printf("%T\t %v\n", f, f)
}
