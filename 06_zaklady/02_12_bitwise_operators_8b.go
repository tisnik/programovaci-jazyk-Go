// Základy programovacího jazyka Go
//
// - bitové operátory pro osmibitové hodnoty

package main

import "fmt"

func main() {
	// dvojice proměnných typu byte
	// (osmibitové celé číslo bez znaménka)
	var x byte = 42
	var y byte = 100

	// výpis původních hodnot
	fmt.Printf("x =      %08b\n", x)
	fmt.Printf("y =      %08b\n", y)

	// výpis výsledků bitových operací
	fmt.Printf("neg x =  %08b\n", ^x)
	fmt.Printf("neg y =  %08b\n", ^y)
	fmt.Printf("x && y = %08b\n", x&y)
	fmt.Printf("x || y = %08b\n", x|y)
	fmt.Printf("x ^ y =  %08b\n", x^y)
	fmt.Printf("x &^ y = %08b\n", x&^y)
}
