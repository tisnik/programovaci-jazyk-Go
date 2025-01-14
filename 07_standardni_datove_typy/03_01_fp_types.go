// Datové typy pro numerické hodnoty s plovoucí řádovou čárkou:
// float32 a float64.

package main

import "fmt"

func main() {
	// proměnné typu float32
	var a float32 = -1.5
	var b float32 = 1.5
	var c float32 = 1e30
	var d float32 = 1e-30

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// proměnné typu float64
	var e float64 = -1.5
	var f float64 = 1.5
	var g float64 = 1e300
	var h float64 = 1e-300

	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
}
