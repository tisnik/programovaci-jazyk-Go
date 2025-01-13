// Základy programovacího jazyka Go
//
// - automatické odvození typu globálních proměnných

package main

import "fmt"

// deklarace globálních proměnných
// s automatickým odvozením jejich typu
var a = 1
var b = 3.1415927
var c = false
var d = "Hello, world!"
var e = 'e'
var f = 1 + 2i
var g = &a

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
