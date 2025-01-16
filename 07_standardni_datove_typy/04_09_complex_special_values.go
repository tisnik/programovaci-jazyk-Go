// Základní datové typy v jazyce Go
//
// - speciální komplexní hodnoty z balíčku math/cmplx

package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// kladné komplexní nekonečno
	a := cmplx.Inf()

	// záporné komplexní nekonečno
	b := -cmplx.Inf()

	// Not A Number
	c := cmplx.NaN()

	// tisk hodnot
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
}
