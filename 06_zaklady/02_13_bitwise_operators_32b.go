// Základy programovacího jazyka Go
//
// - bitové operátory pro 32bitové hodnoty

package main

import "fmt"

func main() {
	// dvojice proměnných typu uint32
	// (32bitové celé číslo bez znaménka)
	var x uint32 = 42
	var y uint32 = 100

	// výpis původních hodnot
	fmt.Printf("x =      %032b\n", x)
	fmt.Printf("y =      %032b\n", y)

	// výpis výsledků bitových operací
	fmt.Printf("neg x =  %032b\n", ^x)
	fmt.Printf("neg y =  %032b\n", ^y)
	fmt.Printf("x && y = %032b\n", x&y)
	fmt.Printf("x || y = %032b\n", x|y)
	fmt.Printf("x ^ y =  %032b\n", x^y)
	fmt.Printf("x &^ y = %032b\n", x&^y)
}
