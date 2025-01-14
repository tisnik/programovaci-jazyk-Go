// Datové typy pro numerické hodnoty s plovoucí řádovou čárkou:
// float32 a float64. Hexadecimální mantisa a exponent.

package main

import (
	"fmt"
)

func main() {
	// proměnná typu float64
	var a float64

	fmt.Println(a)

	// jedna z možností zápisu nuly
	a = 0x0
	fmt.Println(a)

	// hexadecimální mantisa i exponent: 2^0
	a = 0x1p0
	fmt.Println(a)

	// hexadecimální mantisa i exponent: 2^1
	a = 0x1p1
	fmt.Println(a)

	// hexadecimální mantisa i exponent: 2^-1
	a = 0x1p-1
	fmt.Println(a)

	// také je možné použít řádovou tečku
	a = 0x15.p-2
	fmt.Println(a)

	// oddělovače
	a = 0xFF_FF
	fmt.Println(a)

	// oddělovače + exponent (vyjde "skoro 2")
	a = 0xFF_FFp-15
	fmt.Println(a)
}
