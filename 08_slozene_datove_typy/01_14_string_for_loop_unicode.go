// Řetězce v programovacím jazyce Go
//
// - výpis hexadecimálních hodnot bajtů tvořících řetězec

package main

import "fmt"

func main() {
	var s string = "Příliš žluťoučký kůň"

	for i := 0; i < len(s); i++ {
		fmt.Printf("%02x ", s[i])
	}
}
