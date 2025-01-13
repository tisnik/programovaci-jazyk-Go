// Základy programovacího jazyka Go
//
// - automatické odvození typu globálních proměnných
// - použití bloku var

package main

import "fmt"

var (
	// deklarace globálních proměnných
	// s automatickým odvozením jejich typu
	a = 1
	b = 3.1415927
	c = false
	d = "Hello, world!"
	e = 'e'
	f = 1 + 2i
	g = &a
)

func main() {
	// tisk typů i hodnot proměnných
	fmt.Printf("%T\t %v\n", a, a)
	fmt.Printf("%T\t %v\n", b, b)
	fmt.Printf("%T\t %v\n", c, c)
	fmt.Printf("%T\t %v\n", d, d)
	fmt.Printf("%T\t %v\n", e, e)
	fmt.Printf("%T\t %v\n", f, f)
	fmt.Printf("%T\t %v\n", g, g)
}
