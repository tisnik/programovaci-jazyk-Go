// Základy programovacího jazyka Go
//
// - konstanty inicializované konstantím výrazem

package main

import "fmt"

const (
	x = 10
	y = 20 * x
	z = x + y<<1
)

func main() {
	const w = x * 0.5

	// výpis hodnot a typů všech čtyř konstant
	fmt.Printf("x=%v (%T)\n", x, x)
	fmt.Printf("y=%v (%T)\n", y, y)
	fmt.Printf("z=%v (%T)\n", z, z)
	fmt.Printf("w=%v (%T)\n", w, w)
}
