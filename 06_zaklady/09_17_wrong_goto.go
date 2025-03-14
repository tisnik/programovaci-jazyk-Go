// Základy programovacího jazyka Go
//
// - neplatné použití goto pro skok mezi funkcemi

package main

import "fmt"

func foo() {
	// cíl skoku
Target:
	fmt.Println("foo")
}

func main() {
	// skok do jiné funkce
	goto Target
}
