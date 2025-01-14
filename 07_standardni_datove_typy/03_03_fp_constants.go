// Datové typy pro numerické hodnoty s plovoucí řádovou čárkou:
// float32 a float64. Zápis literálů s numerickými hodnotami.

package main

import (
	"fmt"
)

func main() {
	// proměnná typu float64
	var a float64

	fmt.Println(a)

	// jedna z možností zápisu nuly
	a = 0.
	fmt.Println(a)

	// oddělení desetinné části tečkou
	a = 0.1
	fmt.Println(a)

	// zápis s exponentem
	a = 1.3e10
	fmt.Println(a)

	// záporný exponent
	a = 1.3e-10
	fmt.Println(a)

	// explicitně zapsaný kladný exponent
	a = 1.3e+10
	fmt.Println(a)

	// exponent zapsaný velkým znakem E
	a = 1.3e10
	fmt.Println(a)

	// lze použít i oddělovače
	a = 1234_5678.
	fmt.Println(a)

	// oddělovač v exponentu
	a = 1234_5678.e1_2
	fmt.Println(a)
}
