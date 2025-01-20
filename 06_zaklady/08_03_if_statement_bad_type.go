// Základy programovacího jazyka Go
//
// - řídicí konstrukce if a hodnoty, které nejsou typu boolean.
// POZOR: tento příklad nebude možné přeložit

package main

import "fmt"

func main() {
	if 1 {
		fmt.Println("true")
	}

	if 0 {
		fmt.Println("false")
	}

	if !1 {
		fmt.Println("false")
	}

	if !0 {
		fmt.Println("true")
	}

	var b1 int = 1

	if b1 {
		fmt.Println("true")
	}
	if !b1 {
		fmt.Println("false")
	}

	b2 := 1

	if b2 {
		fmt.Println("true")
	}
	if !b2 {
		fmt.Println("false")
	}
}
