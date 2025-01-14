// Datové typy pro numerické hodnoty s plovoucí řádovou čárkou:
// float32 a float64. Speciální hodnoty těchto typů.

package main

import (
	"fmt"
	"math"
)

func main() {
	// proměnné typu float64
	var a float64 = math.NaN()
	var b float64 = math.Inf(1)
	var c float64 = math.Inf(-1)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println()

	// proměnné typu float32
	var e float32 = float32(a)
	var f float32 = float32(b)
	var g float32 = float32(c)

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
