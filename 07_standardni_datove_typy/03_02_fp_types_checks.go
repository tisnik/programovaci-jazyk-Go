// Kontroly prováděné překladačem, zda numerické hodnoty s plovoucí čárkou
// nepřekračují povolený rozsah.

package main

import "fmt"

func main() {
	// proměnné typu float32
	var a float32 = 1e300
	var b float32 = -1e300

	// proměnné typu float64
	var c float64 = 1e3000
	var d float64 = -1e3000

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
