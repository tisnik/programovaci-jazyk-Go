// Základní datové typy v jazyce Go
//
// - speciální komplexní hodnoty získané jako výsledky výpočtů

package main

import "fmt"

func main() {
	a := 1 + 1i
	b := 0 + 0i
	c := 1 + 0i
	d := 0 + 1i

	// dělení komplexní nulou
	x := a / b
	y := c / b
	z := d / b

	// podíl dvou komlexních nul
	w := b / b

	// tisk původních hodnot
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)

	fmt.Println()

	// tisk výsledků operací
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Println("z=", z)
	fmt.Println("w=", w)
}
